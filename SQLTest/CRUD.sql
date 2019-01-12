-- 增删改 insert delete update
-- INSERT
INSERT INTO values (name,year) VALUES ("12",1908); -- 往 values 表插入 name和year的数据，数据的实际值是 "12" 1908
-- 多批次增加Insert
INSERT INTO values (name,year) VALUES
("1",12),
("2",3),
("3",23); --  primary key 也就是主键不用写出啦了，它自己会增加，另外有默认值的也可以不写，但是没有默认值并且不能为null的一定要写出来。
