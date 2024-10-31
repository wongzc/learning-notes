# Mysql Log
### source: xiaolincoding
https://xiaolincoding.com/mysql

1. what happen when we run update ? `UPDATE t_user SET name = 'xiaolin' WHERE id = 1;`
    - client establish connection through connector
    - contain `update`, so wont pass through cache, and involve `set`, will clear cache
    - analyzer differentiatie column name, key word etc, build syntax tree, check if syntax correct
    - preprocessor check if syntax exist
    - optimizer decide which index to use ( this case, id)
    - executor find and update data
    - involve 3 log: 
        1. undo log: 
            - in innodb, for transaction atomicity & MVCC
        2. redo log
            - in innodb, for persistency, restore when system failure
        3. bin log
            - in server, backup & master slave replication

2. undo log
    - eventhough no `begin` and `commit`, MySQL will start a hidden transaction when update/insert/delete
        - and auto commit (depends on confi `autocommit`)
    - undo log recorded the information requied to undo, to ensure atomicity
    - for example:
        - insert: record primary index, if undo, just delete the record based on index
        - delete: if undo, just reinsert
        - update: if undo, just roll back old value
    - undo log have:
        - `roll_pointer`: point to previous record
        - `trx_id`: which transaction updated the record
    - undo log + ReadView for MVCC
        - read commit: each select create a new Read View
        - repeatable read: create a ReadView when transaction started
    - so undo log 2 usage:
        - rollback to ensure atomicity
        - for MVCC
    - undolog write to disk through redolog, undo page in buffer pool, changes to undo page will be record in redo log
        - redo log write to disk every seconds/ when commit transaction

3. buffer bool
    - when record update, need to read data from disk , and update it in memory
    - after update done, cache it in buffer pool
    - when read:
        - if data in buffer pool. read from buffer pool
    - when update:
        - if data in buffer pool, update it in buffer pool
        - mark the page as dirty
        - wont immediate write to disk ( to improve I/O)
    - when MySQL start, innoDB request a continuos memory for buffer pool
    - and the memory divided into 16KB size each, which is page (cache page) 
        - when MySQL run, data from disk gradually load into cache page
        - so virtual memory high and physical memory low when DB start
        - physical memory only allocated when the virtual page accessed and triggered page fault, and establish virtual to physical address
    - consist of
        1. data page
        2. index page
        3. undo page
            - when trsnaction started, if it is update operation, need to record old value by generate undo log
            - undo log will be write into undo page
        4. insert buffer
        5. adaptive hash index
        6. lock information
    - when query data, whole page will be cache, and search by index page

4. redo log
    - buffer pool is in memory, data may lost if system fail
    - when update, innoDB update buffer pool and mark it as dirty page, and write the change as redo log
    - Write-Ahead Logging (WAL): 
        - technique that innoDB write change to log, then update dirty page in buffer pool to disk
    - when system fail, just based on redo log to create the latest state
    - if undo log updated, need to record in redo log as well
        - undo log: to rollback if transaction not done
        - redo log: to restore if transaction done
    - redo log + WAL, innoDB is crash-safe (persistency)
    - why need redo log write to disk when data is also write to disk?
        - redo log by sequence, faster (sequential write)
        - data update need to identify location then write, slower (random write)
    - redo log write immediate to disk?
        - redo log write to redo log buffer
        - redo log buffer default size 16MB (control by `innodb_log_Buffer_size`)
        - when data write from redo log buffer to disk?
            - MySQL shutdown
            - data write into redo log buffer > 50% of total volume
            - innoDB background process, every seconds
            - everytime when transaction commited, controlled by `innodb_flush_log_at_trx_commit`
    - `innodb_flush_log_at_trx_commit` control when transaction commited
        - 0: keep redo log in redo log buffer
        - 1: write redo log from redo log buffer to disk. default value
        - 2: write redo log from redo log buffer to redo log file ( not disk, it is in os page cache)
    - innoDB background process
        - for 0 cases: use `write()` to write to os page cache, use `fsync()` to write to disk.
            - lost data when MySQL or OS fail
        - for 2 cases: `fsync()` to write data from page cache to disk.
            - lost data when OS fail
    - comparison:
        - efficiency: 0>2>1
        - safety: 1>2>0
    - redo log full
        - innoDB redo log group: consits of 2 redo log files `ib_logfile0` and `ib_logfile`
            - same size for the 2
        - circular write: when reach end, go to beginning again
            - write 0 first, then 1, then go back to 0
        - redo log is to avoid dirty page lost ( in memory), so when dirty page write to disk, redo log no use
            - `write pos`: redo log current write location
            - `checkpoint`: to erase location ( after write to disk)
            - if write pos meet checkpoint, means redo log full, MySQL will be blocked
                - will stop and update buffer pool data to disk
                - important to set correct size for redo log!
