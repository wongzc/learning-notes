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
            - covering index: query on index only
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
            1. 


