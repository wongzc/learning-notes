# Mysql basic
### source: xiaolincoding
https://xiaolincoding.com/mysql

0. 2 main layers in mySQL:
    1. server layer
        - incharge of connection, parse and execute SQL
        - other built in function like: date, time, math, encrytion
        - cross storage engine feature: stored procedures, triggers, views
        - component:
            - connector
            - query cache
            - parser
            - preprocessor
            - optimizer
            - executor

    2. storage engine layer
        - incharge od data storage and retrieval
        - support storage engine: innodb, MyISAM, Memory
            - index data structure implement by storage engine layer
            - different engine different search type.
                - innoDB: B+ tree

1. connector
    - connect through linux: `mysql -h$ip -u$user -p`
        -  can skip `-h` if connect to local
        -  `-p` is password, suggest to skip in command and enter in interactive dialogue
    - then 3 handshake of TCP, as mysql use TCP for transport
    - if account & password correct, connector will receive the user acess & keep it
        - so after establish connection, any access change wont impact user, unless user reconnect 
    - `show processlist;` cehck how many client connected
        - command `sleep` means never execute any command ( idle time in `time`)
    - `show variables like 'wait_timeout';` to check max timeour ( default 28800)
    - `kill connection +<id>` to disconnect
        - disconnected user only will know after receive error in next query
    - `show variables like 'max_connections';` to show max connection number
    - 2 type of connection:
        - one-time connection (short)
            - use less memory
        - persistent connection (long)
            - no need to keep establish connection
            - but high load to memory. to solve:
                - regularly disconnect long connection
                - client actively reset connection
                    - client use `mysql_reset_connection()` to reset connection after heavy operation.
                        - session variable, temporary table will be removed.
                    -  auto reconnect.

