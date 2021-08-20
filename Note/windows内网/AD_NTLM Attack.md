<a name="IJjUW"></a>
# NTLM 协议

---

NTLM(NT LAN Manager)是一种身份验证协议,可用于本地计算机验证和域用户验证.
<a name="ne8ax"></a>
##### NTLMv1(Net-NTLMv1)
NTLM协议主要用于服务器端和客户端挑战应答模式中,v1版本同时使用NT和LM哈希,具体由配置决定.只存在于很老的系统中.
<a name="cc7kP"></a>
##### NTLMv2(Net-NTLMv2)
现在系统主要使用的NTLM协议版本,算法不同,更难进行破解.
<a name="IhgvV"></a>
##### NTLM2
安全性比NTLMV2弱的但比NTLMV1强的版本.
<a name="E0kyo"></a>
##### NTLM Hash
对服务器质询的响应,由NT HASH进行计算,也叫NET-NTLM 哈希或者NTLM响应.
<a name="Jn3ED"></a>
##### NTLMv1 Hash
由NTLMV1创建的Hash.
<a name="x1gCC"></a>
##### NTLMV2 Hash 
由NTLMV2创建的Hash.
<a name="dNILV"></a>
##### NT Hash
由用户密码生成的Hash
<a name="Qqxlj"></a>
##### LM Hash
老版本LAN Manager生成的Hash.
<a name="EckhC"></a>
## NTLM 身份验证
验证包括三部分,协商,挑战,认证.<br />NTLM 验证过程
```
client               server
  |                    |
  |                    |
  |     NEGOTIATE      |
  | -----------------> |
  |                    |
  |     CHALLENGE      |
  | <----------------- |
  |                    |
  |    AUTHENTICATE    |
  | -----------------> |
  |                    |
  |    application     |
  |      messages      |
  | -----------------> |
  | <----------------- |
  | -----------------> |
  |                    |
```

1. 协商:客户端初始完上下文后调用NTLM SPP的InitializeSecurityContext 来发送一个协商信息(NEGOTIATE Message)到服务端,包含了一些安全选项,比如NTLM版本等.

