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
-- 投影+条件
SELECT name,year FROM coastroad.room WHERE id>10 LIMIT 100;
-- 排序 ORDER BY
SELECT name,year FROM coastroad.room WHERE id>10 ORDER BY id LIMIT 100;
-- 排序 ORDER BY + 倒序：DESC
SELECT name,year FROM coastroad.room WHERE id>10 ORDER BY id DESC LIMIT 100;
-- 多条件排序 ORDER BY + 倒序：DESC
SELECT name,year FROM coastroad.room WHERE id>10 ORDER BY id, name  DESC LIMIT 100;-- 这样会按照优先级 id 然后再name 来排序 order要在where后面
-- 分页1
SELECT name,year FROM coastroad.room WHERE id>10 ORDER BY id, name  DESC LIMIT 100 OFFSET 0;
-- 分页2
SELECT name,year FROM coastroad.room WHERE id>10 ORDER BY id, name  DESC LIMIT 100 OFFSET 100;-- 意思是按照100 一个集合来分类，offset是偏移了100 意思就是从101开始计算
-- 分页的简写
SELECT name,year FROM coastroad.room WHERE id>10 ORDER BY id, name  DESC LIMIT 100 100 OFFSET 0; -- 可以省略掉OFFSET.
-- 给表设置别名
SELECT * FROM room r
-- 聚合查询 number
SELECT COUNT(*) num FROM room; -- 必须带;不然 sql不知道什么时候才结束。
-- 聚合查询 + where
SELECT COUNT(*) num WHERE gender >1 ORDER BY name DESC LIMIT 100 100;
-- 除了COUNT()函数外，SQL还提供了如下聚合函数：
-- 函数	说明
-- SUM	计算某一列的合计值，该列必须为数值类型
-- AVG	计算某一列的平均值，该列必须为数值类型
-- MAX	计算某一列的最大值
-- MIN	计算某一列的最小值
-- 注意，MAX()和MIN()函数并不限于数值类型。如果是字符类型，MAX()和MIN()会返回排序最后和排序最前的字符。
-- 要特别注意：如果聚合查询的WHERE条件没有匹配到任何行，COUNT()会返回0，而MAX()、MIN()、MAX()和MIN()会返回NULL
--- 分组聚合
SELECT COUNT(*) num FROM number GROUP BY student_id ;
-- 多表查询
SELECT c.name, d.name FROM room c,dd d