5. bin log
    - record all db structure change or table data change, not record query
    - generate by server layer everytime when update done
    - write to bin log file when transaction completed
    - why do we need redo log when we have bin log?
        - MySQL at first use MyISAM, which not crash-safe
        - binlog can only used for archiving
        - innoDb then introduced and redo log added for crash-safe
    - bin log vs redo log:
        1. purpose:
            - binlog created by MySQL server layer, all storage engine can use
            - redo log created by innoDB
        2. format:
            - binlog have 3 types:
                1. STATEMENT
                    - record SQL statement that ran on master, then ran on slave.
                    - problem with dynamic function like uuid, may cause difference btw master & slave
                    - called logic log
                2. ROW
                    - record how's the change done on data
                    - may cause bin log size super huge ( when STATEMENT maybe just 1 sentence)
                3. MIXED
                    - use ROW or STATEMENT depends on situation
            - redo log
                - physical diary of what change done on which table, which page
        3. write in method:
            - binlog is append, if full, create another file
            - redo log is circular, fixed space
        4. usage:
            - binlog for backup, master-slave replication
                - if DB deleted, can use binlog, but not redo log
                - cos redo log is not full record
            - redo log for system crash-safe
    - master slave replication 
        - asynchronously calling statement in binlog ( from master ) to slave
        - 3 step:
            1. write to binlog
                - after client sent order, master write to binlog first
                - then commit transaction & update local data
            2. sync binlog
                - slave create a I/O thread connect to master log dump thread to receive master binlog
                - copy binlog to all slave, slave write binlog to relay log
                - then reply master "copy success"
            3. replay binlog
                - replay binlog and update data
        - more slave, master need more I/O thread for log dump
            - usually master with 2-3 slave
        - master slave replication model: 3 type
            1. Synchronous:
                - master transaction thread need to wait slave complete bin log sync
                - low efficiency, slow
            2. Asynchronous:
                - master transaction thread dont wait for bin log to finish sync to slave before return to client
                - if master crash, data loss
            3. Semi-Synchronous:
                - master transaction thread wait for some slave complete dupliacte bin log
    - when bin log write to disk?
        - during transaction, binlog write to binlog cache
        - when transaction commited, write from binlog cache to binlog file
            - in page cache, not in disk, need `fsync` to disk
        - binlog can not be splitted. due to 1 thread 1 transaction & atomicity of transaction
            - must be write to disk in 1 operation
        - binlog cache: memory that every thread have for binlog caching
            - size depends on config `binlog_cache_size`
        - `sync_binlog` control when transaction commited
            - 0: write to page cache, os to decide fsync. default
                - faster, but risky
            - 1: write and fsync
            - N: write N time before 1 fsync

6. full flow for update: `UPDATE t_user SET name = 'xiaolin' WHERE id = 1;`
    1. optimizer optimize
    2. executor use innoDB API to check if data in buffer pool or need to read from disk to buffer pool
    3. executor get clustered index and compare before and after change
    4. innoDB record undolog, write to undo page in buffer pool
    5. innoDB start to update and mark as dirty page in memory, write to redo log
    6. background thread to decide when to write to disk ( WAL)
    7. record binlog and save to binlog cache
    8. 2 stage submit
7. 2 stage submit
    - after transaction commited, both redo log & bin log need to save to disk.
        - if redo log success and bin log fail: master updated, slave outdated
        - if redo log fail and bin log success: master outdated, slave updated
    - 2 stage submmit:
        - to makesure atomicity of redo log & bin log
        - using internal XA transaction
        1. prepare stage
            - write XID ( ID of XA transaction) to redo log
            - set redo log transaction status as prepare
            - save redo log to disk (if `innodb_flush_log_at_trx_commit` = 1)
        2. commit stage
            - write XID to binlog
            - save binlog to disk (if `sync_binlog` = 1)
            - server API to set redo log status as commit
    - fail during 2 stage commit
        - MySQL scan redlog sequentially, if any redo log in prepare stage, cehck if XID in binlog
        - XID not in binlog: fail before binlog saved, rollback
        - XID in binlog: fail after binlog saved, commit
    - uncommited transaction redo log also saved in disk, but will rollback if system fail & restart
        - binlog only save to disk when commited
    - problem of 2 stage commit:
        - high disk I/O
            - binlog save in binlog cache
            - redo log save in redo log buffer
            - memory to disk control by config, if both 1, every commit means 2 I/O
            - impact to efficiency
        - high lock contention
            - `prepare_commit_mutex`: transaction need to obtain this lock to move to `prepare` and release after `commit`
                - problem when multithread, causing competition
            - group commit
                - combined multiple commited transaction into 1
            - new stage:
                1. prepare
                    - in mySQL 5.7, transaction dont fsync redo log to disk seperately, do it in flush
                2. flush
                    - transaction sequentially write binlog to file ( not disk)
                    - in MySQL 5.7, fsync redo log to disk
                3. sync
                    - fsync binlog file to disk 
                    - frequency depends on `Binlog_group_commit_sync_delay` ( 1 fsync for multiple binlog changes)
                    - if more than `Binlog_group_commit_sync_no_delay_count` will auto save
                4. commit
                    - innoDB commit sequentially
8. optimize disk I/O for MySQL
    - `binlog_group_commit_sync_delay`: to delay binglog write to disk ( wont fail if MySQL crash, fail if OS crash)
    - `binlog_group_commit_sync_no_delay_count`: same as above
    - `sync_binlog`: set to mroe than 1 ( usually 100-1000), only fsync after N write for binlog ( risk for losing N binlog if OS crash)
    - `innodb_flush_log_at_trx_commit`: set to 2, only write to file in page cache, os to control save to disk ( risk for losing data if OS crash)