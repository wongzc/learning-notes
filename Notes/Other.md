1. Thread vs Process  

    |         | Process | Thread  |  
    |---------|-----|----------|  
    | Memory  | isolated, own memory address  | thread within same process share same address space      |  
    | Communication  between| slow, need to go through inter-process comm, need kernel context switch   | fast, share same address, can direct read/ write memory  |
    | Creation Overhead | high, need allocate resource  | fast |
    | Failure isolate | 1 process fail not affect others | 1 thread fail can crash entire process  |
    | context switch  | expensive, switch memory context | faster as resource share |
    - can run how many depends on CPU
        - hyper-threading: can run 2 thread/core
        - 4-core/8-thread CPU: run truly 8 thread/process simutaneously
    - python thread
        - has thread, but only execute 1 thread at a time in interpreter, due to GIL 
            - which purpose to simplify memory management
            - as python object not memory safe
        - python can use multithread for I/O ( network, file, request)
            - not for CPU bound task
            - when I/O, operations handle by `system calls` which outside of interpreter
                - python run time will release the GIL when making calls
                - so another thread can acquire the GIL and run
        - so python threading and asyncio same purpose
            - but asyncio is single thread, with lower overhead
    - python process
        - use `multiprocessing` for true parallelism
        - each with own interpreter and GIL
        - data must be serialized ( pickled) to send between process
            - serialized into byte stream, and another to load it

2. Network data transfer
    - from local to server
        - app layer
            - app build request and serialized into JSON etc
        - transport layer
            - app send data over TCP/UDP
            - if TCP: split data into segment, with number and retry if lost
        - network layer
            - data wrap in IP packet
            - with a source IP and destination IP
        - link layer
            - ip packet sent as frame via wifi/ 4g/5g
            - router forward frame through ISP
        - routing over internet
            - DNS used to translate domain name to IP
            - packet hops across routers on internet until reaches servers IP
        - server receive
            - server accept packet, assemble into full TCP streams, pass to application
            - then server read and process

3. router and TCP
    - when connect to same wifi, public IP will be same ( to server), but ports different
    - private ip, assigned by router will be different
    - router keep a `Network Access Translation` (NAT) table to match from private ip to public ip+ports
    - when device initiate outgoing TCP connection, it pick random temporary port call Ephemeral port
        - once connection closed, it will be reused later
        - each device even under same network, will have its own TCP
    - if different app in single device, will have its own TCP connection
    - id different request in same app:
        - same host ( website): 
            - share one TCP/TLS
            - use stream id to identify where the data goes to
                - stream id assigned by browser
        - different host
            - own TCP
    - router ARP (address resolution protocol) map private ip to MAC
    - if same devicem same private ip, but differnt port

4. TCP 
    - TCP is uniquely identify by (src IP, src port, dst IP, dst port)
    - must be unique n network stack of both end
        - if 2 device connect to same server and port, but source is differnt
            
    - MAC identify device at layer 2
    - IP identify at layer 3

5. mysql ACID
    - Atomicity
        - all failed or all success
        - use transaction logs ( redo/ undo logs)
        - if transaction failed midway, innoDB rolled back all changes using undo logs
    - Consistency
        - from 1 valid state to another, never violate constraint
        - FK/Check/Data types constraint
        - by Isolation+ atomicity+durability
    - Isolation
        - transaction is isolated, not inferencing others
        - MVCC (Multi-Version Concurrency Control)
            - let transaction read consistent snapshot of data without blocking views
                - by keeping multi version of rows
            - use undo log + readview for MVCC
                - undo log: old version of logs
                - readviews: 
                    - which version a transaction can see
                    - snapshot that created when transaction started
                        - contain 
                            - smallest transaction id that hasnt committed( oldest active, before it all commit)
                                - smaller than this means can see
                            - next transaction id to be assigned
                                - larger than this means created after the readview, dont read
                            - list of current active transaction id when view created
                                - id in here, means not committed
                            - transaction id that own this read view
                        - if not in active id list, and smaller than next transaction id, means can read as well
                        - if the readview found rows that is too new, will try to look back to old version in undolog
        - each transaction see snapshot of data consistency to its isolation level
            - READ UNCOMMITTED
            - READ COMMITTED
            - REPEATABLE READ (default)
            - SERIALIZABLE
    - Durability
        - once transaction committed, the change will still remain even crash
        - redologs flush to disk before aknowledge a commit
    