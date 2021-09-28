# Potato系列

主要用于服务账户到system账户的权限提升,如iis,sqlserver用户.

## 0X01 Hot Potato

Hot Potato(又名：Potato)利用 Windows 中的已知问题在默认配置中获得本地权限提升,即 NTLM 中继(特别是 HTTP->SMB 中继)和 NBNS 欺骗.

### Local NBNS Spoofer

#### NBNS协议

NBNS 是 Windows 环境中常用的名称解析广播 UDP 协议。在DNS查询时,如果hosts文件中查询不到就会进行DNS查询,如果DNS查询不到则进行NBNS查询,即向本地广播域中进行广播查询"XXX的IP是多少?",并且任何主机都可以对此进行回复.

#### 响应欺骗

1. 如果知道该主机(system权限)将要请求的NBNS主机名,则可以用低权限用户使用假的NBNS响应来回复该主机,因为是UDP协议且是发往本地的数据包所以可以很快的响应该主机的请求.
2. 而NBNS请求数据包中有一个TXID字段,该字段必须要在响应的数据包中匹配,而我们无法得到请求数据包,但可以直接穷举65536个可能值来强行匹配.
3. 如果网络中有请求的主机名的话,则不会进行NBNS请求,但我们可以绑定本地每一个UDP端口,这样将无法发出DNS请求,因为没有DNS请求的源端口,则会使用NBNS协议进行查询.

### Fake WPAD Proxy

#### WPAD

IE浏览器默认会访问`http://wpad/wpad.dat`来加载网络代理配置.同时默认windows服务(windows更新)也会请求该url来加载代理配置.但一般该url一般不会存在于网络中,`wpad`也不会存在于DNS服务器中,所以就会进行NBNS请求了.

#### Fake proxy

现在得到了本地主机(system权限)要请求的主机名为`wpad`(或者WPAD.DOMAIN.TLD),那么我们在本地(低权限用户)用大量主机名为`wpad`的NBNS响应数据包来响应本地主机的NBNS请求数据包,并在响应包中告诉本地主机(system)`wpad`的ip地址为127.0.0.1.

同时在本地(127.0.0.1)开启一个http服务,并在接受到`http://wpad/wpad.dat`请求时响应主机代理地址为本地127.0.0.1:80,这样目标的所有流量都会被重定向到我们开启的127.0.0.1:80.

### HTTP -> SMB NTLM Relay

现在所有的HTTP流量通过本地的http服务时会再被重定向到`http://127.0.0.1/xxxxx`(xxxxx是一个唯一标识符),在请求该url时将会被要求进行401 NTLM认证.

如果该请求来自于高权限用于如(system)那么我们将可以得到system用户的NTLM认证信息,再将其Relay到本地的SMB协议,即HTTP到SMB的中继,那么则可以获取该system用户的权限.

### Summary

主要是利用window的一些系统服务进行请求时会使用IE的代理配置来请求获取代理配置(且我们已经请求的URL为`http://wpad/wpad.dat`),我们使用NBNS欺骗使目标使用我们开启的http服务作为代理,那么所有流量都会经过我们的http服务,并将其重定向到强制到要求401认证的页面来获取其NTLM协议认证信息,最后将其中继回本地的SMB协议中来获取请求目标的权限.

### 参考

https://foxglovesecurity.com/2016/01/16/hot-potato/

## 0X02 Rotten Potato

在Hot Potato中使用的是windows服务并结合NBNS欺骗来获取高权用户的NTLM认证信息,但在不同的系统上,windows服务表现并不一致,并且windows服务的触发也不稳定,导致该利用并不通用和稳定.

而在Rotten Potato中则是使用DCOM/RPC来获取NTLM认证信息,因为DCOM/RPC在windows的各个版本上表现都一致,且可立即触发,通用性和稳定性都更好.

### DCOM请求

从IStorage中根据Guid获取一个BITSv1的实例,并且从127.0.0.1:6666端口来加载该实例.

```C
public static void BootstrapComMarshal()
{
IStorage stg = ComUtils.CreateStorage();
 
// Use a known local system service COM server, in this cast BITSv1
Guid clsid = new Guid("4991d34b-80a1-4291-83b6-3328366b9097");
 
TestClass c = new TestClass(stg, String.Format("{0}[{1}]", "127.0.0.1", 6666)); // ip and port
 
MULTI_QI[] qis = new MULTI_QI[1];
 
qis[0].pIID = ComUtils.IID_IUnknownPtr;
qis[0].pItf = null;
qis[0].hr = 0;
 
CoGetInstanceFromIStorage(null, ref clsid, null, CLSCTX.CLSCTX_LOCAL_SERVER, c, 1,       qis);
}
```

### DCOM <-> Middle <-> RPC







## 0x03 Juicy Potato

