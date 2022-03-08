# 漏洞修复
## XSS
## CSRF
## XXE
配置相关FEATURE来禁用外部实体。
```java
"http://apache.org/xml/features/disallow-doctype-decl", true 
"http://apache.org/xml/features/nonvalidating/load-external-dtd", false
"http://xml.org/sax/features/external-general-entities", false
"http://xml.org/sax/features/external-parameter-entities", false
```
```java
XMLConstants.ACCESS_EXTERNAL_DTD, ""
XMLConstants.ACCESS_EXTERNAL_STYLESHEET, ""
```
### 
## SQL注入
## SPEL表达式注入
使用`SimpleEvaluationContext`代替`StandardEvaluationContext` .  
[官方文档](https://docs.spring.io/spring-framework/docs/5.0.6.RELEASE/javadoc-api/org/springframework/expression/spel/support/SimpleEvaluationContext.html)