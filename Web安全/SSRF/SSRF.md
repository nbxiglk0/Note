- [Server-side request forgery (SSRF)](#server-side-request-forgery-ssrf)
  - [Impact](#impact)
  - [产生SSRF的函数](#产生ssrf的函数)
    - [php](#php)
  - [SSRF中URL伪协议](#ssrf中url伪协议)
  - [云上SSRF](#云上ssrf)
  - [Bypass](#bypass)
    - [进制转换](#进制转换)
    - [域名欺骗](#域名欺骗)
    - [编码](#编码)
    - [`@`绕过](#绕过)
    - [`#`绕过](#绕过-1)
    - [白名单子域名](#白名单子域名)
    - [302重定向](#302重定向)
    - [短链接](#短链接)
  - [Bind SSRF](#bind-ssrf)
  - [修复方案](#修复方案)
  - [参考](#参考)
# Server-side request forgery (SSRF)
服务端请求伪造,简单来说就是可以让目标服务器作为代理去请求指定的任意地址。
## Impact

* 访问本地或者内网服务器监听的端口  
一些监听在本地的端口正常情况下无法被其他人访问，但通过SSRF可以使服务器自身做代理来访问到本地的端口。  
```
redis 6379
fpm 9000
...
```
* 绕过认证  
有一些接口是需要认证才能进行访问的，但是认证只限于其它IP进行访问时，当通过内网来请求这些接口时可能不需要验证。 
* 读取文件
* 直接用已知EXP攻击内网Server(Str2,redis,elastic...).
## 产生SSRF的函数
### php
* file_get_contents
* sockopen()
* curl_exec()
* readfile()
* fopen()
* ...
## SSRF中URL伪协议
当没有限制请求的协议时,可以尝试使用其他的协议扩展攻击面.
file:/// 从文件系统中获取文件内容，如，file:///etc/passwd  
dict:// 字典服务器协议，访问字典资源，如，dict:///ip:6739/info：  
sftp:// SSH文件传输协议或安全文件传输协议  
ldap:// 轻量级目录访问协议   
tftp:// 简单文件传输协议  
gopher:// 分布式文档传递服务，可使用gopherus生成payload  
## 云上SSRF
Amazon Elastic Compute Cloud (Amazon EC2)中的每个实例，都可以通过执行`curl -s http://169.254.169.254/user-data/`，对IP 169.254.169.254发出HTTP请求，来获取该云实例自身的元数据。
## Bypass
### 进制转换
127.0.0.1 -> 2130706433 -> 017700000001 -> 127.1
### 域名欺骗
将恶意域名的ip解析为127.0.0.1。Tools:spoofed.burpcollaborator.net
### 编码
将关键字符url编码或大小写混淆。
1. 双字编码
2. URL编码
3. 16进制编码
4. 8进制编码
### `@`绕过
一些白名单中只是匹配了url的起始或者是否包含某些白名单关键字,可以使用url的一些来绕过.
* 通过RFC标准,url`@`前面的部分将会被视为用户密码,而`@`后面的部分才会被视为目标服务器. 
`https://whitelist-host@evil-host`
### `#`绕过
* 通过`#`锚点在恶意host中加入白名单host.
`https://evil-host#whitelist-host`
### 白名单子域名
在自己的域名下注册一个白名单host的恶意子域名.
* `https://whitelist-host.evil-host`
### 302重定向
因为很多防御措施都是在请求前对路径进行过滤和检测,如果该SSRF漏洞后端支持重定向的话则可以利用重定向来绕过很多黑名单,如果应用自身就存在Openredirection漏洞的话也可以绕过大部分白名单,先请求一个合法的远程服务器,通过控制远程服务器返回302状态码在location Header来再次请求任意地址.  

php快速搭建302跳转服务器,默认执行-t指定目录下的index.php.  
`php -s localhost:80 -t ./`  
```php
<?php
Header('Location: http://localhost:8080/console')
?>
```
### 短链接
1. 使用4m短网址或站长工具即可（短链接本质是302跳转）
2. is.gd可以自定义后缀的短网址
百度短地址等等
## Bind SSRF

## 修复方案
1. 如果可以,使用白名单
2. 对响应进行验证,过滤返回信息，验证远程服务器对请求的响应是比较容易的方法。如果web应用是去获取某一种类型的文件。那么在把返回结果展示给用户之前先验证返回的信息是否符合标准。
3. 统一错误信息，避免用户可以根据错误信息来判断远端服务器的端口状态。
4. 限制请求的端口为http常用的端口，比如，80,443,8080,8090。
5. 限制访问内网,黑名单内网ip。避免应用被用来获取获取内网数据，攻击内网。
6. 协议限制,禁用不需要的协议。仅仅允许http和https请求。可以防止类似于file:///,gopher://,ftp:// 等引起的问题。
## 参考
https://portswigger.net/web-security/ssrf   
https://blog.csdn.net/qq_43378996/article/details/124050308  
https://www.blackhat.com/docs/us-17/thursday/us-17-Tsai-A-New-Era-Of-SSRF-Exploiting-URL-Parser-In-Trending-Programming-Languages.pdf  
https://xz.aliyun.com/t/7198