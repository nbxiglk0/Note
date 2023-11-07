# PHP-AuditBasic
- [PHP-AuditBasic](#php-auditbasic)
  - [PHP配置相关](#php配置相关)
    - [register\_globals(全局变量开关)](#register_globals全局变量开关)
    - [allow\_url\_include(包含远程文件)](#allow_url_include包含远程文件)
    - [magic\_quotes\_gpc(runtime, sybase)(魔术引号过滤)](#magic_quotes_gpcruntime-sybase魔术引号过滤)
    - [safe\_mode(安全模式)](#safe_mode安全模式)
    - [open\_basedir(可访问目录)](#open_basedir可访问目录)
    - [disable\_function(禁用函数)](#disable_function禁用函数)
    - [display\_errors,error\_reporting](#display_errorserror_reporting)
  - [SQL注入](#sql注入)
    - [宽字节](#宽字节)
    - [urldecode()/rawurldecode()](#urldecoderawurldecode)
    - [addslashes()](#addslashes)
  - [文件包含](#文件包含)
    - [include(),include\_once()](#includeinclude_once)
    - [require(),require\_once()](#requirerequire_once)
    - [文件包含截断](#文件包含截断)
  - [任意文件读取](#任意文件读取)
  - [任意文件写入](#任意文件写入)
  - [任意文件上传](#任意文件上传)
  - [任意文件删除](#任意文件删除)
  - [XXE](#xxe)
  - [SSRF](#ssrf)
  - [代码执行](#代码执行)
    - [preg\_replace()](#preg_replace)
  - [命令执行](#命令执行)
    - [防御](#防御)
  - [变量覆盖](#变量覆盖)
  - [反序列化](#反序列化)
    - [常见魔术方法](#常见魔术方法)
  - [逻辑问题](#逻辑问题)
  - [Tricks](#tricks)
    - [解析标签](#解析标签)
    - [php:// 输入输出流](#php-输入输出流)
    - [file\_put\_content()](#file_put_content)
    - [upload\_progress](#upload_progress)
    - [str\_replace](#str_replace)
## PHP配置相关
和安全相关的php配置.
### register_globals(全局变量开关)
1. 将用户GET,POST提交的数据自动注册为全局变量并初始化为对应的值.
2. 5.4被移除.
### allow_url_include(包含远程文件)
1. 开启后可以包含远程文件.
2. 5.2之后默认为off关闭.
### magic_quotes_gpc(runtime, sybase)(魔术引号过滤)
1. 开启后GET,POST,COOKIE变量中的单引号,双引号,反斜杠,空字符(null)的前面会加上反斜杠进行转义.
2. 不会过滤$_SERVER变量.
3. runtime和gpc区别在于只对数据库和文件中的数据进行过滤.
4. sybase的区别在于只转义空字符和把单引号变成双引号.
5. 上述5.4之后该配置被取消.
### safe_mode(安全模式)
1. 开启后文件操作属性会有属主权限限制.
2. popen(),system(),exec()等函数执行命令会失败.
### open_basedir(可访问目录)
1. 限制PHP能访问的目录.
### disable_function(禁用函数)
1. 禁止执行的函数名单.
### display_errors,error_reporting
1. 是否显示php脚本内部错误.
2. error_reporting用于选择显示的级别.
## SQL注入
### 宽字节
mysql连接时.设置字符集为`set character_set_client=gbk`时出现或者使用相关编码转换函数时也可能出现.
### urldecode()/rawurldecode()
双重解码导致绕过过滤.
### addslashes()
1. 同GPC类似进行过滤,参数类型必须是string类型.
2. 宽字节,双重编码,数字型,使用了stripaddslashes()可绕.
## 文件包含
### include(),include_once()
遇到错误时会继续执行脚本.include_once()只会包含一次.
### require(),require_once()
遇到错误时会报错退出.require_once()只会包含一次.
### 文件包含截断
1. %00会被gpc和addslashes()过滤,且在5.3版本之后被完全修复.
2. 5.3之前还可以尝试超长文件名进行截断.windows下240个点,Linux时2038个`/.`可以截断.
3. 远程包含时,使用?来进行伪截断.
## 任意文件读取
常见文件读取函数:
```
file_get_contents(),highlight_file(),fopen(),readfile(),fgetss(),fgets(),parse_ini_file(),show_source(),file()
```
文件包含函数如include()可以使用php输入输出流`php://filter/`来读取.
## 任意文件写入
* file_put_conten()
* fwrite()
## 任意文件上传
文件上传函数:move_uploaded_file()
## 任意文件删除
文件删除函数:unlink()
## XXE
* simplexxml_load_string
* Libxml2.9.0 以后 ，默认不解析外部实体.
* 支持的协议: file,http,ftp,php,compress,data,glob,phar
## SSRF
* curl
* file_get_content()
* fsockopen()
* 支持的协议:gopher协议,dict协议,file协议,http/s...
## 代码执行
* eval()
* assert()
* call_user_func()
* ...
### preg_replace()
该函数代码执行需要第一个参数存在`/e`参数,会导致第二个参数的值会被当成php代码执行,而第二个参数的值来自于第一个参数匹配的结果.
## 命令执行
* system()
* exec()
* passthru()
* 反引号``(其实调的shell_exec()函数)....
### 防御
过滤函数:escapeshellcmd(),escapeshellarg().
## 变量覆盖
* extact()
* parse_str()
* `$$`覆盖
## 反序列化
### 常见魔术方法
* __construct: 对象创建时调用.
* __destruct: 对象销毁时调用.
* __toString: 对象被当作字符串时调用.
* __sleep(): 对象反序列化之前运行.     
* __wakeup():对象被反序列化之前调用.
## 逻辑问题
* in_array():判断是否存在于数组中,但会自动进行类型转换造成绕过(如2aaa会被转化为2).
* is_numeric():判断是否为数字,但为hex时会直接返回true,而hex在sql中可以使用(insert,update语句中).
* `==`和`===`:`==`在比较前会做类型转换可能导致绕过(如2aaa会被转化为2),而`===`不会.
* exit(return):未手动退出程序导致代码继续执行.
## Tricks
### 解析标签
* 脚本标签:`<script language="php"> <script>`
* 短标签:`<? ... ?>`
* asp标签:`<% ... %>`
### php:// 输入输出流
* php://input: 获取post原始的输入流(不能读表单)
* php://filter: 文件操作过滤器,可读写文件.
### file_put_content()
file_put_contents的第二个参数,可以传入一个数组.PHP会将这个数组拼接成字符串,写入文件中.可以用来绕过关键字检测(将关键字拆分).
### upload_progress
### str_replace