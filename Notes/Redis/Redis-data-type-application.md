# Redis data type and application
### source: xiaolincoding
https://xiaolincoding.com/redis/

1. String
    - max 512m
    - string encoding type: int, raw, embstr
        raw & embstr -> SDS
    - example: 
        - if string value is long:
            - encoding: INT
            - ptr to long value
        - if string is str & <=44 byte:
            - encoding: embstr
            - ptr to SDS
            - use continuos memory for redisObject & SDS
                - less call of memory allocation/ free up
                - better use of CPU cache
            - when changing str length, memory relocate needeD!
                - embstr is read only, when updating, it will become raw and execute
        - if string is str & >44 byte: ( exact number depends on redis version)
            - encoding: raw
            - ptr to SDS
            - use non-continuos memory for redisObject & SDS
                - 2 call for memory allocation/ free up
    - command:
        1. `SET <key> <value>` set value for a key
        2. `GET <key>` return value
        3. `EXISTS <key>`
        4. `STRLEN <key>`
        5. `DEL <key>`
        6. `MSET <key1> <value1> <key2> <value2>` set multiple key & value
        7. `MGET <key1> <key2>`
        8. `INCR <key>` increase by1 ( for number value only)
        9. `INCRBY <key> 10` increase by 10
        10. `DECR <key>`
        11. `DECRBY <key> 10`
        12. `EXPIRE <key> 60` set key to expire after 60 seconds, default is never
        13. `TTL <key>` check key TTL
        14.  `SET <key> <value> EX 60` or `SETEX <key> 60 <value>` set expiry when create key-value pair
            - `SET <key> <value> PX 60` for miliseconds
        15.  `SET NX <key> <value>` create if doesnt exist

    - application 
        1. caching:
            - cache whole json, eg: `SET user:1 '{"name":"xiaolin", "age":18}'`
            - use key to segregate, eg: `MSET user:1:name: xiaolin user:1:age 18 user:2:name haha user:2:age 99`
        2. counting:
            - eg: `INCR abc:readcount`
        3. distributed lock
            - use `SET NX` to set if key doesnt exist
            - usualy will `PX` to set expiry
            - when unlock:
                1. compare value from called if equal to value of key in memory
                2. call `DEL` to remove key
        4. information sharing
            - use same redis to store the session informatio of a distributed system

2. List
    - array, element order by insert sequence
    - can prepend or append
    - max: 2^32 -1 ( 40 billion elements)
    - command:
        1. `LPUSH <key> <value1> <value2>`
        2. `RPUSH <key> <value1> <value2>`
        3. `LPOP <key>`
        4. `RPOP <key>`
        5. `LRANGE <key> <start index> <stop index>` eg: `LSTART mylist 0 -1`
        6. `BLPOP <key1> <key2> timeout` pop from left from key1, if no, key 2, else block for timeout second, if timeout==0, block indefinitely
        7. `BRPOP <key1> <key2> timeout`
    - application
        1. messaging queue
            - key for messageing queue:
                1. ordering
                    - List is FIFO, can use `LPUSH` + `RPOP` for ordering
                    - to avoid consumer keep using while loop to `RPOP` ( and waste CPU)
                        - use BRPOP, block untill new element available
                2. handle duplicated
                    - use ID for each message, and consumer need to record consumed ID
                    - need to add in ID when  insert into list
                3. reliable
                    - `BRPOPLPUSH` read from 1 list and push to another list
                        - `BRPOPLPUSH <source> <destination> timeout`
                    - to keep the message that is reading, incase system down after read
                        - read is 1 time, need to pop to read
            - downside as message queue:
                - cannot multiple consumer to same message, as message removed when a consumer pop it

