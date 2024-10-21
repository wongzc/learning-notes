1. `show databases`
2. `create datbase database_name`
3. `use database_name`
4. ```CREATE TABLE `t_user` ( `id` int(11) NOT NULL, `name` VARCHAR(20) DEFAULT NULL, `phone` VARCHAR(20) DEFAULT NULL, `age` int(11) DEFAULT NULL, PRIMARY KEY (`id`) USING BTREE ) ENGINE = InnoDB DEFAULT CHARACTER SET = ascii ROW_FORMAT = COMPACT;```
5. ```CREATE TABLE test ( `name` VARCHAR(65535) NULL ) ENGINE = InnoDB DEFAULT CHARACTER SET = ascii ROW_FORMAT = COMPACT;```
6. `drop table table_name;`