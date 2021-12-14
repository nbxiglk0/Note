- [JNDI](#jndi)
  - [动态转换](#动态转换)
  - [命名引用](#命名引用)
  - [JNDI注入](#jndi注入)
    - [RMI](#rmi)
    - [LDAP](#ldap)
  - [JNDI回显](#jndi回显)
  - [打法总结](#打法总结)
  - [补丁](#补丁)
# JNDI
JNDI全称为Java Naming and Directory Interface,也就是Java命名和目录接口.Java中使用最多的基本就是rmi和ldap的目录服务系统.
## 动态转换
JNDI能根据传递的URL协议自动转换与设置了对应的工厂与PROVIDER_URL。如果设置了工厂与PROVIDER_URL,但lookup时参数能够被控制,也会优先根据lookup的url进行动态转换.
## 命名引用
引用由Reference类表示，并且由地址和有关被引用对象的类信息组成，每个地址都包含有关如何构造对象。  
## JNDI注入
JNDI注入简单来说就是在JNDI接口在初始化时，如：InitialContext.lookup(URI)，如果URI可控，那么客户端就可能会被攻击.

正常JRMP协议流程:
1. 服务端A启动一个RMI的注册中心,并将要被远程调用的方法暴露在注册中心,其中存储着该方法的stub信息和该方法所处的ip地址和端口.
2. 客户端B连接服务端A启动的RMI注册中心,根据名称查询到对应的JNDI,并将其下回本地
3. 然后RMI通过下回来的数据得到对应方法的IP和端口,通过JRMP协议发起RMI请求
4. 
### RMI

### LDAP
## JNDI回显
## 打法总结
1. 打Registry注册中心
通过使用Registry连接到注册中心，然后把gadget chain对象bind注册到注册中心，从而引起注册中心反序列化RCE

2. 打InitialContext.lookup执行者
通过使用JNDI的实现，也就是rmi或ldap的目录系统服务，在其中放置一个某名称关联的Reference，Reference关联http服务中的恶意class，在某程序InitialContext.lookup目录系统服务后，返回Reference给该程序，使其加载远程class，从而RCE

3. JRMP协议客户端打服务端
使用JRMP协议，直接发送gadget chain的序列化数据到服务端，从而引起服务端反序列化RCE

4. JRMP协议服务端打客户端
使用JRMP协议，当客户端连上后，直接返回gadget chain的序列化数据给客户端，从而引起客户端反序列化RCE
## 补丁
jdk8u121版本开始，Oracle通过默认设置系统变量com.sun.jndi.rmi.object.trustURLCodebase为false，将导致通过rmi的方式加载远程的字节码不会被信任.
绕过:
1. 使用LDAP替换
2. 使用tomcat-el利用链  

在在JDK 11.0.1、8u191、7u201、6u211之后 com.sun.jndi.ldap.object.trustURLCodebase属性的值默认为false,也限制了利用LDAP远程加载.
绕过:
1. 
2. 通过加载本地类
**参考:**  
https://paper.seebug.org/1091/#weblogic-rmi  
https://blog.csdn.net/li_w_ch/article/details/110114397  
https://xz.aliyun.com/t/7079
https://www.blackhat.com/docs/us-16/materials/us-16-Munoz-A-Journey-From-JNDI-LDAP-Manipulation-To-RCE.pdf