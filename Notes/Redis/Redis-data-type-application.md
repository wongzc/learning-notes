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
        15.  `SETNX <key> <value>` create if doesnt exist

    - application 
        1. caching:
            - cache whole json, eg: `SET user:1 '{"name":"xiaolin", "age":18}'`
            - use key to segregate, eg: `MSET user:1:name: xiaolin user:1:age 18 user:2:name haha user:2:age 99`
        2. counting:
            - eg: `INCR abc:readcount`

2. Hash

3. List

4. Set

5. Zset

6. bitMap

7. HyperLogLog

8. Geo

9. Stream