3. Hash
    - key value pair
    - command:
        1. `HSET <key> <field> <value>` set the field value for a key
        2. `HGET <key> <field>`
        3. `HMSET <key> <field1> <value1> <field2> <value2>`
        4. `HMGET <key> <field1> <field2>`
        5. `HDEL <key> <field1> <field2>`
        6. `HLEN <key>` number of field
        7. `HGETALL <key>` retuen all field and value
        8. `HINCBY <key> <field> n`
    - application
        1. caching
            - use string + json to store
            - but for constant chaging value, can take out and save as Hash

4. Set
    - command:
        1. `SADD <key> <value1> <value2>`
        2. `SREM <key <value1> <value2>` remove value
        3. `SMEMBERS <key>` get values from key
        4. `SCARD <key>` get number of element
        5. `SISMEMBER <key> <value>` check if value in set
        6. `SRANDMEMBER <key> <count>` get count of element from set
        7. `SPOP <key> <count>` similar to SRANDMEMBER, but remove after get, also random!
        8. `SINTER <key1> <key2> <key3>` get intersect
        9. `SINTERSTORE <destination> <key1> <key2> <key3>` save intersect to destination
        10. `SUNION <key1> <key2> <key3>`
        11. `SUNIONSTORE <key1> <key2> <key3>`
        12. `SDIFF <key1> <key2> <key3>` get diff
        13. `SDIFFSTORE <key1> <key2> <key3>`
    - application
        1. union, intersect, diff calculation complexity is big, may block redis
        2. like:
            - ensure 1 user 1 like
        3. common following
            - use intersect
        4. lucky draw
            - SPOP (if cannot repeatly get award) or SRANDMEMBER

