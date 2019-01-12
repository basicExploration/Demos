-- 从room表取出所有的数据
SELECT * FROM coastroad.room
-- 从room表取出所有的数据前提是id必须 >10
SELECT * FROM coastroad.room WHERE id >10;
-- 不带FROM的语句是为了测试，测试数据库连接是否正常
SELECT 1;
-- 条件查询 限制100个。
SELECT * FROM coastroad.room WHERE id>10 AND name="ap" LIMIT 100;
SELECT * FROM coastroad.room WHERE id>10 OR name="ap" LIMIT 100;
SELECT * FROM coastroad.room WHERE id>10 NOT name="ap" LIMIT 100;-- name 不等于"up" sql中没有赋值，所以= 就是等于的意思了，没有==
SELECT * FROM coastroad.room WHERE (id>10 OR name="ap") AND tt > 12 LIMIT 100;-- 使用小括号划分。如果不加括号，条件运算按照NOT、AND、OR的优先级进行，即NOT优先级最高，其次是AND，最后是OR。加上括号可以改变优先级。用小括号是一个非常好的喜欢，不论是sql或者是go或者是js。




