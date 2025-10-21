# Mysql Transaction
### source: xiaolincoding
https://xiaolincoding.com/mysql

1. what is transaction
    - support by innoDB
    - not by MyISAM
    - 4 requirement of transaction:
        1. atomicity
            - for list of tasks, either all success or all fail. cannot stop at middle
        2. consistency
            - before and after execution, data is consistent, valid state
            - sum of money is the same, none-negative value etc
        3. isolation
            - concurrent transaction not interfere each other
            - intermediate state of each transaction not visible to others
        4. durability
            - changes are permanent after each transaction commited
    - how innoDB achieve these?
        - atomicity: undo log
        - consistency: atomicity + isolation + durability 
        - isolation: MVCC or lock ( multi version concurrent control)
        - durability: redo log

2. what issue happen when concurrent?
    - MySQL allow multiple client connect at the same time, means handle multiple transaction
    - problem:
        1. dirty read [D]
            - happen when transaction A read a non-commited data from another transaction B
            - which if B rollback, A get wrong data
        2. non-repeatable read [NRR]
            - two read of the same item in 1 transaction give different result
            - maybe due to item being updated by another transaction
            - reading same row twice and get different data
        3. phantom read [P]
            - in 1 transaction, if number of records matching differs between 2 queris
            - execute range query twice and get more/less rows ( inserted/deleted by another transaction)

3. transaction isolation
    - seriousness: 
        - dirty > non-repeatable > phantom
            - dirty: violate isolation and atomicity as read uncommited data
            - non-repeatable: violate isolation but only for existing row
            - Phantom read: violate isolation at set level, only bad when we depends on aggregated data
    - SQL created 4 type of isolation level to avoid:
        1. READ UNCOMMITTED: 
            - uncommitted transaction can be seen by other transaction
            - may have D,NRR,P
        2. READ COMMITTED: 
            - only committed transaction can be seen by other transaction
            - may have NRR, P
        3. REPEATABLE READ: 
            - default in MySQL innoDB
            - data that a transaction read is always same
                - no defense againt new/ delete data
                - but actually MySQL RR is good enough to protect Phantom read
            - may have P
        4. SERIALIZABLE
            - lock a record
            - if any conflict, the later need to wait for the earlier
            - all good
    - but in reality, MySQL repeatable read can avoid most phantom read
        - so no need serialize, which affect performance
        - for snapshot read: `SELECT`
            - using MVCC to solve phantom read
            - for new data that inserted, repeatble read will not show it!
        - for current read: `SELECT ... FOR UPDATE`
            - using next-key lock
            - when others trying to insert into range of next-key lock, will be blocked
    - how to implement the 4 isolation level?
        1. serialize: use a lock for read write
        2. repeatable read: create a `read view` ( snapshot) before the whole transaction start
            - readview + next key lock ( record+gap lock)
        3. read committed：create a `read view` before each command
        4. read uncommitted: always read latest data
    - 2 type of transaction start command:
        1. begin/start transaction
            - after this, transaction doesnt start.
            - only start after first `SELECT`
        2. start transaction with consistent snapshot
            - transaction start after this

4. read view
    - 4 columns in read view:
        1. creator_trx_id
            - id of transaction that create this read_view
        2. m_ids
            - list of active uncommited transaction in database when read view created
        3. min_trx_id
            - smallest in m_ids
        4. max_trx_id
            - transaction id for next new transaction when read view created
            - largest transaction id +1
    - 2 hidden column related to transaction in clustered index record  
        1. trx_id
            - when a transaction changed a clustered index record, it record the id in
        2. roll_pointer
            - when changed cluster index record, older ver will be written into undo log
            - pointer to undo log
    - so when transaction visit records:
        - trx_id < min_trx_id:
            - it committed before read view created, can be seen by current transaction
        - trx_id >= max_trx_id:
            - it created after read view, cant be seen by current transaction
            - need to use roll pointer to check down
        - min_trx_id <= trx_id < trx_id:
            - if trx_id in m_ids: 
                - havent commit when read view created, cant be seen
                - need to use roll pointer to check down
            - if not in m_ids:
                - committed when read view created, can be seen
        - this is call `MVCC`, control by version chain

5. repeatable read
    - create a read view when transaction started, and use the same read view through out transaction

6. read committed
    - create a read view every time read data

7. MySQL repeatable read completely solved phantom read?
    - snapshot read solved by MVCC
    - current read: `insert`, `delete`, `update`
    - next-key locks
        - when `select * from table where id>5 for update`
        - means id range from 5 to unlimited will be locked
            - any insert into this range will be blocked
    - but!
        - still can have phantom, when the new inserted value is larger than max?
    
