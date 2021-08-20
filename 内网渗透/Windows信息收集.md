<a name="iurdI"></a>
###### Powershell命令执行记录
```powershell
type (Get-PSReadlineOption).HistorySavePath
```
<a name="nvjPH"></a>
###### 获取安装软件列表
```powershell
wmic product get name,version
poweshell "Get-WmiObject -class Win32-Porduct | Select-Object -Property name,version
```
<a name="7dFUs"></a>
###### 获取本地服务信息
`wmic service list brief`
<a name="XSoTG"></a>
###### 获取启动程序信息
`wmic startup get command,caption`
<a name="idF1J"></a>
###### 查看计划任务
`schtasks /query /fo LIST /v`
<a name="9H4i8"></a>
###### 开机时间
`net statistics workstation`
<a name="reexi"></a>
###### 重启时间
<a name="JRwKT"></a>
###### <br />`dir /a c:\pagefile.sys<br />dir /a \machine\c$\pagefile.sys`
<a name="PHk2l"></a>
###### 查看当前在线用户
`query user | qwinsta`
<a name="I6Dye"></a>
###### 列出会话
`net session`
<a name="KNR5C"></a>
###### 查看补丁
  1. `systeminfo`

2. `wmic qfe get`
<a name="IVFoC"></a>
###### 查看共享列表
`wmic share get name,path,status`
<a name="bDeKc"></a>
###### 防火墙设置

1. 关闭防火墙(Administrator Require)

version <= 2003<br />`netsh firewall set openmode disable`<br />version > 2003<br />`netsh firewall set allprofiles state off`

2. 查看防火墙配置

`netsh firewall show config`

3. 允许指定程序全部连接

version <= 2003<br />`netsh firewall add allowedprogram c:\shell.exe "allowed" enable`<br />version > 2003<br />`netsh advfirewall firewall add rule name="allowed" dir=(in)(out) action=allow program="c:\shell.exe"`

4. 允许端口放行(3389)

`netsh advfirewall firewall add rule name="Remote Desktop" protocol=TCP dir=in localport=3389 action=allow`

5. 修改防火墙配置文件存储路径

`netsh advfirewall set currentprofile logging filename "C:\tmp\tmp.log"`
<a name="0u80w"></a>
###### 查看代理配置
`reg query "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Internet Settings"`
<a name="d3DqO"></a>
###### 查看远程桌面端口
`REG QUERY "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp" /V PortNumber`
<a name="aXFum"></a>
###### 开启远程桌面
`Version == 2003`<br />```<br />`Version 2008 & 2012`
```powershell
wmic /namespace:\\root\cimv2\terminalservices path win32_terminalservicesetting where (_CLASS !="") CALL setallowtsconnections 1
wmic /namespace:\\root\cimv2\terminalservices path win32_tsgeneralsetting where (TerminalName='RDP-Tcp') call setuserauthenticationrequired 1
reg add "HKLM\SYSTEM\CURRENT\CONTROLSET\CONTROL\TERMINAL SERVER" /v fSingSessionPerUser /t REG_DWORD /d 0 /f
```
<a name="8m0jV"></a>
###### 探测内网存活主机
`For /L %I in (1,1,254) DO @ping -w 1 -n 1 192.168.1.%I | findstr "TTL="`<br />查询域成员计算机列表<br />`net group "domain computers" /domain`<br />获取域信任关系<br />`nltest /domain_trusts`<br />域内用户详细信息<br />`wmic useraccount get /all`<br />获取域内存在用户<br />`dsquery user`<br />寻找域控<br />Domain name:test.lab<br />`nltest /dclist:test.lab`<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1626970526447-d41bb520-521e-4a39-9107-8b4142207bf4.png#clientId=ud78b5e20-a7a8-4&from=paste&height=60&id=u2640c462&margin=%5Bobject%20Object%5D&name=image.png&originHeight=119&originWidth=800&originalType=binary&ratio=1&size=10078&status=done&style=none&taskId=ud5060bbb-6ae2-475d-baf1-97392459aeb&width=400)<br />`nslookup -q=srv _ldap._tcp.dc._msdcs.test.lab`<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1626970550347-97abca2e-9b23-45f1-83ae-9248129262dc.png#clientId=ud78b5e20-a7a8-4&from=paste&height=147&id=uf091e427&margin=%5Bobject%20Object%5D&name=image.png&originHeight=294&originWidth=785&originalType=binary&ratio=1&size=18792&status=done&style=none&taskId=u5f694423-0c07-48e6-a598-1039bcab1c5&width=392.5)
<a name="yAW1d"></a>
###### 寻找域管理员登录的会话

1. 查询域管理员列表 `net group "Domain Admins" /domain`
1. 向域控制器查询域的所有活动会话
1. 将得到活动会话列表和管理员列表交叉对比得到域管理员正在登录的会话

导出所有SPN信息<br />`setspn -T domain -q */*`
