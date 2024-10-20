# Redis interview
### source: xiaolincoding
https://xiaolincoding.com/redis/

1. what is redis
    - in memory database, fast
    - caching, message queue, distributed lock
    - support:
        - String
        - Hash
        - List
        - Set
        - Zset ( ordered set)
        - Bitmaps
        - HyperLogLog
        - GEO
        - Stream
2. redis vs memcache
    - same
        - both in memory, use for cache
        - high performance
        - cache invalidation
    - different
        - redis support more type of data (String、Hash、List、Set、ZSet)
        - memcache only key-value
        - redis can save data permanently into disk, memcache cannot
        - redis support cluster

3. redis as cache for mysql
    - redis QPS ( query per second) is 10x of MySQL
    - redis cache data that query frequently for next query

4. main data type
    - String
        - cache session token, basic key value storage
    - List
        - add element to head and tail, to implement queue in FIFO
    - Sets
        - unordered collection of unique string
        - union, intersection and difference
    - Zset
        - score associated with each member, allow order retrieval
        - lader biards, ranking system, schedule task priority
    - Hash
        - key-value pair
        - representing objects
    - Bit map
        - compact way to store bits
        - manipulate individual bits
        - track user activity
    - HyperLogLogs
        - A probabilistic data structure used for approximating the cardinality (unique count) of a set
        - count unqiue visitor to website ( for large data set)
    - Streams
        - store multiple entries in time-ordered manner
        - message brokering, real-time analytic
    - Geospatial Indexes
        - Store and query geographic data using longtitude & latitude

5. redis key value pair
    - key: string
    - value: string, list, hash, set, zset
    - hash table to achieve key value pair, can search in O(1)
    - key in hash table are call hash bucket, hold pointer that point to key-value data

6. redis datatype to underlying structure
    - string: SDS
    - list: quicklist
    - hash: listpack + hash table
    - set: hash table + int set
    - zset: skip list + list pack

7. is redis single thread?
    - redis is single thread, as 1 main process handle receive request -> execute request -> send to client
    - but redis program is not single thread, it run BIO ( background I/O) when started
    - 3 BIO:
        1. `BIO_CLOSE_FILE` use clode(fd) to close file
        2. `BIO_AOF_FSYNC` sync AOF to disk
        3. `BIO_LAZY_FREE` free(obj) for command like " unlink key / flushdb async / flushall async"
            - that's why use unlink for big key, instead of del, which block mian process
    - run the 3 as BIO, cause these task are time consuming, will block main process

8. why redis single thread but so fast?
    - redis can handle 100k request per second
    - fast because most operation in memory and using efficient data structre
    - limitation is usually network or memory, not cpu
    - so single thread is enough. and it avoid multiple thread conflict and mutex lock
    - IO multiplexing: 1 thread handle many I/O stream ( select/epoll mechanism)
        - mutiple listener socket and connected socket
        - one request received, send to thread to handle

9. why redis 6.0 introduce multithread?
    - with network hardware become better, sometime network I/O can be bottleneck
    - redis 6.0 use multithread to improce network I/O
    - but for command execution, still single thread

10. how to implement delayed queue in redis?
    - use zset, add data with score
    - use zrangebyscore to get data that within requirement

11. how to handle redis big key?
    - big key is not key itself big, is the value of the key is big
    - string>10kB
    - hash, list, set, zset >5000 length
    - impact of big key:
        - block client due to time consuming of big key
        - network overloaded
        - block main process: del big key need more time
        - memory inbalance due to big key
        - use `redis-cli --bigkeys` to find big key
        - use `scan`
    - how to delete big key:
        - batch deletion: 
            - hash: `hscan` get 100 count every time and `hdel` 1 by 1
            - list: use `ltrim` delete part by part
            - set: `sscan` to get 100 count and `srem` 1 by 1
            - zset: `zremrangebyrank` to delete 100 at 1 time
        - async delete:
            - use `unlink` instead of delete
            - unlink will exxecute delete in async manner
    - config for lazy delete
        - `lazyfree-lazy-eviction no `: use lazy delete when memory max
        - `lazyfree-lazy-expire no` : use LD when key expire
        - `lazyfree-lazy-server-del no`: use LD when command that involve `del`, like `rename` when key alr exist
        - `slave-lazy-flush no`: when slave clean up data before full RDB

12. redis pipeline
    - to execute multiple command at the same time
    - but need to take note when command too big and overload netwrok

13. redis rollback:
    - redis does not support rollback
    - redis does not guarantee atomicity ( i.e., all command in the task to be all success)
    - redis creator think that error usually from wrong syntax, which in dev env, not in prod env. and redis want simplicity

14. redis distributed lock:
    - use `SET NX`
    - pros:
        - efficiency
        - convenient
        - avoid single point of failure ( with cluster)
    - cons:
        - timeout hard to set: need to create a guardian thread, to extend timeout when needed
        - master-slave replica is async, distributed lock maybe not reliable
    - redlock:
        - distributed lock created by redis developer
        - client request lock with multiple indipendent redis node.
        - steps:
            1. get client time (t1)
            2. client sequentially request lock from N redis node
                - use `SET` with `NX/EX/PX`
                - set a timeout for redlock to ensure working when 1 of the node down
                - if client get lock from more than half node & finish time -t1 < threshold : success