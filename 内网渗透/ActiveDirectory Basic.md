---



<a name="Nww6z"></a>
# Basic

1. Functional Modes,域控需求的最小操作系统,不同的Mode(操作系统)有不同的功能特点,如windows2012R2Mode,其拥有`Protected Users group`组.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1626528058941-1766efb8-185b-4046-8e48-ff8a8d567ee1.png#height=250&id=BzCW5&margin=%5Bobject%20Object%5D&name=image.png&originHeight=499&originWidth=1404&originalType=binary&ratio=1&size=129519&status=done&style=none&width=702)
1. 计算机账户在计算机加入域时自动生成，账户名为加入域时的计算机名后面加一个$符号，密码为自动生成，且30天后会自动更新.
1. 通过NTLM或者Kerberos进行身份验证的用户不会缓存凭据在目标计算机中(除非启用了Kerberos委派).
1. 从2008R2开始,默认不会缓存明文密码在内存中,但可以                                                                                                                                             修改注册表HKLM\SYSTEM\CurrentControlSet\Control\SecurityProviders\WDigest\UseLogonCredential的值为1来重新开启该功能,或者通过Pacthing Digest SSP.
<a name="YMs7f"></a>
## Groups
根据作用域的不同可以分为三种类型的组:

- 通用组: 同一个林下任何域的用户,同一个林下其他域的全局组,同一个林下任何域的通用组.
- 全局组: 同一个域下的用户,同一个域下其他的全局组.
- 本地域组: 
<a name="RGNjj"></a>
##### Domain Admins
域管理员组,其中的用户对域内的成员拥有管理员权限.默认添加到域中所有计算机都Administrators组中.
<a name="DTeWk"></a>
#### Enterprise Admins
对所有林中的所有用户拥有管理员权限,只在林的根域上存在,但默认加入在林中所有域中的Administrator组.
<a name="Q85zU"></a>
#### DnsAdmins
该组允许其成员在域控中使用任意的DLL来执行代码.[参考](https://www.semperis.com/blog/dnsadmins-revisited/)
<a name="t68i1"></a>
#### Protected User
该组用户强制执行一些安全措施,防止NTLM中继或者委派攻击.

- 不允许使用NTLM认证(在Kereros协议中).
- 不允许在Kerberos预认证中使用DES或者RC4算法.
- 不允许非约束性委派或者约束性委派.
- 不允许在请求TGT之后四小时后续约TGT.
<a name="kSV0P"></a>
#### Remote Desktop Users
该组用户可以进行远程桌面连接.
<a name="mIim8"></a>
#### Group Policy Creator Owners
该组成员能够修改组策略
<a name="tofNo"></a>
## Port

- 53->DNS
- 88->Kerberos
- 135->RPC 端点映射
- 139->Netbios
- 389->LDAP
- 445->SMB
- 464->kpasswd
- 593->RPC HTTP 端点映射
- 636->带SSL的LDAP
- 3268->LDAP 全局目录
- 3269->带SSL的LDAP全局目录
- 5985->WinRM,通过CIM对象或者远程Powershell进行远程管理的服务
- 9389->ADWS,管理域数据库的Web服务 



<a name="YawfL"></a>
## Hashs
[NT_LAN_Manager](https://en.wikipedia.org/wiki/NT_LAN_Manager#NTLMv1)<br />导出用户Hash的格式为`<username>:<rid>:<LM>:<NT>:::`.<br />LM未使用的话值为	aad3b435b51404eeaad3b435b51404ee,为空字符串的加密值.
<a name="aDjm1"></a>
### LM hash
长度为14位,不足的用空,字节填充,超过14位的截断,从Vista/windows2008开始默认关闭,可通过GPO重启开启,存储在本地SMA数据库或者域控的Ntds数据库中,长度为16位.<br />生成流程:

1. 统一转换为大写.
1. 用空字节填充不足的或者截断超过14位的.
1. 将14位数据块分割为两块7位作为DES加密的密钥.
1. DES加密字符串""KGS!@#$%".
1. 将加密后的字符串拼接得到LMhash.
<a name="hqNEG"></a>
### NT hash
也叫NTLM-hash,存储在SAM数据库或者域控的Ntds数据库中,主要用于工作组环境下用户登录,使用MD4加密算法加密.Pass the hash即用的该Hash.
<a name="jD1zj"></a>
## Trusts

1. 同一个林下不同域的用户默认情况下可以访问对方域的资源,不同林的用户默认情况不能访问对方域的资源.
1. Windows NT下,信任关系是不可传递并且单向的,win2000-2008及以后则是可传递并且双向信任的.
1. 当一个客户端想要使用NTLM协议认证来访问另一个域的资源时,该资源服务器必须联系含有客户端账户的域控来验证该客户端的凭据.
1. 当两个域建立信任关系之后,在两个域中都会创建一个用户来存储信任秘钥,该账户的名称为另一个域的NetBios名称加一个$符号.
<a name="yjbCM"></a>
## SPNEGO
<a name="qFJhN"></a>
### GSS-API/SSPI
一种应用编程的接口,定义了由安全包实现的过程和类型,可以通过调用GSS-API过程来使用MIT Kerberos.<br />一些GSS-API过程如下:

- gss_acquire_cred:返回凭据句柄.
- gss_init_sec_context:初始化安全上下文.
- gss_accept_sec_context:接受安全上下文.
- gss_get_mic:计算MIC(消息完整性).
- gss_verify_mic:检查MIC.
- gss_wrap:将MIC附加到消息中并加密消息内容.
- gss_unwrap:验证MIC并解密消息内容.

Program that can use Kerberos or NTLM authentication
```
                     .---------------------------.
                     |   Kerberos Library        |
                     .---            .----       |
               .---> | GSS-API  ---> | Kerberos  |
               |     '---            '----       |
               |     |                           |
 .---------.   |     '---------------------------'
 |  user   |---|
 | program |   |     .---------------------------.
 '---------'   |     |       NTLM  Library       |
               |     .---            .----       |
               '---> | GSS-API  ---> | NTLM      |
                     '---            '----       |
                     |                           |
                     '---------------------------'
```
Program that uses Negotiate (SPNEGO)
```
                                             Kerberos
                                         .-------------------------.
                                         |      kerberos.dll       |
                                         |-------------------------|
                                         .---           .----      |
                   Negotiate       .---> | GSS-API ---> | Kerberos |
                 .-------------.   |     '---           '----      |
                 | secur32.dll |   |     |                         |
                 |-------------|   |     '-------------------------'
 .---------.     .---          |   |
 |  user   |---->| GSS-API ----|>--|
 | program |     '---          |   |         NTLM
 '---------'     |             |   |     .-------------------------.
                 '-------------'   |     |       msv1_0.dll        |
                                   |     |-------------------------|
                                   |     .---           .----      |
                                   '---> | GSS-API ---> | NTLM     |
                                         '---           '----      |
                                         |                         |
                                         '-------------------------'
```
<a name="JYGyu"></a>
### 协商机制

1. 客户端调用_GSS_Init_sec_context_并指示将使用 SPNEGO,得到一些带有安全选项的列表(mechTypes)和初始令牌(mechToken),包含在NegTokenInit消息中发送给服务器.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627400289758-0520cb74-c1cd-405a-834e-05353b292e34.png#clientId=u3891bd8e-3af0-4&from=paste&height=656&id=uae400b64&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1311&originWidth=3806&originalType=binary&ratio=1&size=286691&status=done&style=none&taskId=u5e9ba205-0c86-4e49-9bac-f095967200c&width=1903)
1. 服务器应用程序将初始令牌和安全机制列表传递给_GSS_Accept_sec_context_。然后返回以下结果之一并在_NegTokenResp_消息中发送（_NegTokenResp_ 与Wireshark 显示的_NegTokenTarg_相同）：
- 不接受任何安全机制。服务器拒绝协商。
- 如果所选的安全机制是客户端的首选，则使用收到的令牌。创建包含_accept-complete_状态的协商令牌 。
- 选择了首选机制之外的其他机制，则创建具有_accept-incompleted_或_请求request-mic_ 的协商令牌。

![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627400468688-2ec9038c-93bb-4c1e-b4d3-3f7d74d9a190.png#clientId=u3891bd8e-3af0-4&from=paste&height=656&id=u0c6d1b1e&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1311&originWidth=3751&originalType=binary&ratio=1&size=307947&status=done&style=none&taskId=u54f96b51-b668-4820-8495-5f401258660&width=1875.5)

3. 如果协商返回给客户端，则将其传递给 _GSS_Init_sec_context_并对其进行分析。协商一直持续到客户端和服务器都同意安全机制和选项。

_SPNEGO 协商_
```
                                     Client              Server
                                        |                 |
 GSS_Init_sec_context(SPNEGO=True) <--- |                 |
                                   ---> |   NegTokenInit  |
                            1) Kerberos | --------------> |  
                               (Token)  |    Security?    |  
                            2) NTLM     |    1) Kerberos  |
                                        |       (Token)   |
                                        |    2) NTLM      | Kerberos (Token)
                                        |                 | ---> GSS_Accept_sec_context()
                                        |   NegTokenResp  | <---
                                        | <-------------- | (Token)
                                        |     (Token)     | accept-complete
                                  Token | accept-complete |
            GSS_Init_sec_context() <--- |                 | 
                                        |                 |
                                        |                 |
```
<a name="C4NEp"></a>
## Services(SPN)
不是所有远程服务都在AD数据库中注册,但是使用kerberos验证的服务都要在域中注册.<br />AD中每个注册的服务包含以下信息:

- 运行该服务的用户
- 服务类别
- 运行该服务的机器
- 服务端口(可选)
- 服务路径(可选)

注册的服务名(SPN)格式如下:<br />`service_class/machine_name[:port][/path]`<br />machine_name可以是主机名或者FQDN名称,如下两种都兼容<br />`ldap/DC01`<br />`ldap/DC01.test.lab`<br />默认情况下,机器注册了一个HOST服务类,用于Windows默认自带的服务的别名.
<a name="Vm88d"></a>
## Principals
<a name="VmVBy"></a>
### Service Principal Name(SPN)
服务主体名称(SPN),在AD中,同一个服务可以运行在不同的主机上,而一个主机上可以运行不同的服务,而SPN则是用主机名和服务名的组合来区别不同的服务.<br />例如: `service_class/hostname_or_FQDN(:Port)`,而其中的service_class即服务类为一种集合,表示一类服务,如所有的Web服务都为"www"类,而如果使用了自定义端口的话,则可以将端口号添加在主机名后面.
<a name="cJQCN"></a>
#### Host service_class
Host SPN并不是真正的服务类,而是一组服务类,包含了许多预定义的服务类,可以通过以下命令列出`Get-ADObject -Identity "CN=Directory Service,CN=Windows NT,CN=Services,CN=Configuration,DC=HALO,DC=NET" -properties sPNMappings`.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628495269355-3a3287b0-4834-4000-b71b-ee0209f927fc.png#clientId=u4ebbf284-a4a6-4&from=paste&height=165&id=udeb1c2a9&margin=%5Bobject%20Object%5D&name=image.png&originHeight=329&originWidth=1168&originalType=binary&ratio=1&size=108557&status=done&style=none&taskId=u60662f62-9485-4b14-8f64-0ca636f01a0&width=584)<br />即请求预定义里面的class时,可以直接指定service_class为host即可.
<a name="mKfi3"></a>
## SID
为了标识不同的主体,每一个主体都有一个SID来进行标识.
<a name="mCFHu"></a>
### Domain SID
域SID,用来标识不同的域.
```powershell
PS C:\> $(Get-ADDomain).DomainSID.Value
S-1-5-21-1372086773-2238746523-2939299801
```
<a name="hhT2p"></a>
### Principals SID
用来标识不同的主体,该ID由域SID加相对标识符组成(RID).
```powershell
PS C:\> $(Get-ADUser Anakin).SID.Value
S-1-5-21-1372086773-2238746523-2939299801-1103
```
<a name="spYfP"></a>
### 其它常见SID

- S-1-5-11 -> 已经登录的用户组.
- S-1-5-10 -> 主体自身,用于在安全描述符中引用自己.
- S-1-5-21-domain-500 -> Adminstrator
```powershell
PS C:\> $(Get-ADUser Administrator).SID.Value
S-1-5-21-1372086773-2238746523-2939299801-500
```

- S-1-5-21-domain-512 -> Domain Admins
- S-1-5-21-domain-513 -> Domain users
- S-1-5-21-root  domain 519-> Enterprise Admins
<a name="PaTZX"></a>
## DistinguishedName(DN)
用于指示对象在数据库中的层级
```powershell
PS C:\> Get-ADComputer dc01 | select -ExpandProperty DistinguishedName
CN=DC01,OU=Domain Controllers,DC=contoso,DC=local
```
<a name="WRjbE"></a>
### 组成部分:

- Domain Component(DC):一般用来标识数据库的域部分,如test.lab则为`DC=test,DC=lab.`
- Organizational Unit(OU):用于标识多个对象的容器,类似于组,但目的不同,OU 的目的是组织数据库中的对象，而安全组用于组织域/林中的权限.但有时候会将OU直接映射到安全组.可以将GPO直接应用到OU.
- Common Name(CN):标识对象,有时候有多个CN是因为有时候有些对象同时还是容器
<a name="Kp5YN"></a>
## 全局目录
为了加快林中搜索对象的速度，某些域控还包含了其他域的对象子集，但只有只读权限，监听端口为3268( ldap).
<a name="u3gQK"></a>
## 数据库查询
<a name="abQrN"></a>
### LDAP
[LDAP Wiki](https://ldapwiki.com/)
```markdown
                      .-------------
                      |
                    .---
           .--TCP-->| 389 LDAP
           |        '---
           |          |
           |        .---
           |--SSL-->| 636 LDAPS
 .------.  |        '---
 | LDAP |--|          |
 '------'  |        .---
           |--TCP-->| 3268 LDAP Global Catalog
           |        '---
           |          |
           |        .---
           '--SSL-->| 3269 LDAPS Global Catalog 
                    '---
                      |
                      '-------------
                                          LDAP ports
```
<a name="UWm7k"></a>
### ADWS
AD web服务，基于soap消息查询和操作对象，和ldap兼容，本质是dc在内部进行ldap查询。
```markdown
                              .---------------------------------------
                              |          Domain Controller
                            ,---
                            | 389 (Domain) <------------.
                            '---                        |    .------.
                              |                         |----| LDAP |
                            .---                        |    '------'
                            | 3268 (Global Catalog) <---'       |
                            '---                                ^
                              |                                 |
 .------.     .------.      .---                                |
 | ADWS |>--->| SOAP |>---->| 9389  >----------------->---------'
 '------'     '------'      '---
                              |
                              '---------------------------------------
ADWS related ports and protocols
```
<a name="M7S2S"></a>
## DHCP
DHCP是一种在 UDP 上工作的应用程序协议。它在服务器中使用端口 67/UDP，并要求客户端从端口 68/UDP 发送消息。
```markdown
客户端服务器
 -----. .----- 
      | | 
     ---。.------. .--- 
 68/UDP |>---->| DHCP |>---->| 67/UDP 
     ---' '------' '--- 
      | | 
 -----''-----
																				DHCP 端口
```
<a name="DVORO"></a>
### DHCP 过程(DORA):

1. 服务器发现:客户端向255.255.255.255或者网络广播地址发送广播请求来请求IP地址,DHCP服务器会接受此请求.
1. IP租用提议:DHCP服务器提供一个IP和配置选项进行广播回应.
1. IP租用请求:客户端收到IP租用提议后,确认使用该IP和接受对应的配置选项,向DHCP服务发出使用该IP的请求.
1. IP租用确认:DHCP服务器确认该IP可以被客户端使用.
```markdown
  客户端服务器
    | | 
    | 发现| 
    | ----------> | 
    | | 
    | 提议| 
    | <--------- | 
    | | 
    | 请求 | 
    | ----------> | 
    | | 
    | 确认| 
    | <--------- | 
    | |
    																				DHCP流程
```
<a name="sOSJ3"></a>
## NetBIOS
netbios分为三个服务,NetBios Name Service用于解析NetBios名称,和NetBios Datagram和NetBios Session用于传输信息.
```markdown
                       .-----
                       |
     .------.        .---
     | NBNS |--UDP-->| 137
     '------'        '---
                       |   
    .-------.        .---
    | NBDGM |--UDP-->| 138
    '-------'        '---
                       |
    .-------.        .---
    | NBSSN |--TCP-->| 139
    '-------'        '---
                       |
                       '-----
										NetBIOS ports
```
<a name="hvfea"></a>
### NetBIOS Session Service
与TCP类似,面向连接的消息传输,使用139/TCP端口.
<a name="gHdjO"></a>
### NetBIOS Datagram Service
与UDP类似,用作无连接的应用程序协议消息传输,使用138/UDP端口.
<a name="TVdLi"></a>
### NetBIOS Name Service
该服务用于将NetBIOS名称解析为IP地址,提供节点的状态,注册或者删除NetBIOS名称,使用端口137/UDP.<br />
<br />查看本地NetBIOS名称
```powershell
C:\Users\fibr3>nbtstat -n

VMware Network Adapter VMnet8:
节点 IP 址址: [192.168.96.1] 范围 ID: []

                NetBIOS 本地名称表

       名称               类型         状态
    ---------------------------------------------
    MYLAPTOP       <00>  唯一          已注册
    WORKGROUP      <00>  组           已注册

```
[常见类型参考](https://web.archive.org/web/20031010135027/http://www.neohapsis.com:80/resources/wins.htm#sec2)
<a name="IJHvA"></a>
## WPAD
web代理自动发现(wpad),浏览器动态获得一个代理文件来确定使用的代理协议,默认情况下不使用该协议,但可以在浏览器或者系统设置中设置,或者通过GPO设置.
<a name="nUZha"></a>
## ACL
安全描述符，每个对象在其NTSecurityDescriptor属性都有一个关联的安全描述符，以二进制格式存储

- 对象主体的sid
- 所属组的sid
- DACL(自由访问控制列表）(可选)
- SACL(系统访问控制列表）(可选)
<a name="nqu99"></a>
### ACE
访问控制条目，ACL即ACE的列表.ACE分为几个部分:

- 类型:表明是拒绝还是允许访问
- 可继承性:该ace是否被继承
- 权限:ace正在应用的访问类型
- 主体:应用ace的主体sid
- 对象类型:扩展权限，属性，或者子对象的guid
- 继承类型: 