2. query cache ( server layer )
    - after connection establish, user can send sql command
    - server will analyze sql command, if first word is `select`, `will check query cache.
    - cache key is sql statement, value id result. if cache miss, will query table and write to cache later.
    - for table that frequent renew, cache hit rate is low. big table that in cache can be clear out easily and didnt use for once.
    - mySQL8.0 removed cache.
    - mySQL <8.0 need to change `query_cache_type` to `demand` to remove cache

3. parse command
    - parse SQL statement from user, identify keyword (like select, from)
    - check if syntax correct, then build a syntax tree

4. execute SQL
    - every `select` can be breakdown into 3 stage:
        1. prepare: preprocessor
            - check if table of column exist
            - expand `*` to all column
            - if table doesnt exist, error will be thrown here
        2. optimize: optimiser
            - to decide to use which index ( if there are multiple)
            - can use `explain` before statement to check using which index
                - `null` means not using index
            - covering index: query on primary index only ( or on secondary, but result need primary index only, which skip lookup)
            - secondary index: in B+ tree, leaf node  of secondary index store main index
        3. execute: executor
            - executor interact with storage engine
            - 3 ways:
                1. primary key index query `select * from product where id = 1;`
                    - main id is unique, so optimizer decide to use const as access type
                    - executor use `read_first_record` to point to innonDB query API, and pass id=1 to it, to get the first record
                    - storage engine use B+ tree to find the first id=1, return error if cant find
                    - executor check if result fulfill user requirement and send it. skip if it doesnt.
                    - executor is a while loop, will check again, but this time use `use_record`. access type is const, so will point to -1 and end.
                2. full table scan `select * from product where name = 'iphone';`
                    - optimizer decide to use all as access type as doesnt involve index
                    - use `read_first_record` to point to innonDB scan all api, storage engine return first record.
                    - executor check and skip if it is not correct, and send to client directly
                        - it actual send data row by row to client, by client side only show table when it completed.
                    - while loop, executor use `read_record`, point to innoDB scan all API, ask innoDB to return the next record.
                    - excutor check, and send to client if ok
                    - repeat untill executor receive query finish from storage engine
                3. index pushdown `select * from t_user where age > 20 and reward = 100000;`
                    - assume age and reward is composite index
                    - conposite index:
                        - index that create on multiple columns of a table
                        - `CREATE INDEX idx_name_age ON employees (name, age);`
                        - composite index stop after `>` or `<`
                            - so this case, age can composite index, but reward cant
                    - step without index pushdown:
                        - storage engine get the first record that fullfil the secondary index (age>20), and **lookup**, send full row to executor
                        - executor check if reward =100000, if not then skip, else send to client
                        - executor ask next record from storage engine, untill all completed.
                        - need many lookup!
                    - with index pushdown:
                        - after storage engine find a record that fullfil secondary index, it check if the `reward` column = 100000, if yes, lookup and return data to executor, else skip.
                        - executor check
                        - untill all done
                            - extra mention: "using index condition" means using index lookdown 

                    
5. where MySQL save data to?
    -  `show variables like data dir;`
    - each database will create a folder in this dir
        - each table will be a .frm, .ibd & dp.opt file in the folder
            - db.opt: store default character set and collation rule for database
            - tabelname.frm: table structure & meta data
            - tablename.ibd: 
                - if it is named with `ibdata1`: share tablespace, multiple table cross database can store in same file. ( but maintain own .frm file)
                - `tablename.ibd` means dedicated tablesapce. when `innodb_file_per_table` set to 1 ( default)
        - sturcture of .frm file, consists of:
            1. segemnt
            2. extent
            3. page
            4. row

1. segment
    - segment consist of extent, and multiple segemnt form table.
    - segemnt have: 
        1. data segment: store leaf extent of B+ tree
        2. index segment: store non leaf extent of B+ tree
        3. rollback segment: store rollback data
            - for MVCC (Multi-Version Concurrency Control) 
2. extent
    - innoDB use B+ tree, each layer in B+ tree connected with doubly linked list.
    - if use page as unit, the two adjacent page can be dar physically. when disk lookup, will have many random I/O, which is slow.
        - random I/O: disk need to move read/write head randomly to "seek"
    - to solve, use extent to group page, size ~1 MB, which is 64 page. then can sequential I/O.
3. page
    - data saved in row, but database read data in page, to improve efficiency
    - default page size 16KB ( continuos space)
    - smallest unit for innodb (for read to/ write from memory )
    - many type of page
        - data pages
        - undo log pages
        - overflow pages
4. row
    - row_format:
        1. Redundant: 
            - old format, no one use
        2. Compact: 
            - design to save more data in 1 page, default for 5.1
            - complete record has 2 part:
                ([length list][null list][record header][row id][trx id][roll ptr][actual data])
                1. additional infomration
                    - variable-length field length list
                        - char is fixed length, varchar is variable length
                        - need to save actual length of varchar, to read it
                        - data length save inversely
                            - row: ['a','abc','abcd'] will be save as 04,03,01
                            - `NULL` wont be save in actual data and so do length list
                        - so, left side is inversely saved data length, rght side is actual data, middle is the record header information.
                        - record header information point to location between next record header information and actual data, so read to left, we can get record header information, read to right, can get actual data.
                        - also the front data and length info more likely cache hit
                        - if all int, wont have length list
                    - NULL value list
                        - inversely save, 0 means not null, 1 means null
                        - then saved as hexadecimal
                        - if column set as `NOT NULL`, wont have null list, save 1 byte atleast
                            (1 byte for 8 record, 9 will be 2 byte)
                    - record header information ( 5 byte)
                        - delete_mask: mark if this row is deleted. (1 means deleted)
                        - next_record: point to next record, btween record header information and actual data, to left get record headre, to right get actual data
                        - record_type:
                            - 0: nomral record
                            - 1: non-leaf in B+ tree
                            - 2: minimum record
                            - 3: maximum record
                2. actual data
                    - row_id ( 6 byte )
                        - if specify primary key when create table, will not have row_id
                        - else innodb will add this in
                    - trx_id ( 6 byte )
                        - indicate which transaction generated this record
                    - roll_ptr ( 7 byte )
                        - pointer to last version
        3. Dynamic: default now, similar to compact
        4. Compressed: similar to compact



5. max 65535 for each row in mysql
    - exclude `TEXT`, `BLOBs`, Hidden column, record header information
        - hidden column: row_id, trx_id, meta data purpose
    - include storage overhead: length list, null list

6. varchar(n), max of n?
    - n is number of character, not byte
    - 1 row include null list & length list
    - null list, if <8 columns, 1 byte/row
    - lengthlist, if varchar<=255, 1 byte, else 2 byte/row
    - so... if 1 column, max n is 65535-2-1=65532
        - if character set is ascii
    -```CREATE TABLE test ( `name` VARCHAR(65532) NULL ) ENGINE = InnoDB DEFAULT CHARACTER SET = ascii ROW_FORMAT = COMPACT;```
    - if utf-8, 1 char is 3 byte, so 65532/3=21844
    - formula: sum of all varchar n + sum of [1 if n<=255 else 2 for n in varchar] + null list ( 2 if >8 column else 1)

7. row overflow
    - unit for mySQL interaction with memory is page, 16KB, 16384 byte
    - TEXT, BLOB can save more than 65535, which cause overflow
    - compact: if 1 page not enough, innodb save some data in actual table, and save extra data into overflow page, and actual data use 20 byte to point to address of the this extra data.
    - compress & dynamic: only 20 byte address in actual table, all data in oveflow page

