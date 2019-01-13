-- 关于mysql本身

-- 创建一个数据库 CREAT DATABASE  [databasename]
CREATE DATABASE coastroad;
-- 切换到这个数据库
USE coastroad;
-- 删除一个数据库
DROP DATABASE coastroad;
-- 显示 ：显示出所有的数据库 显示出所有的表
SHOW DATABASES ; SHOW TABLES;
--- 断开client和server的连接。
EXIT;
--- 事务
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ ;
BEGIN ;-- 开始内容
INNER INTO values(id ,name n)values (99,92);
COMMIT ;-- 停止

