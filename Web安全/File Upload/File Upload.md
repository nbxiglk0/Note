- [File Upload](#file-upload)
  - [ByPass](#bypass)
    - [前端JS过滤](#前端js过滤)
    - [MIME类型](#mime类型)
    - [解析漏洞](#解析漏洞)
    - [0x00截断](#0x00截断)
    - [文件内容头校验（GIF89a)](#文件内容头校验gif89a)
    - [二次包含](#二次包含)
  - [防御](#防御)

# File Upload
## ByPass
### 前端JS过滤
Burp抓包修改
### MIME类型
针对根据MIME类型来确定文件类型的.
1. 修改Content-type类型绕过白名单或者黑名单. 
2. 直接删除Content-type.
### 解析漏洞
1. Apache
2. IIS
3. Nginx
### 0x00截断
php version < 5.2  
java jdk < 1.6
### 文件内容头校验（GIF89a)
### 二次包含
1. 只检测(jsp,php,aspx等后缀)的内容.
先上传包含恶意代码的txt等不会被检测的文件后缀,再上传一个正常的(jsp,php,aspx文件)来包含引用txt文件.
## 防御
1. 上传目录不可执行
2. 白名单
3. 文件名随机数