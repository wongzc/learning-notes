# Redis availability
### source: xiaolincoding
https://xiaolincoding.com/redis/

1. Availability
    - AOF & RDB to make sure most data persist when system restart
    - but if system fail:
        - restart take time, cant take in new request
        - disk issue, then data lost
    - to solve, backup to other server
        - but how to makesure data consistency between main & backup?
        - redis provide master-slave duplication  for this
            - read/write seggregate: 
                - read on both slave/master
                - write on master, master then write to slave

2. establish master-slave relation
    - on slave server, run `replicaof <ip of master> <port of master>`
    - 3 steps:
        1. establish connection, and synchronize
            - after run `replicaof`, slave will send psync to master to request for sync
                - psync contain `runID` and `offset`
                - the first time, it will be `psync ? -1` 
                    - runID is a unique ID created when redis server start, use "?" as slave doesnt know
                    - offset is the sync progress, -1 for the forst time
            - master send `FULLRESYNC runID offset` back, slave will record runID and offset
                -`FULLERSYNC` means full replicate
        2. master sync data to slave
            - master will run `bgsave` to create RDB ( in subprocess ) and send to slave
            - slave empty its data and load RDB
            - new data write into to master may not in RDB and so in slave
                 - to make data consistent, master will keep write instruction in `replication buffer` during below time:
                    1. when master creatign RDB
                    2. when master sending RDB
                    3. when slave loading RDB
        3. master send new write to slave
            - after slave loaded data, will send ack to master
            - master send data in replication buffer into server
            - data become same between master & slave

3. command propagation based on a long-lived connection
    - after first sync, a TCP will be maintain between master & slave
    - it is a long-lived TCP, to avoid performance overhead by frequent TCP connect/disconnect

4. load distribute for master server
    - 2 time consuming operation in first sync:
        1. create RDB
        2. send RDB
    - big problem if too much slave, main process will be busy at `fork()`, and block main process
    - sending RDB will cause pressure to network bandwidth, impacting request to master.
    - to solve this, create a slave master to incharge of slave. by `replicaof <ip> <port>`

5. Incremental Replication
    - happend after network discnnected between master and slave.
    - to sync data, instead of full RDB, use incremental replication
    - 3 step:
        1. slave send `psync runID offset` offset will be non-1
        2. master send `CONTINUE` to inform slave use incremental
        3. master send data over.
            - master check offset sent by slave and compare to mastr offset.
            - if data still in `repl_backlog_buffer`, use incremental
                - recent data sent by master will be in `repl_backlog_buffer`
                - default 1 M
                - data overwrite older data if full
                - can set it based on estimated_recovery_time_needed * write_size_per_second
                - set in config file `repl-backlog-size`
            - if data not in,  use full RDB

6. summary:
- 3 ways of replication:
    - full RDB
    - command propagtion through long lived TCP
    - incremental replication

7. how to know if a node is working?
    - use ping-pong
    - if more than half node ping it and no pong, will determine as failed, and connection close
    - master every 10 secs send ping to slave to check
        - can be change by `repl-ping-slave-period`
    - slave every 1 sec send `replconf ack {offset}` to master
        - to check network connection
        - to make sure data consistency

8. evicted key will be send as a `del` from master to slave

9. replica is async, by master write to buffer, async send to slave

10. replication buffer vs repl backlog buffer
    - replication buffer:
        - for full RDB to keep new write that need to send to slave later
        - every new slave will have 1
        - when maxout, connection lost, buffer delete, RDB restart
    - repl backlog buffer
        - for incremental replica, to keep instruction, to reuse after master-slave reconnected
        - share 1 for all slave
        - when maxout, overwrite previous data

11. data difference in master & slave:
    - due to async replication
    - after master execute, it return data directly to client, wont wait for slave to finish

12. how to solve data inconsistent in master and slave?
    - makesure good network connection between master & slave, like same network
    - use external program to monitor `master_repl_offset` and `slave_repl_offset`
    - if delta larger than preset threshold, redirect client to another node.
        - but delta cant be too small!

13. how to minimize data lost for master-slave switching?
    - 2 situation may cause data lost in master-slave switch:
        1. async replica data loss
            - happen when master unable to sync new command to slave when system fail
            - to solve:
                - `min-slaves-max-lag`, if all slave delay by this number compare to master, master stop accept request
                - in seconds, time required for master to sync data to slave
                - for client, can set it write to local or kafka queue when master stop accept request
        2. cluster split-brain data loss
            - happen when master connecting with client, but master lost connection with slave
                - client keep sending data to master
                - sentinel detect this and determine master down, a new master will be selected from slave 
                - when network restore, old master being downgrade to slave and data removed.
                    - data sent by client will be lost.
            - to solve:
                - `min-slaves-to-write x` master reject request if less than x number of slave ( too much slave disconnected)
                - `min-slaves-max-lag x` master reject request if master slave delay more than x seconds ( network delay )
14. promote new master when old master down?
    - must use sentinel!

15. Sentinel:
    - in charge of 3 things:
        1. monitor
            - ping all node (master+slave) every second
            - if node didnt reply within time ( `down-after-milliseconds` ), will be marked as **subjective offline**
                - master node maybe delsay response due to system load, so have subjective delay
                - use sentinel cluster to check ( >=3 sentinel)
            - one a node is subjective offline, sentinel send request to gather feedback from other sentinel by `is-master-down-by-addr`
            - then based on `quorum`, if number of sentinel agree reach this number, node will be marked as down
        2. promote master
            1. choose leader from sentinel to execute master-slave switch
                - the sentinel that mark master as subjecive offline is the candidate
                - all sentinel vote to choose leader from candidate ( can be multiple candidate)
                - all sentinel only 1 vote
                - candidate that get move than 50% and >quorum become leader
                    - as quorum used to decide objective offline and choose leader, and choose leader need >50%
                    - if quorum <50%, eventhough mark as objective offline, cant swap.
                    - so suggestion is quorum = n//2+1 (n is number of sentinel)
            2. choose new master
                - filter out offline node
                - filter out node with bad network connection
                - filter out node that had bad network connection. ( based on number of down happened)
                - then choose based on:
                    1. priority:
                        - based on redis config `slave-priority`
                        - usually set for server with larger memory
                    2. replication progress
                        - choose node with `slave_repl_offset` that nearer to `master_repl_offset`
                    3. ID number
                        - smaller ID first
                - then send `SLAVEOF no one` to new master
                - leader will send `INFO` to new master every second and check if it promoted.
                - leader send `SLAVEOF <new master ip> <new master port>` to all slave
            3. notify client new master
                - by redis pub/sub
                - leader send to `+switch-master` channel to inform client
                - client then switch
            4. change old master to slave
                - sentinel keep monitor old master
                - send `SLAVEOF` to it when it online
                
        3. notify
            - notify about switch master, master down, slave relocate
            - pub/sub

16. how to create sentinel cluster?
    - command to create sentinel:
        - `sentinel monitor <master-name> <ip> <redis-port> <quorum>`
        - sentinel detect each other by redis pub/sub
        - use channel `__sentinel__:hello` on master to find out each other
            - by sending ip and port to the channel
    - how sentinel know about slave info?
        - every 10 sec, send `INFO` to master to get!
            
     