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

5. How MySQL lock
    - which SQL statement use row-level lock?
        - innoDB, (MyISAM dont have)
        - normal select is snapshot, use MVCC
        - lock select that use row lock:
            - share lock: `select ... lock in share mode;`
            - exclusive lock: `select ... for update;`
        - update 
            - `update table .... where id = 1;`
            - X lock
        - delete 
            - `delete from table where id = 1;`
            - X lock
    - How MySQL lock?
        - for index: next-key lock
            - when can use gap-lock or record lock to prevent phantom record, next-key lock will become record/gap lock
        - if using secondary index, it lock both primary & secondary index
        1. unique index equal query `select * from user where id = 2 for update;`
            - if record exist
                - next-key lock become record lock
                - X-locked
                - use `select * from performance_schema.data_locks\G` to check which lock is used
                    - under `LOCK_MODE` ( which under `LOCK_TYPE`: `RECORD`, `TABLE` is for table level)
                        - `X`: next-key
                        - `X, REC_NOT_GAP`: record lock
                        - `X, GAP`: gap lock
                - primary-index is unique, and with record-lock, others can delete it, so count cannot be change.
                    - use record lock is able to solve phantom record
            - if record doesnt exist
                - gap lock between the last that meet criteria until the first that fail criteria
                    - if record is 1,5,6,7
                    - search 2 will lock (1,5), which is 2,3,4
                - why dont use next-key lock?
                    - it will lock 5, which we dont need
        2. unique index range query
            - use next-key lock for all index, and change to gap/record in below scenario
                - for `>=X`:
                    - if `X` exist:
                        1. record lock X
                        2. next-key lock `(X, largest record]`
                        3. next-key lock `(largest record, supremum pseudo-record]`
                    - if `X` doesnt exist:
                        1. next-key lock `(X, largest record]`
                        2. next-key lock `(largest record, supremum pseudo-record]`
                - for `>X`: 
                    1. next-key lock `(X, largest record]`
                    2. next-key lock `(largest record, supremum pseudo-record]`
                    - so no one can insert value after X
                - for `<X`:
                    - if `X` doesnt exist:
                        1. next-key lock `(infimum, smallest record]`
                        2. next-key lock `(smallest record, largest that <X]`
                        3. gap lock `(largest that <X, smallest that>X)`
                    - if `X` exist:
                        1. next-key lock `(infimum, smallest record]`
                        2. next-key lock `(smallest record, X]`
                        3. gap lock `(X, smallest that>X)`
                - for `<=X`:
                    - if `X` doesnt exist:
                        1. next-key lock `(infimum, smallest record]`
                        2. next-key lock `(smallest record, largest that <X]`
                        3. gap lock `(largest that <X, smallest that>X)`
                    - if `X` exist:
                        1. next-key lock `(infimum, smallest record]`
                        2. next-key lock `(smallest record, X]`

        3. non-unique index equal query
            - when query secondary index, lock applied to secondary index, and under certain criteria, for primary index
            - when value exist:
                - may have more records that equal, keep scanning until exceed
                1. next-key lock `(last value that smaller than X, X]`
                    - avoid insert X infront
                2. gap lock `(X, first value that larger than X)`
                    - avoid insert X from the back
                - record lock primary index
            - when value doesnt exist:
                - gap lock `(last that smaller than X, first that bigger than X)`
                - but value that is `last that smaller than X` or `first that bigger than X` can insert failed also
                    - `first that bigger than X`: when the insert primary id is smaller than primary id of `first that bigger than X`
                    - `last that smaller than X`: when the insert primary id is bigger than primary id of `last that smaller than X`
        4. non-unique index range query `select * from table where col >= X  for update;`
            - always next-key lock
            1. next-key lock `(last that smaller than X, X]`
            2. next-key lock `(X, largest value]`
            3. next-key lock `(largest value, supermum]`
            4. record lock all primary key that >=X
        5. no index query
            - next key lock all ! as it is full table scan


summary
<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E8%A1%8C%E7%BA%A7%E9%94%81/%E5%94%AF%E4%B8%80%E7%B4%A2%E5%BC%95%E5%8A%A0%E9%94%81%E6%B5%81%E7%A8%8B.jpeg">
<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E8%A1%8C%E7%BA%A7%E9%94%81/%E9%9D%9E%E5%94%AF%E4%B8%80%E7%B4%A2%E5%BC%95%E5%8A%A0%E9%94%81%E6%B5%81%E7%A8%8B.jpeg">
