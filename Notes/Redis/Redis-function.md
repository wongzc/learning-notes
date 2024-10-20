# Redis function study note
### source: xiaolincoding
https://xiaolincoding.com/redis/

1. Expiration Deletion Policy
    - set by:
        1. `expire <key> <seconds>`
        2. `pexpire <key> <milliseconds>`
        3. `expireat <key> <timestamp>` expire at sepcific timestamp
        4. `pexpireat <key> <timestamp>`
        5. `set <key> <value> ex <seconds>`
        6. `set <key> <value> pex <milliseconds>`
        7. `setex <key> <seconds> <value>`
    - check by:
        1. `ttl <key>`
    - cancel by
        1. `persist <key>`
    - when query key, redis check if key in expires table, if no, return value, if yes, cehck if expire before return
    - 3 policies
        1. Scheduled deletion
            - when cerate key, create a schedule event to delete key
            - pros:
                - makesure ley deletion asap, good for memory
            - cons:
                - too much expiry key may cause CPU pressure when deleting
        2. Lazy deletion
            - when query key, check if expire, delete if yes
            - pros:
                - good for CPU, only check key when query
            - cons:
                - memory pressure when too much expired key
        3. Periodic deletion
            - periodically, randomly choose key to check and delete expired key
            - pros:
                - less cpu & memory impact
            - cons:
                - not best for both cpu & memory
                - hard to gauge execute number & frequency, too much, become schedule, too less become lazy
    - what policy redis using?
        - lazy + periodic
        - lazy:
            - using function `expireIfNeeded`
            - before update/ query, will use this function to check
            - deletion can be async or sync, depends on `lazyfree_lazy_expire` setting
                - return `null` after deletion
        - periodic:
            - every 10 secs ( default, can change in redis.conf)
            - check 20 random key
                - if more than 5 (out of 20),  continue to pick 20 random and check
                - else, stop and wait
                - max execute time 25 ms to avoid blocking

2. Eviction Policy
    - use eviction policy to delete key when redis memory exceed threshold
    - set at `maxmemory <byte>` in redis.conf
        - default value for 64 bit is 0, ie, no limit
        - default for 32 bit is 3G, as max memory is 4GB for 32 bit system
    - 8 type of policy
        1. noeviction:
            - default
            - error when trying to insert after maxout memory
            - but still can read/ del
        2. volatile-random: 
            - random delete key that set with expiry time
        3. volatile-ttl:
            - ttl shorter will be deleted first
        4. volatile-lru:
            - least recent used key deleted first
        5. volatile-lfu
            - least frequent used key deleted first
        6. allkeys-random
        7. allkeys-lru
        8. allkeys-lfu
    - check current evicton policy
        - `config get maxmemory-policy`
    - update eviction policy. 2 method:
        1. use `config set maxmemory-policy <policy>`
            - effective immediately
            - lost after restart
        2. update redis config `maxmemory-policy <policy>`
            - persist after restart
            - need to restart to be effective
    - how redis LRU ?
        - not by linked list, as extra cost to move data when visited, and extra memory to maintain
        - by record last visited time at struct
        - randomly pick 5 and delete the least recently used one
        - but cant solve memory problem that cause by large dataset read in which only 1 time use
            - use LFU!
    - how redis LFU?
        - redis object record number of visit
        - the lru in redis object serve different use for LRU/LFU
            - LRU, 24 bits to record recent used timestamp
            - LFU, 16 bits for last decrement time (timestamp), 8 bit for logistic counter (is frequency, decrese with time, start with 5)
                - when redis visited key: it decay logc, then add up value to logc 
                - `lfu-decay-time` to control logc decay speed, larger, decay slower
                - `lfu-log-factor` to control logc increment, larger, slower

        