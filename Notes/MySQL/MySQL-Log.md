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
4. redo log
5. bin log
6. 2 stage submit
7. optimize disk I/O for MySQL