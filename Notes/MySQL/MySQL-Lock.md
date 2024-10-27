# Mysql Lock
### source: xiaolincoding
https://xiaolincoding.com/mysql

1. type of lock in MySQL based on scope:
    1. global lock
    2. table-level lock
    3. row-level lock

2. global level lock:
    - `flush tables with read lock;`
    - whole database become read only
    - block all others like `insert`, `delete`, `update`, `alter table`, `drop table`
    - to release: `unlock tables`
    - ususally used for full datbase backup, to avoid backup inconsistent btw tables
        - like balance table adjusted, but sold table didnt update etc
    - problem:
        - block operation! 
    - to allow backup and non-block operation, can begin transaction before backup
        - readview & MVCC help to ensure consistent
        - when backup with mysqldump, add `-single-transaction` parametr, it will begin transacton before backup
            - *if the storage engine support repeatle read*
                - like innoDB
                - not like MyISAM (which must use global lock when backup)

3. table level lock ( 4 types ):
    1. table lock
        - `lock tables table_name read;`
        - lock the table for all thread, include the locker itself
        - to release: `unlock tables`
        - avoid table lock for innoDB, we have more fine-grained option
    2. metadata lock (MDL)
        - auto executed
        - when CRUD: 
            - acquire MDL read lock
            - other thread: allow read, but block structural change
        - when change table structure: 
            - acquire MDL write lock
            - block all other thred operation
        - will only release after transaction done & commited
        - if a thread cannot acquire lock ( which taken by previous, it will be blocked)
            - if too mahe, can cause thread overloaded!
            - write lock priority > read lock priority
            - if thread witing write lock, all after read will be block
    3. intent lock
        - before apply `shared lock`, need apply `intention shared lock` first
        - before apply `exclusive lock`, need apply `intention exclusive lock` first
        - when create/update/delete, will add `intention exclusive lock` first
            - select wont lock by default, just use MVCC to read consistently
                - share lock: `select ... lock in share mode;`
                - exclusive lock: `select ... for update;`
        - intent lock is table level
            - not conflict with shared lock & exclusive lock (which is row level)
            - only conflict with lock table read & lock table write
            - table lock & row lock:
                - read read share
                - read write exlcude
                - write write exclude
            - intent lock help to check which record in table is locked
                - without intent lock, need iterate 1 by 1
    4. auto-inc lock
        - primary index are usually auto-increment
            - due to `AUTO_INCREMENT`
        - when insert, no need to specify primary index, DB will auto assign value through auto-inc lock
        - when insert, lock by table level auto-inc lock, only release after whole row done.
            - block other insert
        - new ver: ligth weight lock
            - released after assigned the value, not waiting for whole row done
        - can be config using `innodb_autoinc_lock_mode`
            - 0: auto-inc lock
            - 2: light weight lock
            - 1: 
                - for `insert`,light weight
                - for `insert ... select`, bulk insert, auto-inc lock
        - light weight lock is best for efficiecny, but can cause inconsistency for master slave!
            - session A and session B trying to insert 4 rows to table C
            - session A insert 1 row, then session B take over and insert 1 row and repeat
            - session A & B inserted row id not continuos!
            - but in binlog, it only record the 2 insert command in session! (id `binglog_format` is `statment`)
                - will be 1 after another, not in the same time
            - when it execute in slave, it execute sequentially, so row created will be cluster together
            - master slave different!!!
            - to solve:
                - change `bing_format` to `row`
                - then bin_log wil record the changes in row in master and apply to slave!!

4. row-level lock
    - innoDB suppoer, MyISAM does not
    - normal `select` wont lock, it is snapshot
    - if want to lock when select:
        - share lock: `select ... lock in share mode;`
        - exclusive lock: `select ... for update;`
    - row level lock auto released after transaction done
        - transaction initiate by
            - `begin`
            - `start transaction`
            - `set autocommit = 0`
        - share lock ( S lock ) & exclusive lock (X lock) interaction
            |   |       X|       S|
            |---|--------|--------|
            | X |not comp|not comp|
            | S |not comp|comp    |
    - row-level lock type:
        1. Record Lock
            - X lock & S lock
                - when trans applied S lock, other trans can apply S lock, but not X lock
                - when transa applied X lock, other can not apply any lock
        2. Gap Lock
            - in repeatable read
            - specify a range to lock, others cannot insert
            - gap lock compatable with each other!
        3. Next-Key Lock
            - Gap + record lock
            - specify a range that others cannot edit or insert
            - X & S type, if taken X type, other's S/X will be blocked on the same range
        4. insert intention lock
            - not intent lock!!
            - happen when want to insert, but blocked by gap or next-key lock
            - the lock will be in pending status untill previous lock released