1. Connect to redis
    - r = redis.Redis(host='localhost', port=6379, db=0)

2. String Operation
    1. set: `r.set("key","value")`
    2. get: `r.get("key","value")`
    3. increase (must be numeric): `r.incr("counter")`
    4. decrease: `r.decr("counter")`
    5. Append: `r.append("key", " world") `
        - append: orginal_value + new_value
    6. set multiple: `r.mset({"k1":"v1","k2":"v2"})`
    7. get multiple: `r.mget(["k1","k2"])`
    8. delete:`r.del("key")`

3. List Operation
    1. push to left: `r.lpush("mylist","A")`
    2. push to right: `r.rpush("mylist","A")`
    3. pop from left: `r.lpop("mylist")`
    4. pop from right: `r.rpop("mylist")`
    5. get element in range: `r.lrange("mylist",0,-1)` #(0,-1) is get all

4. Set Operation
    1. add: `r.sadd("myset","a',"b","c")`
    2. get all elements: `r.smembers("myset")`
    3. remove element: `r.srem("myset","b")`
    4. check if element in: `r.sismember("myset","a")`
    5. check set intersecr: `r.sinter("set1","set2")`

5. Hash (dictionary) Operation
    1. set: `r.hset("user:1000", "name", "Alice")`
    2. set multiple field: `r.hset("user:1000",mapping={"name":"alice","age":30})`
    3. get: `r.hget("user:1000","name")`
        - only return value
    4. get all: `r.hgetall("user:1000")`
        - return both key and value

6. Sorted Set (ZSet) Operation
    1. add: `r.zadd("leaderBoard",{'Alice':100,"Bob":200})`
    2. get ranked item: `r.zrange("leaderBoard",0,-1, withscores=True)`
    3. increase score: `r.zincrby("leaderBoard",10,"Alice")`
    4. get score: `r.zscore("leaderBoard","Alice")`

7. Key Management
    1. check if exist: `r.exists("key")`
    2. set TTL: `r.expire("key",60)`
    3. get remaining time: `r.ttl("key")`
    4. rename key: `r.rename("old","new")`
    5. search keys: `r.keys("*")` or `r.keys("user*")` # wild card search
        - `r.keys("user:*")` to get all key related to user like `user:100`, `user:1`

8. Transaction ( for atomic operation )
    > with r.pipeline() as pipe:  
    >   pipe.set("x",1)  
    >   pipe.incr("x")  
    >   pipe.execute()  
    1. prevent race condition: 
        - `r.watch("balance")` watch "balance" and abort if it was changed, to prvent concurrent update
    2. execute: `r.exec()`
    3. update all comand at the same time with: `r.multi()`, then command, then `r.exec()`
        > r.multi()  
        > r.set("balance", 100)  
        > r.incr("balance", 50)  
        > r.exec()  

9. Pub/Sub (Message queue)
    1. publish: `r.publish("channel1", "hello")`
    2. create pubsub object: `p = r.pubsub()`
    3. subscribe: `p.subscribe("channel1")`
        - listen after subcribe: `for m in p.listen(): print(m)`
            - keep running until manually terminate by ctrl+c or break loop
    4. unsubscribe: `p.unsubscribe("channel1")`
    5. pattern-based subcribe: `p.psubscribe("user:*")`
        - subscribe to all user:* channel
    6. check subcription status:   `p.subscription_count`

10. HyperLogLog (approx unique count)
    1. add: `r.pfadd("visitor","user1","user2","user3")`
    2. count: `r.pfcount("visitor")`

11. when get, need to use decode() as redis stored binary data
    1. `data={k.decode(): v.decode() for k,v in r.hgetall("user:1000").items()}`
    2. need:
        - String: `SET`, `GET`
        - Hashes: `HSET`, `HGET`
        - Lists: `LPUSH`, `LPOP`
        - SortedSet: `ZADD`, `ZRANGE`
    3. no need:
        - Pub/Sub: `Publish`, `Subscribe` alr in string format

12. Redis namespace style
    - use `:` to seperate different objects of sametype
    - example:
        - user:1000:name, user:1001:name
        - order:500, order:501

13. value returned afer update or add
    - SET: return `"OK"` if success ( only fail when redis down)
    - HSET: return `1` if new, `0` if update
    - LPUSH: length after insert
    - ZADD: number of new elements added
        - `r.zadd("leaderBoard", {"Alice": 100, "Bob": 200}) ` return 2
        - `r.zadd("leaderBoard", {"Alice": 150, "Charlie": 250})` return 1, as only Charlie is new

14. Redis subsystem
    - key and channel is differnt subsystem: we can have same name key and channel without issue
    1. key-value store
        - `string`, `list`, `sortedset`, `hash`
    2. Pub/sub
        - not persistent
    3. Stream
        - lightweight version of kafka
        - store persistent logs
        - support multiple consumer
    4. Transaction
    5. Lua Script
        - execute multiple command as single unit
        - create a scirpt, then use`r.eval(script,1,'mykey')`
            - script: LUA script
            - 1: 1 key is passed
    6. Geospatial
        - `r.geoadd("restaurants", (103.851959, 1.290270, "BurgerPlace"))`
        - `r.geodist("restaurants", "BurgerPlace", "AnotherPlace", unit="km")`
    7. HyperLogLog
        - estimate unique count, use less memory than set
    8. Bitmap
        - store `0` and `1`
        - mark usr active at index 10 : `r.setbit("user_activity", 10, 1)`
        - get bit: `r.getbit("user_activity", 10)`
    9. Bloom Filter
        - check if item exist ( with small error)

15. Redis Stream
    1. stored ordered event with timestamps
    2. like message queue but persists data, histrory and consumer group
    3. add message: `r.xadd("mystream", {"user": "Alice", "message": "Hello, World!"})`
    4. read message: `r.xrange("mystream", "-", "+")`
        - example result: `[('1683145697263-0', {'user': 'Alice', 'message': 'Hello, World!'})]`
    5. create consumer group: `r.xgroup_create("mystream", "mygroup", id="0", mkstream=True)`
        - error if alr created
    6. consume message: `messages = r.xreadgroup("mygroup", "worker1", streams={"mystream": ">"}, count=1)`
        -`>`: read only new message assigned to consumer
    7. message will be PEL until ack
        - `r.xack("mystream", "mygroup", "1711638329173-0")`
    8. XREAD: not consume messgae, XREADGROUP: consume after XACK

16. Bloom Filter
    - check if item exist ( with small error)
    - add: `r.bfadd("seen_users", "user123")`
    - check: `r.bfexists("seen_users", "user123")`
    - do not support remove, can only add
    - memory usage lower than hash
    - speed faster than hash
    - may have false positive ( no for hash)
    - but not support deletion
    - use for
        - spam detection
        - url shorten
        - caching
        - distributed system to reduce common queries
        - wbe crawler
    - example
        > from redisbloom import BloomFilter  
        > bf = BloomFilter(r, 'user_bloom')  
        > bf.add("user:1000")  
        > print(bf.exists("user:1000"))