5. Zset
    - command:
        1. `ZADD <key> <score1> <member1> <score2> <member2>`
        2. `ZREM <key> <member1> <member2>`
        3. `ZSCORE <key> <member>`
        4. `ZCARD <key>` number of element
        5. `ZINCRBY <key> <increment> <member>` add increment to score of member
        6. `ZRANGE <key> <start> <stop> WITHSCORES` get range from lowest to highest, add `WITHSCORES` if need score
            - use (start as exclude boundary
            - default include
        7. `ZREVRANGE <key> <start> <stop> WITHSCORES` get from highest to lowest
        8. `ZRANGEBYSCORE <key> <min> <max> WITHSCORES <LIMIT offset count>`
            - get element within min and max, optional withscores
            - start, end can be -inf or +inf
            - LIMIT 0 1 means skip 0, and only 1 count needed
        9. `ZRANGEBYLEX <key> <min> <max> [LIMIT offset count]`
            - same as above, but sort lexicographical order
            - can use "-" or "+" for all
                - `ZRANGEBYLEX phone - +`
            - min is like (a  (ccc, means all element start with a to ccc
        10. `ZREVRANGEBYLEX key max min [LIMIT offset count]` sort by reverse lexico
        11. `ZUNIONSTORE <destination> <number of set to add in> <key1> <key2>` 
            - `ZUNIONSTORE result 2 zset1 zset2`
        12. `ZINTERSTORE <destination> <number of set to add in> <key1> <key2>`
    - application
        1. ranking board:
            - `ZINCRBY` to increase score
            - `ZSCORE` to check score
            - `ZREVRANGE user:xiaolin:ranking 0 2 WITHSCORES` check top 3 highest
            - `ZRANGEBYSCORE user:xiaolin:ranking 100 200 WITHSCORES` check elements within 100 & 200 score

6. bitMap
    - binary , only 0 & 1
    - locate element using offset
    - save memory
    - underlying structure: string
    - command:
        1. `SETBIT <key> offset value`
        2. `GETBIT <key> offset`
        3. `BITCOUNT <key> start end` get number of 1 in range
        4. `BITOP [operation] [result] [key1] [keyn...]`
            - operation: AND OR XOR NOT
            - for NOT: only 1 key , the other can have n key
        5. `BITOPS [key] [value]`
            - return location of first value (0 or 1) in bitmap
    - application:
        - mark attendance:
            - `SETBIT uid:sign:100:202206 2 1` set attendance to 1 for 2022 06 03 for account 100
        - check if user logged in:
            - `SETBIT login_status 10086 1`
            - save memory, 50 mil user only 6 mb!

7. HyperLogLog
    - based on probability, 0.81% of error
    - to get the cardinality number, i.e., number of unique element
        - use constant memory to calculate even data size very big
        - 12 KB, to cal 2^64 element
        - compare to java, need 2^64*8/1024 KB for same amount of nuber
    - command:
        1. `PFADD <key> <element1> <element2>`
        2. `PFCOUNT <key1> <key2>` count cardinality
        3. `PFMERGE <destination> <source1> <source2>` merge hyperloglog
    - application
        1. unique visitor on a multi million website
            - but may have 0.81% of error

8. Geo
    - for location-Based service (LBS), to check nearby etc
    - underlying: sorted set
    - use geohash encoding to change longitude latitude to score
    - command:
        1. `GEOADD <key> <longitude> <latitude> <member>` add long, lat and member (name) into a key
        2. `GEOPOS <key> <member>` return location of member, `nil` if doesnt exist
        3. `GEODIST <key> <member1> <member2> [m|km|ft|mi]` get distance, optional unit
        4. `GEORADIUS key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]`
            - return set based range & location
    - application
        - uber, grab

9. Stream
    - specifically for message queue ( create on redis ver 5.0)
        - before Stream, list was used for message queue
            - pub-sub method, cannot save message
            - offline user cannot access history of message
            - cannot repetaly consume on same message, 1 consum will delete message
            - producer need to create the unique ID itself
    - command:
        1. `XADD <key> * <field1> <value1>`
            - `*` to tell redis use auto-generated unique ID
            - replace `*` if want to specify own ID
            - after execute, id will be return (example: "1654254953808-0" ), id has 2 part
                - 1654254953808 is server time in miliseconds
                - -0 is the nth message in this ms
        2. `XLEN`
        3. `XREAD [COUNT <count>] [BLOCK <milliseconds>] STREAMS <key> <id>`
            - count: limit number of entry return
            - block: block by ms if no entries available, ignore for no waiting
            - can read from multiple streams
            - will get all message after the specified ID
        4. `XDEL` delete based of message id
        5. `DEL` delete the stream
        6. `XRANGE` return message within range
        7. `XREADGROUP GROUP <groupname> <consumername> [COUNT <count>] [BLOCK <milliseconds>] STREAMS <stream> [<stream> ...] <id> [<id> ...]`
            - reading message from a stream in context of consumer groups
            - id can be '>', means start from first uncomsumed message (unconsumed within the group)
                - diff group still can read the same message
        8. `XPENDING <key> <groupname> [<start> <end> <count> [<consumer>]]` to check for each consumer group, message that consumer read but havent ack
            - can use - + as start end
            - optional to  specify consumer
        9. `XACK` to acknowledge message has been processed
            - stream use a PENDING list to save each message consumed by consumer group, untill it `XACK`
                - make it more reliable, and message can be restore even after system down
        10. `XGROUP`
            - `XGROUP CREATE <stream> <groupname> <id or $> [MKSTREAM]`
                - id: start from id, 0-0: start from begining
                - $: start from new entries, i.e., after group created
                - MKSTREAM: cerate stream if doesnt exist
            - `XGROUP SETID <stream> <groupname> <id or $>`
                - change the ID whwre consumer group start reading
            - `XGROUP DELCONSUMER <stream> <groupname> <consumername>`
                - remove consumer
            - `XGROUP DESTROY <stream> <groupname>`
                - delete consumergroup
    - redis compare to kafka/ rabbitMQ:
        - message may be lost when it is in stream, when:
            - AOF asynchronously writing to  disk
            - Master-Slave Replication, when master & slave swap
            - when stream length larger than max, older will be deleted ( to avoid OOM, out of memory)
        - kafka, rabbit MQ will save data in disk, so no OOM
        - tehy will use multiple node to ensure if 1 node down, the other node working and preserve data
