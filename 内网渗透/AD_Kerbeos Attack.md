# Kerberos V5
## Basic

1. Kerberos v5协议是默认的域认证协议,当某个计算机不支持Kerberos v5协议时则使用NTLM协议.
1. Kerberos v5协议支持对称和受限制的非对称加密.
1. Kerberos协议根据主体(principal)来对用户进行身份验证.
1. Kerberos协议由DC在88/TCP和88/UDP进行监听,同时Kerberos还提供更改域用户密码的服务(kpasswd),该服务监听在DC的464/tcp和464/udp端口.
```
                           .-----
                           |
                         .---
            .----KDC---> | 88
            |            '---   Domain
 Kerberos --|              |
            |            .---  Controller
            '-kpasswd--> | 464
                         '---
                           |
                           '-----
```


大致流程:

1. 用户在客户端使用密码或者智能卡向KDC请求认证.
1. KDC向客户端分发一张ticket-granting ticket(TGT),客户端使用这张TGT来访问ticket-granting service(TGS).
1. 然后TGS分发一张服务票据给客户端.
1. 客户端凭借这个服务票据来请求网络服务,该服务票据提供了用户和该服务的双向认证.

流程图<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1626540933480-31600e40-d552-4379-8255-cbe1b9f528ac.png#height=465&id=tlEPi&margin=%5Bobject%20Object%5D&name=image.png&originHeight=929&originWidth=1427&originalType=binary&ratio=1&size=305920&status=done&style=none&width=713.5)
## 认证摘要所使用的安全组件

