# 漏洞修复
## XSS
## CSRF
## XXE
配置相关FEATURE来禁用外部实体。
```java
"http://apache.org/xml/features/disallow-doctype-decl", true //禁止DOCTYPE 声明
"http://apache.org/xml/features/nonvalidating/load-external-dtd", false //禁止导入外部dtd文件
"http://xml.org/sax/features/external-general-entities", false //禁止外部普通实体
"http://xml.org/sax/features/external-parameter-entities", false //禁止外部参数实体
```
```java
XMLConstants.ACCESS_EXTERNAL_DTD, ""
XMLConstants.ACCESS_EXTERNAL_STYLESHEET, ""
```
## SQL注入
## 命令注入
1. 不可信数据仅作为执行命令的参数而非命令。
不要将参数命令和执行命令做拼接,如图所示,将命令和参数分开传入。
![](1.png)

2. 对外部传入数据进行过滤。可通过白名单限制字符类型，仅允许字符、数字、下划线；或过滤转义以下符号：|;&$><`（反引号）!
## SPEL表达式注入
使用`SimpleEvaluationContext`代替`StandardEvaluationContext` .  
[官方文档](https://docs.spring.io/spring-framework/docs/5.0.6.RELEASE/javadoc-api/org/springframework/expression/spel/support/SimpleEvaluationContext.html)