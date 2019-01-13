-- 增删改 insert delete update
-- INSERT
INSERT INTO values (name,year) VALUES ("12",1908); -- 往 values 表插入 name和year的数据，数据的实际值是 "12" 1908
-- 多批次增加Insert
INSERT INTO values (name,year) VALUES
("1",12),
("2",3),
("3",23); --  primary key 也就是主键不用写出啦了，它自己会增加，另外有默认值的也可以不写，但是没有默认值并且不能为null的一定要写出来。
-- update
UPDATE values SET name='1',year=1 WHERE id >10;
--- update的计算
UPDATE values SET name=name +1 WHERE id>10 AND id <5;
-- ⚠️
UPDATE values SET name="1";-- 像这种没有WHERE的语句，非常的危险，容易更新全部的信息。
-- Delete
DELETE FROM values WHERE id=12;
-- ⚠️
DELETE FROM values --这样就删掉了全部，很危险
-- INNER INTO 和SELECT结合，也就是将values的值用SELECT的值来代替了。
INNER INTO coastroad(name,id) SELECT r.name,r,id FROM room r WHERE id>10;

