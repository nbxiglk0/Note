- [文件上传漏洞 File Upload](#文件上传漏洞-file-upload)
  - [ByPass](#bypass)
    - [前端JS过滤](#前端js过滤)
    - [MIME类型](#mime类型)
    - [解析漏洞](#解析漏洞)
    - [0x00截断](#0x00截断)
    - [文件内容头校验(GIF89a)](#文件内容头校验gif89a)
    - [二次包含](#二次包含)
  - [修复方案](#修复方案)
    - [限制上传目录权限](#限制上传目录权限)
    - [白名单校验](#白名单校验)
    - [文件名随机数](#文件名随机数)
    - [存储分离](#存储分离)

# 文件上传漏洞 File Upload
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
### 文件内容头校验(GIF89a)
### 二次包含
1. 只检测(jsp,php,aspx等后缀)的内容.
先上传包含恶意代码的txt等不会被检测的文件后缀,再上传一个正常的(jsp,php,aspx文件)来包含引用txt文件.
## 修复方案
### 限制上传目录权限
设置目录权限限制，禁止上传目录的执行权限。
### 白名单校验
对后缀做白名单校验
```java
// 获取文件后缀名
        String Suffix = fileName.substring(fileName.lastIndexOf("."));

        String[] SuffixSafe = {".jpg", ".png", ".jpeg", ".gif", ".bmp", ".ico"};
        boolean flag = false;

        for (String s : SuffixSafe) {
            if (Suffix.toLowerCase().equals(s)) {
                flag = true;
                break;
            }
        }
```
### 文件名随机数
对上传的文件回显相对路径或者不显示路径,文件保存时，将文件名替换为随机字符串。
### 存储分离
在OSS静态服务器上存储文件,将上传的文件和服务器分开.