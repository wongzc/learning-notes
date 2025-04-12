# Redis persistency study note
### source: xiaolincoding
https://xiaolincoding.com/redis/

1. Redis persist data to disk

2. AOF (Append-Only File)
    - log every write operation to a file so that data can be reconstructed after restart
        - write include: SET, DEL ( add, update, delete)
        - read doesn't log because it do not modify state of data
    - default is off, need to change redis.conf to allow
    - log example:  
        command:  
        > set name abcdefg  

        log:
        >*3  
        >$3  
        >set  
        >$4  
        >name  
        >$7  
        >abcdefg  

        - *3 means 3 part
        - $ means number of char
    - redis execute before log. 
        - pros:
            - avoid extra cost if user command syntax wrong
            - do not block operation by writing log ( write log is in the main process as well!)
        - cons:
            - data lost if system fail before writing to log
            - writing may block next operation
    - AOF strategy:
        - AOF step: 
            1. execute write
            2. command added to server.aof_buf
            3. write() to write aof_buf data to page cache
            4. kernel to write to disk
        - redis.conf, appendfsync can control step 4
            1. Always: after write, always sync to AOF file
                - less data lost, but block next write!
            2. Everysec: every second, write page cache to disk
                - more balance between data loss & blocking
            3. No: os to decide when to write from page_cache to disk
                - less blocking, but high chance data lost
        - data lost vs writing block next is opposite
            - cant fix both at the same time!
    - AOF rewrite
        - execute when AOf file larger than preset threshold
        - only log the latest value of a key (older will be discard)
        - when rewrite, will write to new file first, then use the new file to overwrite original file
            - why new file? to avoid damage by system failure when writing to orginal AOF
    - AOF background rewrite
        - rewrite is happen at background, not in the main process, as it is time consuming
        - subprocess " bgrewriteaof "
            - subprocess fork a page table from main process, which point to same memory address {page table: map virtual memory to physical address}
                - save memory! but this memory become read only
                - when main try to write to this memory, os will block and run copy on write ( main will use the new one)
                    - become 2 memory
        - not using thread, as multi thread share same memory, need to use lock to avoid conflict, reduce efficiency
        - 2 stage where may block main process
            1. when create sub process: need to copy page table from main process, larger table more time
            2. after create sub process: when main/sub modify share memory & trigger copy on write
        - what if rewrite AOF happening and key-value from main process changed?
            - when bgrewriteaof running, and at the same time where main process writing, it will:
                - execute command from client
                - log to AOF buffer
                - log to AOF rewrite buffer
            - so when bgrewriteaof done, it will signal main process
            - main process will then add AOF rewrite buffer to new AOF file, so rewrited and orginal AOF became the same

    
3. RDB (Redis Database Backup)
    - create snapshot of datasets at specific interval
    - faster than AOF

    - 2 command to create RDB:
        - save: run in main process, may block if writing time too long, but dont need fork, impact redis perfrm
        - bgsave: run in subprocess, avoid blocking main process, need fork, no redis pefrm impact 
        - default setting for save:  
            > save 900 1  
            > save 300 10  
            > save 60 10000  
            - 1 change in 900 seconds
            - 10 change in 300 seconds
            - 10000 change in 60 seconds
            - any of it meet will trigger RDB
        - full snapshot: all data in memory save to disk
            - usually set at 5 mins interval to avoid affecting redis perfromance as writing full snapshot data
            - data lost may be more than AOF
        - when bgsave execute:
            - fork will create subprocess and duplicate the page table, memory address of main & sub point to save physical memory
            - similarly, copy on write will be execute to create a new memory when main/sub trying to change the memory
                - so... data changes by main process when bgsave running wont be save in this bgsave, onyl get save in next bgsave
                - also, memory mey become 2 time larger during bgsave, if all the value being updated and having copy on write

4. mix AOF and RDB
    - change this setting to yes  
        `aof-use-rdb-preamble yes`
    - faster due to writing in RDB format, smaller size also, also RDB format load faster to redis
    - binary, so harder debug by human! also may harder to manual recover from broken AOF
    - when AOF rewrite:
        - subprocesss created by fork will use RDB to write the share memory into AOF first
        - new command added during rewrite will be write to AOF rewrite buffer, where subprocess will pick up and write in using AOF way
        - then inform main process to use the new file to replace orginal AOF

5. impact of large key to persistency
    - affect AOF fsync(), if set to "always", will block main process as take more time to write to disk
        - if "no" will not execute fsync()
        - if "everysec" , will run in async, not blocking

    - AOF log file will be fill up in short time, trigger AOF rewrite
        - more often fork, block main process
        - if copy on write execute, will take more time as well

    - client side time out: large key take more time to load
    - network blocking: generate significant network load for large key
    - thread blocking: `del` command take block thread when deleting large key
    - uneven memory distribution: large key consume more memory, may have data & query skew in a evenly distributed slots.

    - to avoid large key:
        - break large into small during design
        - regularly check, avoid use `del` to delete large key ( which block main process), use `unlink` to delete in async manner
    