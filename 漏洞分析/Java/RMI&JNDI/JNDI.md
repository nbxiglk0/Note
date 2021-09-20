# JNDI
## 动态转换
JNDI能根据传递的URL协议自动转换与设置了对应的工厂与PROVIDER_URL。如果设置了工厂与PROVIDER_URL,但lookup时参数能够被控制,也会优先根据lookup的url进行动态转换.
## 命名引用
引用由Reference类表示，并且由地址和有关被引用对象的类信息组成，每个地址都包含有关如何构造对象。  

**参考:**  
https://paper.seebug.org/1091/#weblogic-rmi  
https://blog.csdn.net/li_w_ch/article/details/110114397