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
6. 2 stage submit
7. optimize disk I/O for MySQL