- Kerberos.dll : 提供了一系列工业标准的安全协议和服务认证的方法
- Kdcsvc.dll : KDC服务
- Ksecdd.sys : 提供用来和LSASS在用户模式下进行交互的内核安全驱动
- Lsasrv.dll : LSA服务,实施安全策略和作为LSA的安全包管理
- Secur32dll : 用户模式下作为提供认证的一部分提供安全接口访问
<a name="C2Z3v"></a>
## Ticket
Kerberos协议中有两种类型的票据,ST(TGS)和TGT,同时也有两种格式:ccache和krb,分别为Linux和Windows,可以使用[ticket_converter](https://github.com/eloypgz/ticket_converter)进行转换.<br />Tikcket的部分结构为加密的,包括

- 该票据对应的主体
- 客户端的名称和域等
- 密钥
- 有效期时间戳
<a name="cnbXm"></a>
### PAC(特权属性证书)
通常包含在票据的`authorization-data`字段中,PAC包含了客户端相关的安全属性等,包括

- 客户端域的相关信息(域名和域SID)
- 客户端用户的用户名和RID
- 客户端在域中所属组
- 一些没有域的组SID,用于域间的身份验证,如Well-Known SID.
- 用加密票据的密钥生成的PAC内容的服务i签名
- 由KDC 密钥生成的KDC签名,用来验证该PAC是由KDC生成和预防银票据攻击,但是没有对应的检验.
- 由KDC 密钥生成的票据签名用来防止Bronze Bit 攻击.
<a name="enpRm"></a>
### ST(TGS)(Service Ticket)
服务票据(Serivce Ticket),客户端把它发送给要请求的服务主机,然后服务主机验证通过后提供相应的服务,而KDC负责向客户端发送该票据.在域中的客户端可以得到注册在域数据库中的服务ST,无论该用户是否有权限访问,或者该服务是否运行.而且ST是由该服务账户的密码生成的密钥进行加密.
<a name="u2kQs"></a>
### TGT
正常认证过程中,我们需要向KDC来获得ST,而为了证明请求ST的权限,需要先提供另一种票据,即TGT(Ticket Granting Ticket)
```
																										Kerberos process
                       KDC (DC)
   .----->-1) AS-REQ------->   .---.
   |                          /   /| -------8] PAC Response-----------.
   | .--<-2) AS-REP (TGT)--< .---. |                                  |
   | |                       |   | '                                  |
   | | .>-4) TGS-REQ (TGT)-> |   |/  <-7] KERB_VERIFY_PAC_REQUEST-.   |
   | | |                     '---'                                |   |
   | | | .<-5) TGS-REP (ST)--<'                                   |   |
   | | | |                                                        |   v
   | v | v                                                        ^   
   ^   ^                                                          .---.
    _____                                                        /   /|
   |     |   <----3) Authentication negotiation (SPNEGO)---->   .---. |
   |_____|                                                      |   | '
   /:::::/   >-------------------6) AP-REQ (ST)------------->   |   |/ 
   client                                                       '---'  
             <-------------------9] AP-REP------------------<  AP (service)
```
<a name="KeWaw"></a>
## 认证流程
<a name="ILiVo"></a>
### AS-REQ
用户向kdc发送AS-REQ来请求TGT,默认还包含了一个加密的时间戳,为预认证部分.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628523588796-5b82b98a-7c13-4d90-b369-aacb441f331d.png#clientId=u4ebbf284-a4a6-4&from=paste&height=573&id=ue4cb1def&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1145&originWidth=2537&originalType=binary&ratio=1&size=159523&status=done&style=none&taskId=ua22082f2-43ec-4324-835c-86a7c38b694&width=1268.5)
<a name="SoZzR"></a>
### 预认证
在客户端发送AS-REQ之后,KDC返回的AS-REP中包含了KDC密钥加密的TGT和客户端密钥加密的客户端数据,为了防止暴力破解所以在客户端AS-REQ请求时默认需要进行预认证,即客户端在发送AS-REQ之中需要用自己密钥加密的时间戳,然后KDC成功解密后才返回AS-REP,但该选项可以在账户的属性中关闭,<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628523063788-dbdc593d-3560-4950-8a10-22ca5d843876.png#clientId=u4ebbf284-a4a6-4&from=paste&height=441&id=ueb420f2c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=881&originWidth=2499&originalType=binary&ratio=1&size=177778&status=done&style=none&taskId=u74257ba1-cc25-4d41-8a91-22818e8c001&width=1249.5)
<a name="SwhdR"></a>
### AS-REP
KDC接受到请求后,如果预认证通过,则返回包含两个加密部分的AS-REP响应,一部分为KDC密钥加密的TGT和客户端密钥加密的客户端数据.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628523638076-776368dc-b7c2-460d-b89e-b7f16c21428f.png#clientId=u4ebbf284-a4a6-4&from=paste&height=516&id=uffe5fbf8&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1031&originWidth=2552&originalType=binary&ratio=1&size=143872&status=done&style=none&taskId=u6f3748d0-a004-490f-b0cc-409b41e80b7&width=1276)
<a name="pVPgv"></a>
### TGS-REQ
客户端发送上一步得到TGT和请求目标的SPN给KDC请求对应的ST,同时包含一些使用会话密钥加密的客户端数据和时间戳,来验证连接.

![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628524072939-85083046-8b65-4b6c-87dd-f19909c36358.png#clientId=u4ebbf284-a4a6-4&from=paste&height=574&id=u892cb0d5&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1147&originWidth=2534&originalType=binary&ratio=1&size=145537&status=done&style=none&taskId=u7b3c7a1d-f833-4744-b928-d5c9670e1ee&width=1267)


### TGS-REP
KDC解密得到的TGT,得到用户名和发送的用户名进行对比验证,通过验证之后,KDC返回包含两个加密部分的TGS-REP响应,加密部分分别为服务账户密钥加密的目标服务的ST和会话密钥加密的客户端数据.![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628524591883-7c988b64-8c59-42cc-93e0-fba8b45e86bc.png#clientId=u4ebbf284-a4a6-4&from=paste&height=304&id=u57636767&margin=%5Bobject%20Object%5D&name=image.png&originHeight=607&originWidth=1209&originalType=binary&ratio=1&size=51134&status=done&style=none&taskId=uffe1489d-5768-4804-9ee4-fddc2eff12c&width=604.5)

### AP-REQ
最后客户端拿着得到ST再次发送AP-REQ去请求对应的服务.

![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628525272894-99baba72-51a4-4511-aa26-dbaeac4fb2eb.png#clientId=u4ebbf284-a4a6-4&from=paste&height=499&id=u0543c504&margin=%5Bobject%20Object%5D&name=image.png&originHeight=997&originWidth=2367&originalType=binary&ratio=1&size=137917&status=done&style=none&taskId=u82b3e349-8bd8-421b-aa30-8a7837c46ba&width=1183.5)

### 可选
验证PAC:如果服务要验证ST中的PAC,那么服务可以使用NETLogon协议向域控验证该PAC签名是否正确.

服务验证:客户端还可以要求服务发送一个AP-REP响应(使用会话密钥加密)来证明自己是真的服务.

## 跨域Kerberos
跨域认证流程
```
  KDC foo.com                                                    KDC bar.com
    .---.                                                          .---.
   /   /|                       .---4) TGS-REQ (TGT bar)------->  /   /|
  .---. |                       |    + SPN: HTTP\srvbar          .---. |
  |   | '                       |    + TGT client > bar.com      |   | '
  |   |/                        |                                |   |/ 
  '---'                         |   .--5) TGS-REP--------------< '---'
  v  ^                          |   | + ST client > HTTP/srvbar
  |  |                          |   |
  |  |                          ^   v                                   .---.
  |  '-2) TGS-REQ (TGT foo)--<  _____                                  /   /|
  |   + SPN: HTTP\srvbar       |     | <----------1) SPNEGO---------> .---. |
  |   + TGT client > foo.com   |_____|                                |   | '
  |                            /:::::/ >----6) AP-REQ---------------> |   |/
  '--3) TGS-REP--------------> client     + ST client > HTTP/srvbar   '---'  
    + TGT client > bar.com    (foo.com)                               srvbar
                                                                    (bar.com)
```

1. 客户端和想要请求的服务器(HTTP/srvbar)协商协议,使用kerberos进行认证.
1. 客户端使用自己的TGT向自己所在域的kdc(foo.com)发送TGS-REQ来请求对应服务(HTTP/srvbar)的ST.
1. 客户端所在域的KDC(foo.com)识别到请求的服务在另一个信任域(bar.com)中,然后使用域间共享密钥为客户端创建一个另一个域(bar.com)的TGT票据.同时包含一份foo.com域TGT相同的PAC.
1. 客户端使用第二个TGT向服务所在域的kdc(bar.com)发送TGS-REQ请求来请求HTTP/srvbar服务的ST.
1. bar.com域kdc使用域间共享密钥解密TGT后为客户端创建请求的ST.
1. 最后客户端得到ST请求对应的服务.
### Kerberos keys
在kerberos协议中作为用户认证的代表,由用户的密码生成.该key可以代表用户申请Kerberos票据.

根据不同算法生成不同的key,Kerberos协议支持的Key类型:

- AES 256 Key: 由[AES256-CTS-HMAC-SHA1-96](https://tools.ietf.org/html/rfc3962) 算法使用,也是Kerberos协议最常采用的Key.
- AES 128 Key: 由[AES128-CTS-HMAC-SHA1-96](https://tools.ietf.org/html/rfc3962) 算法使用.
- DES Key: 由被遗弃的[DES-CBC-MD5](https://datatracker.ietf.org/doc/html/rfc3961#section-6.2.1) 算法使用.
- RC4 Key: 由RC4 算法生成的用户NT hash.

该Key可被用于Pass the key攻击来模拟用户请求票据,同时可以使用[cerbero_hash](https://gitlab.com/Zer1t0/cerbero#hash)来根据用户密码计算其Kerberos密钥.
### Other key

- Long-term key: 只有目标服务和KDC知道的密钥,在客户端的票据中加密存储,该key由明文密码经过加密函数后产生,所有Kerberos v5认证实现都必须支持DES-CBC-MD5.
- Client/Server session key: 一个临时单一的会话密钥,当认证通过后,KDC创建该密钥用来加密客户端和服务器之间的通信.
- timestamps: 时间戳用来防止包重放攻击.
- User Keys: 当用户创建时,使用该用户的密码来生产该Key,在AD域中,user key与用户对象一起存储在AD域中,在工作站上,当用户登录时该Key被创建.
- System Keys: 当工作站或者服务加入域时,由该系统账户的密码生成该Keys.
- Service Keys: 由krbtgt账号的密码生成,同一个域的所有域控都使用同一个Service Key.
- Inter-relam Keys:KDCS共享的密钥,主要用于跨域验证,具有父子关系的Actice Directory域之间也共享域间密钥,该密钥使用信任账户的密码进行使用RC4算法加密得到.
## Kerberos Security Support Provider(SSP)

- 安全支持提供,在kerberos中使用动态链接库(DLL)作为SSP,在系统启动时由Local Sercurity Authority(LSA)加载.
- SSPI即实现SSP的接口

默认的SSPs包含了Negotitate(SPNEGO),Kerberos,NTLM,Schannel,Digest Authtication protocals(认证摘要协议),都被当做插件一样通过DLL的形式包含在了SSPI中.

SSPI结构图:

![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1626541307400-42a306f0-f3e3-4cb6-9c4c-ea760537e442.png#height=489&id=Jrwfd&margin=%5Bobject%20Object%5D&name=image.png&originHeight=978&originWidth=2101&originalType=binary&ratio=1&size=303338&status=done&style=none&width=1050.5)

通过SSPI,当两个应用之间需要认证时,可以不管使用的协议,而直接通过使用统一的SSPI来完成认证过程.

   - Digest Authtication: 工业标准协议,主要用于LDAP和web认证,通过MD5或者信息摘要在网络中传递凭证.
   - Schannel: 实现了SSL和TLS的互联网标准认证协议,主要用于Web Server认证.
   - Negotiate: 用来协商指定的认证协议,如果一个应用指定该SSP进行认证,则Negotiate会根据请求和用户配置的安全策略来选择最好的SSP来处理该请求.

支持Kerberos SSP的域服务包括:

- LDAP
- 使用RPC调用的远程管理
- 打印服务
- C/S认证
- 远程文件访问(CIFS/SMB)
- 文件分发管理系统
- IIS内网认证
- IPSec
- 用户或者计算机通过证书请求的对应服务证书
### Kerberos Physical Structure
Kerberos认证主要使用对称密钥,加密对象,Kerberos服务.

Kerberos要素:

![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1626546058858-f59c42e2-a40a-4959-9e80-f9599e862f06.png#height=412&id=CoRE9&margin=%5Bobject%20Object%5D&name=image.png&originHeight=823&originWidth=870&originalType=binary&ratio=1&size=156129&status=done&style=none&width=435)

## Attack Kerberos
### Kerberos brute-force
暴力破解用户密码,Kerberos的几种错误类型:

- KDC_ERR_PREAUTH_FAILED：密码错误
- KDC_ERR_C_PRINCIPAL_UNKNOWN：用户名无效
- KDC_ERR_WRONG_REALM：无效域
- KDC_ERR_CLIENT_REVOKED：禁用/阻止用户

Tips:身份验证错误时会按照预认证失败(日志4771)进行记录,而不是正常的登录失败(4625).
### Kerberoast
因为AD中的任何用户都可以通过SPN请求域数据库中注册的服务的ST.而ST是由运行该服务的账户密码派生的kerberos密钥进行部分加密,而大部分服务是用机器账户运行的,而机器账户的密码是随机生成的120字符密码,且每月更新,所以破解该密钥机率很低,但有些服务是由普通用户账户运行,如果该账户密码强度不高的话,则可能会被破解得到该用户密码.请求ST时,使用RC4算法的话破解难度会更低. 
### ASREProast
如果某个账户关闭了预认证,那么就可以没有身份凭证的情况下通过发送方AS-REQ请求来让KDC返回包含该账户密钥加密的AS-REP响应,然后直接暴力破解该响应来获得该账户密码.
### Pass-the-key/over-pash-the-hash
和NTLM协议类似,Kerberos认证过程中也不是传输明文密码,而是使用的其由密码派生的Kerberos密钥,如果获得了Kerberos密钥即可模拟该用户请求TGT.

在Windows中Kerberos密钥缓存在LSASS进程中,使用`mimiakatz sekurlsa:ekeys`可以获得其缓存,或者转储LSASS进程进行离线破解获取密钥,然后使用[Rubeus_asktgt](https://github.com/GhostPack/Rubeus#asktgt)来请求TGT.

### Pass-the-Ticket
如果可以直接获得相应的TGT或者ST则可以直接使用该Ticket进行请求,在Windows中,Ticket同样缓存在lsass进程中,使用`mimikatz sekurlsa::tickets`可进行提取,或者转储lsass后使用mimikatz或者pypykatz进行离线提取,然后使用`mimikatz kerberos::ptt`或者[Rubeus_ptt](https://github.com/GhostPack/Rubeus#ptt)将Ticket注入进lsass进程中.即可模拟该Ticket的权限.
### 金票据
类似于访问KDC的ST.所有与ST类似,TGT也使用KDC的内置krbtgt账户密码生成的密钥进行加密,如果得到了krbtgt账户的密码则可以自己生成对应的TGT,即金票据,因为凭借该票据可以向KDC请求任何服务的ST票据.同时必须在20分钟之内使用,不然20分钟之后KDC会检测PAC签名,该票据将失效.
### 银票据
TGS是由提供服务的账户密码派生的密钥进行加密,所以只要知道密钥或者该账户密码,我们就能自己生成一张TGS来访问该账户提供的所有服务.即银票据,同时如果该用户运行了多个账户,因为使用的同一个密钥进行加密,则可以通过修改票据的sname(未加密)来直接访问该账户运行的另一个服务,虽然PAC中含有KDC签名,但其实服务并不会检验该签名.
### SID History attack
#### SID-History
该属性用于域间迁移,当A域的对象迁移到B域时,对象会创建一个新的SID,而旧的SID会被添加到该对象的SID History熟悉中.

在A域和B域跨域认证时,A域会创建一个B域的TGT,其中该TGT中会复制A域TGT的PAC,而在PAC中有一个额外的SID字段,而该字段包含了SID History属性的值. 这样当用户访问旧域资源时,旧域可以在PAC中获取到SID History,从而授予用户对应的访问权限.

而SID还有过滤策略,ST中的PAC不会复制该额外的SID,默认情况下域允许同一个林中的其它域的SIDS,但不会允许其它林中的SIDS,根据安全边界规则,林为AD域的安全边界.

所以在制作黄金票据时,可以在PAC的额外字段添加SIDS,如Enterprise Admins的SID,该组用户拥有林中所有域的管理员权限.当用该TGT进行跨域认证时,该TGT的PAC会被复制到新的TGT中,导致在另一个域中也获得管理员权限.

### Kerberos委派
委派通常用于以下场景,用户请求某项服务A时,在该服务过程中还需要请求服务B,这时则需要服务A以用户的身份来请求服务B,如用户请求Web服务时,而Web服务过程中需要请求数据库服务,所以需要委派Web服务以用户身份访问数据库.
#### 无约束委派
服务A其使用用户A服务内ST内包含的TGT来模拟用户访问KDC请求服务B的ST.

流程如下

```
																			kdc
 .-----------1) TGS-REQ-------------> .---. <--------6) TGS-REQ-------------.
 |         + SPN: HTTP/websrv        /   /|     + SPN: MSSQLSvc/dbsrv       |
 |         + TGT client             .---. |     + TGT client - FORWARDED    |
 |                                  |   | '                                 |
 |  .--------2) TGS-REP-----------< |   |/ >--------7) TGS-REP-----------.  |
 |  |  + ST client > HTTP/websrv    '---'   + ST client > MSSQLSvc/dbsrv |  |
 |  |    - OK-AS-DELEGATE           ^   v                                |  |
 |  |    - FORWARDABLE              |   |                                |  |
 ^  v                               |   |                                |  |
  _____                             |   |                                |  |
 |     | >-----3) TGS-REP-----------'   |                                |  |
 |_____|  + SPN: krbtgt/domain.local    |                                |  |
 /:::::/  + TGT client                  |                                |  |
 ------                                 |                                |  |
 client  <-----4) TGS-REP---------------'                                |  |
   v        + TGT client - FORWARDED                                     v  ^
   |                                                                    .---.
   |                                                                   /   /|
   |                                                                  .---. |
   '----------------------------5) AP-REQ---------------------------> |   | '
                          + ST client > HTTP\websrv                   |   |/
                          + TGT client - FORWARDED                    '---'
                                                                      websrv
                                                                       v
                                  .---.                                |
                                 /   /|                                |
                                .---. | <--------8) AP-REQ-------------'
                                |   | '   + ST client > MSSQLSvc\dbsrv
                                |   |/
                                '---'
```

1. 客户端使用其TGT向KDC请求访问Websrv的ST.
1. KDC检查Websrv服务的运行账户是否设置了TRUSTED_FOR_DELEGATION标志位.然后KDC返回设置了`OK-AS-DELEGATE`和`FORWAEDABLE`标志位的ST.
1. 客户端检查`OK-AS-DELEGATE`标志位,表示该服务使用了委派,然后客户端向KDC请求一张具有FORWARDED标志位的TGT,用来发送给Websrv.
1. KDC返回具有FORWARDED标志位的TGT.
1. 客户端发送包含设置FORWARDED TGT的ST来访问Websrv服务.
1. Websrv需要访问数据库服务时,将使用ST中设置了FORWARDED的TGT模拟用户向KDC请求访问数据库服务的ST.
1. KDC返回数据库服务的ST.
1. Websrv服务使用返回的ST模拟用户访问数据库服务.

而因为在ST中包含了客户端的TGT,所以如果控制了Websrv那么则可以获得包含在ST的TGT进行Pass-the-ticket,对于设置了TRUSTED_FOR_DELEGATION标志位的服务账号来说,其对应的ST中都将包含TGT.如果控制了某个服务器而其账户设置了无约束委派则可以通过钓鱼等方式使客户端来访问该服务器,在客户端发送的ST中则可以得到客户端的TGT,通过该TGT即可访问其它非数据库服务.
#### 约束委派
提供一种机制,即S4U扩展,允许服务A作为客户端不使用客户端的TGT来请求服务B的ST.<br />S4U包括两种扩展:

- S4U2proxy(用户到代理)
- S4U2self(用户到自己)

通过该扩展,可以限制请求的服务为特定的服务,而不需要用户的TGT,从而防止被滥用访问其它服务.
##### S4U2proxy
该扩展允许服务使用客户端发送的ST来代表用户请求另一个服务.<br />在该扩展中服务账户A拥有一个白名单为允许委派请求的服务名单,在服务账户A的msDS-AllowedTodelegateTo属性中拥有一个SPN列表,列表中的服务表示该服务账户A可以代表客户端请求的服务,但修改msDS-AllowedToDelegateTo属性则需要SeEnableDelegationPrivilege权限.<br />显示用户该属性:`get-aduser anakin -Properties msDS-AllowedToDelegateTo`
```
                KDC
                .---. <-----2) TGS-REQ--------------------.
               /   /|    + SPN: MSSQLSvc/dbsrv            |
              .---. |    + TGT websrv$                    |
              |   | '    + ST client > http/websrv        |
              |   |/                                      |
              '---'  >-----3) TGS-REP------------------.  |
                       + ST client > MSSQLSvc/dbsrv    v  ^
                                                      .---.
  ____                                               /   /|
 |    | >-----------1) AP-REQ---------------------> .---. |
 |____|     + ST client > http/websrv               |   | '
 /::::/                                             |   |/ 
 client                                             '---'  
                                                    websrv
                .---.                                  v
               /   /|                                  |
              .---. |<-----4) AP-REQ-------------------'
              |   | '   + ST client > MSSQLSvc/dbsrv
              |   |/ 
              '---'  
              dbsrv
```
<a name="xqogk"></a>
#### 基于资源的约束委派(RBCD)
当服务A想要代替客户端请求的服务B不在同一个域内时则需要用到基于资源的约束委派,在运行服务账户B的msDS-AllowedToActOnBehalfOfOtherIdentity属性中,则列出了能够以委派的方式请求自己的账户名单,即服务B拥有一份白名单,只有白名单内的账户才能够以委托的方式代替第三方请求自身服务,而且基于资源的约束委派始终允许协议转换.
```
                                                  KDC foo.com
       .--------------1) TGS-REQ-------------------> .---.
       |          + SPN: MSSQLSvc/dbsrv.bar.com     /   /|
       |          + TGT websrv$ > foo.com          .---. |
       |          + ST client > http/websrv        |   | '
       |                                           |   |/ 
       |   .----2) TGS-REP-----------------------< '---'  
       |   |  + TGT websrv$ (client) > bar.com     ^   v
       ^   v                                       |   |
        .---.                                      |   |
       /   /| >----3) TGS-REQ----------------------'   |
      .---. |      + SPN: krbtgt/bar.com               |
      |   | '      + TGT websrv$ > foo.com             |
      |   |/                                           |
      '---'   <-----4) TGS-REP-------------------------'
      websrv        + TGT websrv$ > bar.com                      .---.
    (foo.com)                                                   /   /|
      ^  v  v                                                  .---. |
      |  |  '-------------7) AP-REQ--------------------------> |   | '
      |  |          + ST client > MSSQLSvc/dbsrv.bar.com       |   |/ 
      |  |                                                     '---'  
      |  '-------5) TGS-REQ-----------------------> .---.      dbsrv
      |      + SPN: MSSQLSvc/dbsrv.bar.com         /   /|    (bar.com)
      |      + TGT websrv$ > bar.com              .---. |
      |      + TGT websrv$ (client) > bar.com     |   | '
      |                                           |   |/ 
      '---6) TGS-REP----------------------------< '---'  
        + ST client > MSSQLSvc/dbsrv.bar.com   KDC bar.com
```
##### S4U2self
主要为了让不支持Kerberos协议的客户端进行委派,也叫协议转换.使用该扩展需要打开服务A账户的`TRUSTED_TO_AUTH_FOR_DELEGATION`标识位.客户端使用其它协议通过认证后,服务A可以通过发送包含自身的SPN和TGT的TGS-REQ请求,并在请求中指示客户端名称来向KDC请求一张自身服务的ST票据,同时KDC会检查服务账户A的msDS-AllowedToDelegateTo属性,
```
                                                                KDC
                          .---. >---2) TGS-REQ-------------->  .---.
                         /   /|     + SPN: HTTP/websrv        /   /|
  ____                  .---. |     + For user: client       .---. |
 |    | >-1) SPNEGO-->  |   | '     + TGT websrv$            |   | '
 |____|     (NTLM)      |   |/                               |   |/ 
 /::::/                 '---'  <----3) TGS-REP-------------< '---'  
 client                websrv   + ST client > HTTP/websrv
```
根据客户端和服务端相关配置KDC返回的ST也有所不同:

1. 如果客户端在Proteced groups组中,那么返回的ST将被设置为non-FORWARDABLE,即不可用此ST模拟客户端进行委托.
1. 服务账户设置了`TRUSTED_TO_AUTH_FOR_DELEGATION`标志位,返回FORWARDABLE标志位的ST.
1. 服务账户没有设置`TRUSTED_TO_AUTH_FOR_DELEGATION`标志位,但ms-AllowedToDelegateTo属性中含有SPN服务,返回ST将被设置为non-FORWARDABLE,但该ST可以用于基于资源的约束委派
```
																									基于资源S4U2self委派 
                                                          ____       
                                                          |    |
                                   .-----1) SPNEGO------< |____|
                                   |       (NTLM)         /::::/
                                   |                   client (bar)
                                   |
                                   |
  foo KDC                          v                                    bar KDC
   .---. <----2) TGS-REQ------<  .---. >---------4) TGS-REQ------------>  .---.
  /   /|  + SPN: krbtgt/bar     /   /|        + SPN: HTTP/websrv         /   /|
 .---. |  + TGT websrv$ > foo  .---. |        + For user: client        .---. |
 |   | '                       |   | '        + TGT websrv$ > bar       |   | '
 |   |/                        |   |/                                   |   |/ 
 '---'  >-----3) TGS-REP-----> '---'  <----------5) TGS-REP-----------< '---'  
 v   ^   + TGT websrv$ > bar  websrv    + TGT websrv$ (client) > foo
 |   |                         (foo)
 |   |                         v   ^
 |   |                         |   |
 |   '--<<--6) TGS-REQ---<<----'   |                  
 |  + SPN: HTTP/websrv             | 
 |  + For user: client             |     
 |  + TGT websrv$ (client) > foo   |
 |                                 |
 '----->>---7) TGS-REP---->>-------'                
   + ST client > HTTP/websrv
```
#### 禁止委派
有两种方式可以禁止某个账户被其它服务委派:

- 在该账户的UserAccountControl属性中设置NOT-DELEGATED标志.
- 将该账户添加到Proteced Users组中.