![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627402619898-2562f4a5-3dde-4b3a-855b-1b14868de4e4.png#clientId=u9e933a68-158e-4&from=paste&height=698&id=u276a9462&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1396&originWidth=3825&originalType=binary&ratio=1&size=347542&status=done&style=none&taskId=u65a98054-81b2-4b92-a0f5-9bb96bf958a&width=1912.5)

2. 挑战:服务端通过调用NTLM SSP的AcceptSecurityContext来生成一个挑战(Challenge),将其包含在CHALLENGE信息中发送给客户端,还发送包括域名和NTLM版本等确认信息.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627403037561-7f0e3f74-85fa-4b71-b4e3-f223d0300291.png#clientId=u9e933a68-158e-4&from=paste&height=788&id=u17914292&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1576&originWidth=3710&originalType=binary&ratio=1&size=339524&status=done&style=none&taskId=ua828da30-51a3-4c79-9d79-64baa824737&width=1855)
2. 认证:客户端收到挑战信息后,将其发送给InitializeSecurityContext来通过用户密码生成的NT Hash来计算该挑战信息得到对应的响应信息,如果需要的话还会生成一个会话密钥(Session Key)用来加密后续客户端和服务端的通信内容,该会话密钥会由用户的NT Hash加密传输,然后客户端将响应信息和加密后的会话密钥发送回服务端,其中还包括一个由发送内容计算出的MIC信息,用来防止信息篡改.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627403346515-4283a339-5704-47af-ad89-0d2654eefa47.png#clientId=u9e933a68-158e-4&from=paste&height=766&id=u4a2fa0ad&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1531&originWidth=3712&originalType=binary&ratio=1&size=337318&status=done&style=none&taskId=uba543071-ed36-4851-be2c-959f240a7e6&width=1856)
2. 协商结果:最后服务端验证响应是否正确,并调用AcceptSecurityContext来设置安全上下文.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627403571913-273a6f19-fdc9-467c-a1b1-e9c2c35b0b2b.png#clientId=u9e933a68-158e-4&from=paste&height=621&id=udbf4480a&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1241&originWidth=3801&originalType=binary&ratio=1&size=292182&status=done&style=none&taskId=u44a3c241-02b8-4e4f-b717-0d9440ce4f3&width=1900.5)

_NTLM 认证流程_
```
                         client               server
                           |                    |
 AcquireCredentialsHandle  |                    |
           |               |                    |
           v               |                    |
 InitializeSecurityContext |                    |
           |               |     NEGOTIATE      |
           '-------------> | -----------------> | ----------.
                           |     - flags        |           |
                           |                    |           v
                           |                    | AcceptSecurityContext
                           |                    |           |
                           |                    |       challenge
                           |     CHALLENGE      |           |
           .-------------- | <----------------- | <---------'
           |               |   - flags          |
       challenge           |   - challenge      |
           |               |   - server info    |
           v               |                    |
 InitializeSecurityContext |                    |
       |       |           |                    |
    session  response      |                    |
      key      |           |    AUTHENTICATE    |
       '-------'---------> | -----------------> | ------.--------.
                           |   - response       |       |        |
                           |   - session key    |       |        |
                           |     (encrypted)    |   response  session
                           |   - attributes     |       |       key
                           |     + client info  |       |        |
                           |     + flags        |       v        v
                           |   - MIC            | AcceptSecurityContext
                           |                    |           |
                           |                    |           v
                           |                    |           OK
                           |                    |
```
<a name="gz88E"></a>
## v1和v2差异
主要提现在安全性提高,加入了时间戳和一个包含了域名,主机名等Avparis的字段,在认证消息中含有一个MIC值来检验消息是否被篡改,但在V1中响应并不根据标志位来决定,所以在NTLMv1中只要把MIC标识位移除便可以修改消息,在V2中会根据标志位强制检验MIC值.
<a name="IGNBM"></a>
# AcitveDiretory中的NTLM认证
在AD中使用NTLM认证的话则服务端需要向DC请求验证,因为用户的NT哈希只存在DC中,服务端向DC发送一个Netlogon请求来请求DC验证客户端返回的NTLM响应是否正确,DC验证成功之后返回一些相关信息,如会话密钥(Session Key)来加密后续客户端和服务端的通信.
```
  client            server                          DC
    |                 |                              |
    |                 |                              |
    |    NEGOTIATE    |                              |
    | --------------> |                              |
    |                 |                              |
    |    CHALLENGE    |                              |
    | <-------------- |                              |
    |                 |                              |
    |   AUTHENTICATE  |  NetrLogonSamLogonWithFlags  |
    | --------------> | ---------------------------> |
    |                 |                              |
    |                 |        ValidationInfo        |
    |                 | <--------------------------- |
    |                 |                              |
    |   application   |                              |
    |    messages     |                              |
    | --------------> |                              |
    |                 |                              |
    | <-------------- |                              |
    |                 |                              |
    | --------------> |                              |
    |                 |                              |
    																NTLM process with domain accounts
```
<a name="TYb0E"></a>
## 强制使用NTLM
AD下默认使用Kerberos认证,但是通过使用IP的方式而非主机名的方式则可以强制使用NTLM认证,因为Kerberos需要根据主机名来确定主机服务<br />Exp:
```powershell
dir \\dc01\C$ --->Kerboers
dir \\192.168.7.1\C$ --->NTLM
```
<a name="GoMef"></a>
## NTLM->信息收集
在服务端向客户端返回挑战信息时,AvParis的TargetInfo字段含有服务器的一些信息,如主机名,域名等信息.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627403037561-7f0e3f74-85fa-4b71-b4e3-f223d0300291.png#clientId=u9e933a68-158e-4&from=paste&height=788&id=MttAz&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1576&originWidth=3710&originalType=binary&ratio=1&size=339524&status=done&style=none&taskId=ua828da30-51a3-4c79-9d79-64baa824737&width=1855)<br />[NTLMRecon](https://github.com/pwnfoo/NTLMRecon)<br />[ntlm-info](https://gitlab.com/Zer1t0/ntlm-info)
<a name="VhkQB"></a>
## Pash The Hash
因为NTLM认证中使用的是由用户密码产生NT Hash进行计算的,所以只要得到NT Hash即使不知道明文密码也可以完成认证从而模拟该Hash用户的权限.
<a name="Zrr8e"></a>
### NT Hash提取

- lsass.exe
- SAM.hive
- ntds.dit
<a name="xSOrs"></a>
### 限制条件
<a name="Z9KE1"></a>
#### Uac
远程执行任务时,如果账户是Administrator组中的域用户,则不会触发UAC限制,如果是本地Administrator组中的本地管理员用户则会启用UAC限制.
<a name="ptdqe"></a>
##### LocalAccountTokenFilterPolicy
位于`HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`,值为0或1,默认不存在即为0.

- 如果为0(默认情况),只有内置rid为500的本地管理员用户能够远程不受UAC限制执行任务,其它用户只能使用受限的访问令牌.
- 如果为1,则Administrator组的所有用户都可以远程不受UAC限制执行任务.
<a name="UVQ4o"></a>
##### FilterAdministratorToken
位于`HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`,值为0或1,默认为0.

- 如果为0,则表示内置rid为500的管理员能够不受UAC限制远程执行任务,不影响其它用户,只影rid为500的用户.
- 如果为1,则表示内置rid为500的管理员也会受UAC限制远程执行任务,除非`LocalAccountTokenFilterPolicy`为1.

​

矩阵(粗体为默认):<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628439264738-89b5de2c-4448-44bb-b2f1-5d3e59a2b7af.png#clientId=u40f70810-e33a-4&from=paste&id=ub8e6bc05&margin=%5Bobject%20Object%5D&name=image.png&originHeight=209&originWidth=1341&originalType=url&ratio=1&size=19882&status=done&style=none&taskId=u254bcfac-3ac8-427a-999f-d5646655159)
<a name="FhWRP"></a>
## NTLM Relay
攻击者作为中间人冒充服务器的身份与客户端进行协商验证,再冒充客户端将协商过程转发到真正的目标服务器上,从而来获得一个经过真正服务器验证的会话.
```
    client                 attacker               server
      |                       |                     |
      |                       |                -----|--.
      |     NEGOTIATE         |     NEGOTIATE       |  |
      | --------------------> | ------------------> |  |
      |                       |                     |  |
      |     CHALLENGE         |     CHALLENGE       |  |> NTLM Relay
      | <-------------------- | <------------------ |  |
      |                       |                     |  | 
      |     AUTHENTICATE      |     AUTHENTICATE    |  |
      | --------------------> | ------------------> |  |
      |                       |                -----|--'
      |                       |    application      |
      |                       |     messages        |
      |                       | ------------------> |
      |                       |                     |
      |                       | <------------------ |
      |                       |                     |
      |                       | ------------------> |
      									NTLM relay attack
```
但在认证过程中,会话密钥由用户的NT Hash进行加密,即攻击者无法获取该会话密钥,也就无法因为后续的会话过程会用该会话密钥进行签名,服务端就会拒绝这些没有签名的请求,但是否签名是可以协商的,也并不是必须.
<a name="UE6GZ"></a>
### 签名协商
在NTLM协议的协商阶段,客户端和服务端会就是否对后续会话进行签名进行协商,该标识位为NETGOTIATE_SIGN字段.该字段标识为1则标识客户端有能力进行签名,服务端如果也支持签名则也将该字段设置为1.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627918395213-86371eff-b00f-47cc-9780-3f44fa94acf8.png#clientId=u20ec19f2-2fb0-4&from=paste&id=u74a42606&margin=%5Bobject%20Object%5D&name=image.png&originHeight=463&originWidth=658&originalType=url&ratio=1&size=84085&status=done&style=none&taskId=u095f77b2-1942-46ba-990c-09351e79a8d)<br />​

​<br />
<a name="ba3QO"></a>
### 是否强制签名
具体到协议上之后,就需要根据协议对签名要求做出对应的设置,通常有两三个选项来决定.

- 禁用:不启用签名.
- 启用:提供签名功能,但双方根据协商结果决定是否签名.
- 强制:必须进行签名才能进行后续会话.
<a name="PS836"></a>
#### SMB
对于SMBV1中,服务器的签名功能默认是关闭的,对于SMBV2和更高版本来说,是开启的且会强制必须进行签名才能会话.<br />                                                                            	SMB签名矩阵<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627919957158-d9478c41-3ca9-46b4-be24-7993c33f560d.png#clientId=u6665b312-4785-4&from=paste&id=u019f8c58&margin=%5Bobject%20Object%5D&name=image.png&originHeight=496&originWidth=1114&originalType=url&ratio=1&size=82917&status=done&style=none&taskId=u8a413ab0-2e19-4b2b-85b4-950cf9db3a9)  

<a name="oLrVI"></a>
######  修改默认设置
如果要修改服务器的默认签名设置,需要修改注册表`HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services\LanmanServer\Parameters` 下的EnableSecuritySignature 和 RequireSecuritySignature键值为1.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627920576106-2e72e55a-6cc7-4f09-9af3-a3af9453d582.png#clientId=u6665b312-4785-4&from=paste&id=ubfb07df1&margin=%5Bobject%20Object%5D&name=image.png&originHeight=342&originWidth=887&originalType=url&ratio=1&size=50251&status=done&style=none&taskId=u0c982304-fe49-42df-b443-157d548c97f)<br />而对于域控制器来说,是默认需要强制签名的,该安全策略是应用于域管理员组的.在图中可以看到,当域控为SMB的服务器时是启用状态,但当为客户端时,签名则是可选的.<br />![T(0H6ABM6KS@O_H9KM2_[MH.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627921590633-dfe9a737-d818-4be8-8bae-0686e2f2d8e7.png#clientId=u6665b312-4785-4&from=drop&id=ud78e56e6&margin=%5Bobject%20Object%5D&name=T%280H6ABM6KS%40O_H9KM2_%5BMH.png&originHeight=309&originWidth=1044&originalType=binary&ratio=1&size=135621&status=done&style=none&taskId=u092428ef-5012-45bc-b951-aaae0eaf40b)<br />在SMB中,客户端首先标识自己支持签名.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627921614748-c225ffaf-f348-409e-b145-b9c7e9b4a649.png#clientId=u6665b312-4785-4&from=paste&height=105&id=u8a55551c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=209&originWidth=1310&originalType=binary&ratio=1&size=23730&status=done&style=none&taskId=u8a034b1a-b278-4360-8b81-b0bb47fce4e&width=655)<br />然后服务端标识自己也支持签名,并且强制要求签名.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627921628930-0eeed1f2-236b-4b68-ab82-1bf71d7057ae.png#clientId=u6665b312-4785-4&from=paste&height=92&id=u6aa0a7ee&margin=%5Bobject%20Object%5D&name=image.png&originHeight=184&originWidth=1579&originalType=binary&ratio=1&size=26250&status=done&style=none&taskId=u79d2dfc3-eec8-4afc-b912-5b94bd2372e&width=789.5)<br />后续则对通信内容则进行签名.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627921651402-c589faf5-f29b-4c8a-a53c-250a277057df.png#clientId=u6665b312-4785-4&from=paste&height=442&id=u2e19a40a&margin=%5Bobject%20Object%5D&name=image.png&originHeight=884&originWidth=1595&originalType=binary&ratio=1&size=100372&status=done&style=none&taskId=u839ba30c-f127-4297-b792-33469887327&width=797.5)<br />

<a name="xHOGF"></a>
#### LDAP
对LDAP来说,也有三种选项

- 禁用: 不支持签名
- Negotiated Signing:支持签名,如果对方也支持签名,则默认进行签名.
- 强制签名:必须进行签名才能通信.

和SMBV2的区别在于,当双方都支持签名时,则会默认进行签名,而SMBV2则只有当一方需要签名时才签名,而且在AD中,所有主机都有Negotiated Signing设置,但域控不是需要强制签名的.<br />LDAP矩阵<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1627921950901-1fe5d132-4072-4726-b783-55a0afed4287.png#clientId=u6665b312-4785-4&from=paste&id=u4ba184a4&margin=%5Bobject%20Object%5D&name=image.png&originHeight=429&originWidth=1108&originalType=url&ratio=1&size=55364&status=done&style=none&taskId=u6f0ef660-643f-4528-9c98-af39a3c3585)<br />在注册表中LDAP签名的域控设置位于HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services\NTDS\Parameters\ldapserverintegrity中,可以设为0,1,2.域控上默认为1.而客户端的注册表位于HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services\ldap\ldapserverintegrity默认也为1.
<a name="Pyyww"></a>
### 跨协议中继
在通信过程中,认证过程和会话是独立的,也就是说作为中间人可以通过A协议的来和某一端进行认证,同时将A协议认证中的NTLM内容封装为B协议和另一端进行B协议通信.<br />因为我们没有会话密钥,所以在中继过程中无法对消息进行签名,比如在LDAP协议中使用NTLM中继的话则有两个要求来保证不要求签名:

- 服务器不能要求签名.因为客户端默认签名,服务器要求签名的话则无法进行后续会话.
- 客户端协商时也不能设置NEGOTIATE_SIGN标识为1.如果该标志为1,即服务器也会要求进行签名.

而对于WindowsSMB客户端来说,NEGOTIATE_SIGN被设置为了1,所以默认情况无法将SMB身份验证中继到LDAP中.而且对于LDAPS来说,服务器收到NEGOTIATE_SIGN会直接终止身份验证(?).而SMB2协议也可以中继到另一个SMB2协议中,只要另一个SMBv2服务不需要签名即可(非DC).而HTTP在使用NTLM认证时默认不需要签名,即可以将该认证中继到SMB(不要求签名)或者LDAP上.
<a name="fkOW0"></a>
#### 通道绑定(EPA)(Channel Bindings)
该措施用来保护跨协议中继,将身份验证与使用的协议进行绑定,也就无法将A协议中的NTLM转发到B协议中进行验证了,思路为在最后一条NTLM验证消息中放置一条无法修改的消息用来指示请求的服务(SPN).
<a name="j2dHm"></a>
###### 服务绑定
比如在NTLM响应中标识要请求的服务为HTTP,如果将该NTLM转发到SMB服务器,SMB服务器会检验请求的该服务(HTTP)与自己(SMB)是否对应,不对应则拒绝验证,而且该请求服务标识无法被修改,因为它被计算在了NTLM响应中.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628092880475-4ede387c-71ec-4157-9e74-184a58f9b374.png#clientId=u6665b312-4785-4&from=paste&height=573&id=uf73c434c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1145&originWidth=3816&originalType=binary&ratio=1&size=257243&status=done&style=none&taskId=u03492618-8763-4ea6-9fe8-1e683e762a0&width=1908)<br />如图标识了请求的服务为cifs,请求IP为192.168.107.1,如果将此NTLM验证中继到LDAP则会失败.<br />证书绑定<br />简单来说就是如果使用了TLS协议,如LDAPS等,则计算其证书的哈希,并将其加入NTLM响应中进行计算,攻击者将该NTLM验证信息转发到其它服务时,则会因为证书Hash不匹配从而拒绝验证.
<a name="ISuJT"></a>
#### 中继矩阵
![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628093279502-9adcadc6-29b4-4852-9f80-279b816c6418.png#clientId=u6665b312-4785-4&from=paste&id=u36db3879&margin=%5Bobject%20Object%5D&name=image.png&originHeight=936&originWidth=950&originalType=url&ratio=1&size=63646&status=done&style=none&taskId=u989fa6ac-5cef-4851-8860-93b1c27bf80)
<a name="sk5Ue"></a>
#### MIC-消息完整性
MIC是在NTLM身份验证最后的AUTHENTICATE消息中的签名发送,该签名使用HMAC_MD5算法并使用会话密钥作为Key进行对协商信息和挑战信息和认证信息进行加密得到.因为会话密钥是由用户的NT Hash加密的,所以无法得到该会话密钥,也就无法篡改MIC.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628089924435-f672156b-c59d-47b8-a149-a5e8acbdd64f.png#clientId=u6665b312-4785-4&from=paste&height=676&id=uad3de0d8&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1351&originWidth=3791&originalType=binary&ratio=1&size=344225&status=done&style=none&taskId=ud9c6fa7b-e93a-4350-9f20-53ff624b89c&width=1895.5)<br />但MIC标识是可选的,并不是必须的,但也无法直接将MIC标识删除来篡改消息,因为还有一个标志位来表明MIC是否必须存在,即msAvFlags,如果值为0x00000002的话则向服务器表明必须存在MIC,如果服务器收到了该标识而没有见到MIC消息的话则会终止验证.<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/12610959/1628090146332-7569a8de-5850-4166-a7a0-52f51c63b100.png#clientId=u6665b312-4785-4&from=paste&height=657&id=u82542672&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1314&originWidth=3680&originalType=binary&ratio=1&size=281114&status=done&style=none&taskId=u57e661b8-e041-447a-b0f5-e6d51ef6a7e&width=1840)<br />但还是无法直接将msAvFlags设为0,然后将MIC标识直接删除,因为在NTLMv2中,生成的响应中还加入了所有的标志位信息,即指示了MIC必须存在,如果修改了相应的标志位,那么该NTLMv2 Hash将会无效.
<a name="hZFAs"></a>
##### 相关漏洞
<a name="Uocbm"></a>
##### [Drop MIc](https://www.preempt.com/blog/drop-the-mic-cve-2019-1040/)​
[Drop MIC 2](https://www.preempt.com/blog/drop-the-mic-2-active-directory-open-to-more-ntlm-attacks/)
<a name="F8ney"></a>
##### 该漏洞中,即使移除MIC,标志位标识其存在,服务器也不会拒绝身份验证.
<a name="fIy19"></a>
## 会话密钥(Session Key)
<a name="wx0jI"></a>
##### 不同版本Session Key的计算方式
```
# For NTLMv1
Key = MD4(NT Hash)

# For NTLMv2
NTLMv2 Hash = HMAC_MD5(NT Hash, Uppercase(Username) + UserDomain)
Key = HMAC_MD5(NTLMv2 Hash, HMAC_MD5(NTLMv2 Hash, NTLMv2 Response + Challenge))
```
对于本地登录来说,由本地服务器直接计算对应的会话密钥,而对域账户登录来说,则必须由域控来进行计算会话密钥,并将该会话密钥传输返回.
<a name="LowmZ"></a>
##### 相关漏洞
[CVE-2015-0005](https://www.coresecurity.com/core-labs/advisories/windows-pass-through-authentication-methods-improper-validation)<br />该漏洞中域内机器可不经过验证获取任意其它服务器的会话密钥.
