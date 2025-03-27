1. Connect to redis
    - r = redis.Redis(host='localhost', port=6379, db=0)

2. String Operation
    1. set: `r.set("key","value")`
    2. get: `r.get("key","value")`
    3. increase (must be numeric): `r.incr("counter")`
    4. decrease: `r.decr("counter")`
    5. Append: `r.append("key", " world") `
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
    4. get all: `r.hgetall("user:1000")`

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
    5. search keys: `r.keys("*")` # wild card search

8. Transaction ( for atomic operation )
    > with r.pipeline() as pipe:  
    >   pipe.set("x",1)  
    >   pipe.incr("x")  
    >   pipe.execute()  

9. Pub/Sub