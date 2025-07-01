- [不同JDK版本中漏洞利用限制](#不同jdk版本中漏洞利用限制)
  - [高版本JNDI注入](#高版本jndi注入)
    - [jdk21](#jdk21)
  - [JDK17+ Moudle模块化机制](#jdk17-moudle模块化机制)
  - [反射类限制](#反射类限制)

# 不同JDK版本中漏洞利用限制
## 高版本JNDI注入
### jdk21
* 在LDAP反序列化时多了一个trustSerialData字段判断是否允许反序列化数据。
* 在RMI利用本地工厂类时限制了能使用的工厂类包名（RMI）范围。

可以通过RMI协议中的StreamRemoteCall#executeCall()使用JRMP服务进行反序列化。

https://b1ue.cn/archives/529.html  
https://paper.seebug.org/942/
https://srcincite.io/blog/2024/07/21/jndi-injection-rce-via-path-manipulation-in-memoryuserdatabasefactory.html
## JDK17+ Moudle模块化机制
https://mp.weixin.qq.com/s/B5RiEDcT9A2uXntj76asBw
https://whoopsunix.com/docs/java/named%20module/
## 反射类限制
https://mp.weixin.qq.com/s/es3vXQLCs2--7KzOM4z6FQ
https://mp.weixin.qq.com/s/B5RiEDcT9A2uXntj76asBw
https://pankas.top/2023/12/05/jdk17-%E5%8F%8D%E5%B0%84%E9%99%90%E5%88%B6%E7%BB%95%E8%BF%87/