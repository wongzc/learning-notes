# Mysql Buffer Pool
### source: xiaolincoding
https://xiaolincoding.com/mysql

1. Buffer Pool
    - cache data that read from disk in innodb
    - update data in buffer pool and mark as dirty, background process to write back to disk
    - buffer pool size:
        - continuos memory, default 128MB
        - config by `innodb_buffer_pool_size`, set as 60~80% memory
    - cache page: 16kb size page in buffer pool
    - when MySQL start, virtual memory big, physical memory small, when virtua; memory visisted, system connect virtual address with physical address
    - buffer pool cache:
        - index page
        - data page
        - undo page
        - insert cache page
        - adaptive hash index page
        - lock infromation page
    - every cache page has a control block, control
        - cache page memory
        - cache page number
        - cache page address
        - caceh page linked node
    - fragemented space: unused memory in buffer pool
    - when query 1 record, whole page will be cache

2. Buffer Pool Management
    - Free list ( free linked list)
        - to help identify free cache page
        - when need cache page, check from free list and take it out
    - flush list
        - help identify dirty page that need to be write in disk ( after data update)
    - improve cache hit
        - LRU
            - when queried data in buffer pool, move it to head
            - else, add data to head, remove the old data in tail
            - problem with simple LRU
                - prefetch failure
                    - neighbnour page that fetch together into cache, but end up unused
                    - waste space
                - buffer pool pollution
                    - problem when large amount of data queried and flush out data that was in cache
                    - cause I/O when the flushed data required
            - how my SQL solve improve LRU
                - prefetch failure
                    - linked list split into young and old region
                    - control by `innodb_old_blocks_pct`, default 37 ( 63 young, 37 old)
                    - prefetch page will be put in old region, allow faster flush out
                        - move to young if read
                - buffer pool pollution
                    - make data harder to move from old to young
                    - when the data in old is visited, record the time
                    - when the data revisited, if it within `innodb_old_blocks_time` ( default 1000ms)
                        - will keep at old
                        - if larger than that, move to young
                    - keep data into young only when visited and stay at old >1s
                    - also top 1/4 of young wont be move to top again
    - when dirty page write to disk?
        - WAL: Write Ahead Log
            - write to file first then write to disk
        - dirty page write to disk when
            - redo log full
            - buffer pool space not enough
            - MySQL free
            - MySQL normal shut down
