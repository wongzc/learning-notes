# Mysql index
### source: xiaolincoding
https://xiaolincoding.com/mysql

1. what is index?
    - help DB to locate data faster
    - Classification by:
        - "Data Structure": 
            - B+ Tree Index
            - Hash Index
            - Full-text Index
        - "Physical Storage": 
            - Clustered Index (Primary Key Index)
            - Secondary Index (Auxiliary Index)
        - "Field Characteristics": 
            - Primary Key Index
            - Unique Index
            - General Index
            - Prefix Index
        - "Number of Columns": 
            - Single-column Index
            - Composite Index
2. By data structure
    |index type| InnoDB | MyISAM | Memory Engine |
    |----------|--------|--------|---------------|
    |B+ Tree   | Yes    | Yes    | Yes   |
    |Hash Index| No     | No     | Yes   |
    |Full-text | Yes    | Yes    | No    |
    - how innoDB select index when create table:
        1. use primary index as key
        2. else, use unique column that without NULL as key
        3. else, create a hidden incresing unique id as key
        - the other index will be secondary index
        - primary & secondary index default using B+ Tree
    - B+ tree:
        - multiway tree, data in leaf node, branch node store index
        - each parent node index will be in child node, so all index in leaf
        - leaf connect to each through doubly linked list
        <br>
        <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E7%B4%A2%E5%BC%95/btree.drawio.png" alt="Example Image" width="500">
        - each access to node is an I/O, only 3/4 I/O to get data (even for 10mil+)
        - very fast
    - B+ tree for secondary index
        - in primary index: leaf node store data
        - in secondary index: leaf node store primary index
        - process:
            - use B+ tree to find secondary index, get primary index
            - use B+ tree to get primary index, to return full data
                - called "lookup", use 2 B+ tree
            - if use secondary index to get data that already in B+ tree, (i.e., primary index value), no lookup!
                - called covered index, 1 B+ tree can get result
    - compare to other tree:
        - B tree:
            - B tree store data in branch node 
                - B+ store in only leaf, data in each node smaller
                - 1 I/O can access more data
            - B+ tree is double linked, suite for range search
        - Binary Tree
            - complexity: O(d*logn), d= max depth, n= # of node
            - B+ tree, mx height always 3 or 4, faster
        - Hash
            - fast for exact value match
            - but not suit for range search

3. By Physical storage
    - Primary Key Index
    - Secondary Index

4. By Field Characteristics
    - Primary Key Index
        - only allow 1 primary key index, cant be null
        - create by adding `PRIMARY KEY (index_column_1) USING BTREE` when create table
    - Unique Index
        - use unqiue field to index, table can have multiple unique, allow `NULL`
        - create with `UNIQUE KEY(index_column_1,index_column_2,...) `
    - General Index
        - column that used as index, no need to be unqiue or primary key
        - cerate with `INDEX(index_col1, index_col2)`
    - Prefix Index
        - set for char, varchar, binary, varbinary
        - use the first few word to index
        - use the same way like create generate index, but need to define length
            `INDEX(column_name(length))`




