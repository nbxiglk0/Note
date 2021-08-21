<a name="NzP5y"></a>
# Ticks

---

1. <br />利用appcmd工具独出IIS站点运行密码，在.net用户可以成功，asp_pool无法成功

	列出应用池     appcmd.exe list apppool<br />	列出应用池信息 appcmd.exe list apppool "CA" /text:*<br />	列虚拟目录     appcmd.exe list vdir<br />	查看信息       appcmd.exe list vdir "目录名" /text:*

2. ​<br />

利用iis自带工具解密加密的ConnectionString<br />	C:\WINDOWS\Microsoft.Net\Framework\v2.0.50727\aspnet_regiis -pdf "connectionStrings" c:\webconfig所在目录<br />	缺点，会修改web.config为明文，这点值得注意

3. ​<br />

 	IIS应用程序池的配置在目录C:\inetpub\temp\apppools下，并且c:\inetpub\temp\目录对.net用户可写可执行
