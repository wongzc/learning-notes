# Redis data type and application
### source: xiaolincoding

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
                    - 
            

3. Hash

4. Set

5. Zset

6. bitMap

7. HyperLogLog

8. Geo

9. Stream