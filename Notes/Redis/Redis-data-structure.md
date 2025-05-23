# Redis data structure study note
### source: xiaolincoding
https://xiaolincoding.com/redis/


1. redis underlying data structure
    1. simple dynamic strings (SDS)
            - SDS can save as binary data as it use len to decide if a string ended
            - SDS get len complecity is O(1) ( C string is O(n) as it didnt keep len)
            - SDS concate will not result in overflow, as it will check if space enough before execute
            - c string use \0 as the end of string, ( pointer point to start of string)
                - so c string cannot have \0, else will have problem
                - c string cannot store binary data like image, video etc
        - SDS structure
            - len
            - alloc ( allocated space size)
            - flags ( 5 type of sds, difference is len & alloc)
            - buf (actual data)
        - SDS struct byte alignment
            - using packed alignment, i.e., struct with char (1 byte) and int (4 byte), will be 5 byte ( instead of 8 byte)
    2. listNode ( not using anymore)
        - based on list node
            - prev, next, value ( double linked)
            - prev/next can point to null
        - head, tail, len

        - dup: create a copy
        - free: free memory
        - match: compare value
        - good:
            - move to front/ back is O(1)
            - move to head/ tail is O(1)
            - len is O(1)
        - bad:
            - each node is not continuos in memory, not efficient way of using CPU cache
                - CPU cache store frequent access data closer to CPU
                - if data continuos, high chance cache hit
            - each node have pointer to front/back, overhead for memory
        - so redis use ziplist when data small
    3. zip list ( not using anymore)
        - continuos space in memory, good use of CPU cache
        - consist of:
            - zlbytes: number of byte used by ziplist
            - zltail: pointer to last element in ziplist
            - zllen: number of element in ziplist
            - zlend: mark end of ziplist
            - entry:
                - prevlen: previous entry length, for back to front iteration
                    - prevlen <254, use 1 byte
                    - prevelen >=254, us 5 byte
                - encoding: data type and length 
                    - if data is int, use 1 byte to encode
                    - if data is string, use 1/2/5 byte encode, depends on length
                - data
        - chain refresh:
            - when insert element, if space not enough, need to find another space to accomodate the whole zip list
            - if newly inserted elemtn is big, may affect the  next element "prevlen" size, causing chain reaction down the list
        - bad:
            - cannot store too much data
            - when update, memory will relocated, can cause chain refresh
                - only use when data small
            - need to 1 by 1 check element ( except head/tail)
    4. hash table ( dictht )
        - use hash function to calculate key to find the location in table
        - hash table is actually array
        - each element in hash table is pointing to a dictEntry ( hash element)
            - each dictEntry consist of:
                - key value
                - pointer to value
                - pointer to next dictEntry
        - hash key conflict
            - happen when 2 different key, get the same value after hash function
            - redis solve it with linked hash list
                - if 2 (or more) hash key assigned same hash bucket, the later will be connected to the previous by the pointer
                - can be a problem if too long (O(n) speed) 
                    - use rehash to solve
        - rehash
            - in actual usage, redis use 2 dictht for hash
            - data will be first write to hash table 1 
            - when rehash:
                - assign space to hash table 2, usually 2x size of hash table 1
                - move data from ht1 to ht2
                - free up space in ht1, rename ht2 to ht1, create empty ht2 for nest rehash 
            - incremental rehash:
                - can be time-consuming when copy all data in ht1 to ht2
                - with incremental rehash, small number of element is copy over during deletion/ query/ update
                - insertion will be directly insert into ht2
                - so when check, it check both ht1 & ht2
            - when rehash happen:
                - load factor = number of element in hash table/ number of slot in hash table
                - when LF >1:
                    - rehash if no bgsave (RDB) or no bgrewriteaof (AOF rewrite)
                - when LF >5:
                    - must rehash
        - good:
            - O(1) speed for search
        - bad:
            - if data size too big, may have hash key conflict
    5. int set
        - continuos memory
        - consist of:
            - encoding: decide content element data size
            - length: number of element
            - content: int8 but data size depend on encoding
        - intset upgrade, when new element having larger size than current, say insert int32 into int16 set
            - will not assign new space
            - expand based on current, then assign the value to the new space 1 by 1
            - why inset upgrade:
                - save space when the array only have smaller int, only upgrade when big int added
            - after int set upgarde, wont downgrade
    6. skip list
        - only zset use skiplist 
            - in zset struct, actually we have skip list and hash table
        - when insert/update new data, data will be insert into both skip list and hash table
            - support range query, as we use skiplist
            - can get element score in O(1), as we use hash table
            - hash table is only use for retriving elemebt score in constant time in zset
            - the other will use skip list
        - skip list design
            - each skiplistnode keep:
                - element
                - score
                - pointer to front ( point backward)
                - skiplist level, which have:
                    - decide skip how many
            - skiplist node is under skiplist
                - skiplist contain pointer to head, tail, length and level used
            - keep each next level ratio at 2:1
                - example: 4 level 0, 2 level 1, 1 level 2
                - best performance
                - how redis implement this?
                    - randomgly create level when cerate node
                    - generate random number between 0-1, untill it is bigger than 0.25
                        - if <0.25, add 1 level
                    - max 64 level
        - why skip list but not tree?
            - less memory intensive: tree have 2 pointer ( l/r). skip list is 1/(1-p) pointer per node, default is 1.33 when 25%
            - cache locaility of zset is as good as tree, for zrange/ zaverage operation which need to travere the list
            - simpler to implement & debug
        - good:
            - support average O(logN) search
        
    7. quick list
        - is actually listnode+zip list
        - each node is a zip list
        - when insert, if space enough, just add, else, create a new node

    8. listpack
        - to mitigate chain reaction from ziplist
        - remove prevlen, the other similar to ziplist


2. data type to underlying data structure
    - String: SDS     
    - list:
        - quicklist
    - Hash:
        - listpack: small hash
        - hash table: large hash
    - Set:
        - int set: used if all element in set are integer and number<512 ( default) 
        - hash table: if contain non-int or >maxintset, use this
    - zset:
        - listpack: length <128>, element also <64bit
        - skip list+ hash table