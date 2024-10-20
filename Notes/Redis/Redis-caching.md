# Redis caching
### source: xiaolincoding
https://xiaolincoding.com/redis/

0. database usually in disk, disk read write is slow, so use redis as cache between client and database

### 3 problem with cache
1. cache avalanche
    - data in redis usually with a TTL
    - when many data TTL ended at the same time, or Redis fail, client request direct to database, causing DB overloaded
    - many data TTL:
        - to solve:
            1. evenly distributed TTL: use some random number to expiration time
            2. mutex lock: when data not in redis, lock db and makesure only 1 request at a time
                - need to set timelimit for mutex lock, incase process fail and db unable unlock
            3. background refresh cache:
                - no expiration set for cache data, background refresh data by schedule
                - when memory full, some of the data will be lost
                    - when client request, data return null!
                    - to fix it, 2 ways:
                        1. background refersh and check if data still valid, if discradded due to memory full, restore it
                        2. when cleint found that data in valid, send to inform background to refresh
    - redis system down:
        - to solve:
            1. stop client request to prevent overload database or limit the request
            2. build highly available redis cluster

2. cache breakdown
    - for some of the frequent access data, if the data expired, all request will go to DB and causing problem
    - to solve:
        1. mutex lock
        2. dont set expiration for frequent access data, use background async refresh

3. cache penetration
    - data not in cache and DB
    - data cant be loaded into cache from db, so request keep coming in
    - why happen:
        - accidentaly deleted data
        - cyber attack
    - to solve:
        1. invalid request limitation
            - API entrance to check if request param valid
        2. cache null/ default value
            - set null/ default value for data that doesnt exist
        3. bloom filter to check if data exist
            - when data write in DB, mark in bloom filter
            - when client request, check from bloom if data exist
            - bloom filter:
                - start with 0 bit map, with n hash function
                - how it works?
                    - use n hash functions to get n hash values
                    - take modulo of n hash values ( mod of filter length)
                    - set the bit as 1
                - when bloom filter show doesnt exist: 100% doesnt exist
                - when bloom show exist, may not exist

4. consistency of cache & DB
    - update DB or update cache first? both may cause inconsistency between cache & DB
    - cache aside:
        - write: delete cache & update db
        - read: if in cache, return , else go to DB get data and put in cache
        - update data first, then delete cache.
            - why? if we delete cache first, before we update db, another request may bring the old data to cache again
            - memory write in is fast. if we write in db, and delete cache, less likely to happen
        - can set expiration for data, so that eventhough cache db mismatch, it can be updated
    - cache aside will affect cache hit rate, so maybe beter to refresh cache as well, after update db
    - use distributed lock to makesure only 1 request to update cache at a time
    - delay deletion:
        -  can be use for first delete cache then update db
        - delete cache, update db, then sleep and delete cache
        - may not be very effective
    
    - to makesure update & delete cache both success
        - 2 method:
            1. message queue retry deletion
                - if deletion fail, retry until limit
                - if success, remove from queue
                - downside: need to modify code a lot
            2. subscribe MySQL binlog, then operate on cache
                - binlog send to MQ, if success, ack
                - no need to modify code a lot, but many import feature
            both async method