5. by number of column
    - single column: 
        - like primary key index
    - composite index
        - index that create with multiple column
        `CREATE INDEX index_name ON table_name(index_a, index_b, index_c);`
        - B+Tree leaft node will have all index a,b,c
        - when search, use leftmost matching principle
            - if serach conditon involve a,a+b, a+b+c:
                - use composite index ( left most is a)
                - can be c+b+a, optimizer will help to fix
            - if search b or b+c or c:
                - will not use composite index
                - because the composite sort by a, but we dont use a, while b,c is partially sort
        - composite search range
            - when search, maybe some use composite index, some use B+ tree, some doesnt use
            - happen when range search
            - maybe `select * from table when a>1, b=3`, assume a,b composite index
                - will composite search until not match for a, which is a>=2
                - and then it stop and exit, so b cant use!
            - but if `select * from table when a>=1, b=3`
                - will composite for both a,b
                - until doesnt match a>=1
            - for `SELECT * FROM t_table WHERE a BETWEEN 2 AND 8 AND b = 2`
                - in mySQL, between is >= and <=, so include 2 & 8
                - thus, both a & b will be compsite search!
            - for `SELECT * FROM t_user WHERE name like 'j%' and age = 22`
                - use composite serach for both name & age
                - if varchar format is utf8mb4, will be 4 byte
                - in `key_len` ( when we use EXPLAIN ), if it is varchar, will have extra 2 byte
                    - usually 255>=char is 1 byte, because execution planned in server layer, wont check with innoDB about real length, but just use 2 byte
        - index condition pushdown 
            - for `select * from table when a>1, b=3`
            - `Extra` will show `Using index condition`
            - will filter off record that doesnt meet directly within index tranversal, lesser lookup.
        - index selectivity
            - the order of composite index also important
            - in composite index, front index use more
                - so should use a more distinguish ( selectivity) column
                - > selectivity of a column = distinct count/ total count
                - selectivity too low, cant do more filter through that index!
                - Optimizer in MySQL: if found that certain value >30%, will skip index, and use full search
        - `select * from order where status = 1 order by create_time asc`
            - if just execute with index=status, will have `Extra`=`Using filesort` as it need to sort create_time
            - to avoid this, can composite index status+create_time. which is sorted after filtered.

6. when do we index/ not index
    - issues with index:
        - physical storage
        - create & maitain index need time
        - delete/add/update row efficiency impacted. B+ tree will need to update everytime
    - when to index/
        - unique column, like id
        - column that always used in `WHERE`
        - column that always used with `ORDER BY`, `GROUP BY`
    - when not to index
        - column that not used for `WHERE`, `ORDER BY`, `GROUP BY`
        - column that low selectivity (cant filter away much, and optimizer may skip it)
        - data too less
        - column that always update/add/delete

7. Index optimization
    1. prefix index
        - reduce index size, allow more index value in 1 index page, which improve query speed
        - but cant use for:
            - order by
            - cant use prefix index as covered index
    2. covered index
        - all column can be get by B+ tree leaf node.
        - data that get with secondary index, no need to go through primary index and lookup to get
        - example:
            - query: item_id, item_name, item_price
            - create composite index of these 3
            - and only query for these 3, avoid lookup from primary table
    3. self increment primary index
        - innon DB default use primary index as the preimary key
        - when new data come in, add to the back of B+ tree
        - if not self increment, new data may need to insert in middle of B+ tree
        - some data may need to shifted and create a new page, call page splitting
            - causing memory fragmentation
            - index less compact, lower efficiency
        - primary index should be short as well, so in secondary index, leaf node smaller, and less space
    
    4. use `NOT NULL` for index
        1. optimizer harded to optimize if contain `NULL`, example like: `count` will skip `NULL`
        2. `NULL` use space ( by the extra `NULL` list in row data infor)
    
    5. prevent index fail
        - index fail when: ( will show `type` as `ALL`)
            1. using left/ full wild card: `like %xx` or `like %xx%` ( not for right wild card!!)
            2. query with calculation, function, type conversion. `select * from a where b+1=10`
            3. composite index that not following leftmost prefix rule
            4. for `WHERE`, if condition of `OR` 1 is index, another is not
    
    6. for `EXPLAIN`
        1. possible_keys
            - possible used index
        2. key
            - actual used index
        3. key_len
            - used index length (in byte)
        4. rows
            - scanned rows
        5. type
            - `ALL`: full scan, bad.
            - `INDEX`: full scan on index table, better than `ALL` only on this no need to sort, but still bad.
            - `RANGE`: when involve `<` or `>` or `in` or `between`, start to be good with index
            - `REF`: when using non-unique index/ or prefix, cant stop when get the first match, need to do a small range scan. better
            - `EQ_REF`: when use primary or unqiue index. in multiple table join, best
            - `CONST`: when use primary or unqiue index, best than best, ( no need to join!)
        6. extra
            - `Using filesort`: when used `GROUP BY` or `ORDER BY` and cant used index, no choice to use sort. low efficient, need to avoid
            - `Using temporary`: a temporay table used to store intermdeiate result. usually when `GROUP BY` or `ORDER BY`, low efficeint, need avoid
            - `Using index`: covered index, nice!

        