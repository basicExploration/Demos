实例:

SQL模糊查询，使用like比较关键字，加上SQL里的通配符，请参考以下： 
1、LIKE'Mc%' 将搜索以字母 Mc 开头的所有字符串（如 McBadden）。 
2、LIKE'%inger' 将搜索以字母 inger 结尾的所有字符串（如 Ringer、Stringer）。 
3、LIKE'%en%' 将搜索在任何位置包含字母 en 的所有字符串（如 Bennet、Green、McBadden）。 
4、LIKE'_heryl' 将搜索以字母 heryl 结尾的所有六个字母的名称（如 Cheryl、Sheryl）。 
5、LIKE'[CK]ars[eo]n' 将搜索下列字符串：Carsen、Karsen、Carson 和 Karson（如 Carson）。 
6、LIKE'[M-Z]inger' 将搜索以字符串 inger 结尾、以从 M 到 Z 的任何单个字母开头的所有名称（如 Ringer）。 
7、LIKE'M[^c]%' 将搜索以字母 M 开头，并且第二个字母不是 c 的所有名称（如MacFeather）。 

理论:

SELECT 字段 FROM 表 WHERE 某字段 Like 条件


其中关于条件，SQL提供了四种匹配模式：

1，% ：表示任意0个或多个字符。可匹配任意类型和长度的字符，有些情况下若是中文，请使用两个百分号（%%）表示。

比如 SELECT * FROM [user] WHERE u_name LIKE '%三%'

将会把u_name为“张三”，“张猫三”、“三脚猫”，“唐三藏”等等有“三”的记录全找出来。

另外，如果需要找出u_name中既有“三”又有“猫”的记录，请使用and条件
SELECT * FROM [user] WHERE u_name LIKE '%三%' AND u_name LIKE '%猫%'

若使用 SELECT * FROM [user] WHERE u_name LIKE '%三%猫%'
虽然能搜索出“三脚猫”，但不能搜索出符合条件的“张猫三”。

2，_ ： 表示任意单个字符。匹配单个任意字符，它常用来限制表达式的字符长度语句：

比如 SELECT * FROM [user] WHERE u_name LIKE '_三_'
只找出“唐三藏”这样u_name为三个字且中间一个字是“三”的；

再比如 SELECT * FROM [user] WHERE u_name LIKE '三__';
只找出“三脚猫”这样name为三个字且第一个字是“三”的；


3，[ ] ：表示括号内所列字符中的一个（类似正则表达式）。指定一个字符、字符串或范围，要求所匹配对象为它们中的任一个。

比如 SELECT * FROM [user] WHERE u_name LIKE '[张李王]三'
将找出“张三”、“李三”、“王三”（而不是“张李王三”）；

如 [ ] 内有一系列字符（01234、abcde之类的）则可略写为“0-4”、“a-e”
SELECT * FROM [user] WHERE u_name LIKE '老[1-9]'
将找出“老1”、“老2”、……、“老9”；

4，[^ ] ：表示不在括号所列之内的单个字符。其取值和 [] 相同，但它要求所匹配对象为指定字符以外的任一个字符。

比如 SELECT * FROM [user] WHERE u_name LIKE '[^张李王]三'
将找出不姓“张”、“李”、“王”的“赵三”、“孙三”等；

SELECT * FROM [user] WHERE u_name LIKE '老[^1-4]';
将排除“老1”到“老4”，寻找“老5”、“老6”、……

5，查询内容包含通配符时 

由于通配符的缘故，导致我们查询特殊字符“%”、“_”、“[”的语句无法正常实现，而把特殊字符用“[ ]”括起便可正常查询。据此我们写出以下函数：



function sqlencode(str)
str=replace(str,"[","[[]") '此句一定要在最前
str=replace(str,"_","[_]")
str=replace(str,"%","[%]")
sqlencode=str
end function
在查询前将待查字符串先经该函数处理即可，并且在网页上连接数据库用到这类的查询语句时侯要注意：

如Select * FROM user Where name LIKE '老[^1-4]';上面 《'》老[^1-4]《'》

 
