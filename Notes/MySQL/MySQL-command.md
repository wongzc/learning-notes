1. `show databases`
2. `create database database_name`
3. `use database_name`
4. ```CREATE TABLE `t_user` ( `id` int(11) NOT NULL, `name` VARCHAR(20) DEFAULT NULL, `phone` VARCHAR(20) DEFAULT NULL, `age` int(11) DEFAULT NULL, PRIMARY KEY (`id`) USING BTREE ) ENGINE = InnoDB DEFAULT CHARACTER SET = ascii ROW_FORMAT = COMPACT;```
5. ```CREATE TABLE test ( `name` VARCHAR(65535) NULL ) ENGINE = InnoDB DEFAULT CHARACTER SET = ascii ROW_FORMAT = COMPACT;```
6. `drop table table_name;`
7. create table primary key ( when crate table)
    `CREATE TABLE table_name ( .... PRIMARY KEY (index_column_1) USING BTREE );`
8. create unique index ( whne create table)
    `CREATE TABLE table_name ( .... UNIQUE KEY(index_column_1,index_column_2,...) );`
9. change to unique index (after crate table)
    `CREATE UNIQUE INDEX index_name ON table_name(index_column_1,index_column_2,...);`
10. create general index ( when create table)
    `CREATE TABLE table_name ( .... INDEX(index_column_1,index_column_2,...) );`
11. create general index ( after create table)
    `CREATE INDEX index_name ON table_name(index_column_1,index_column_2,...);`
12. prefix index ( when create table)
    `CREATE TABLE table_name( column_list, INDEX(column_name(length)) );`
13. prefix index ( after create table)
    `CREATE INDEX index_name ON table_name(column_name(length));`
14. show table ( need to use database first)
    `show tables`