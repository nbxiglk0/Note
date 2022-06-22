- [漏洞修复方案](#漏洞修复方案)
  - [XSS](#xss)
  - [CSRF](#csrf)
  - [XXE](#xxe)
  - [SQL注入](#sql注入)
  - [命令注入](#命令注入)
  - [SPEL表达式注入](#spel表达式注入)
    - [StandardEvaluationContext](#standardevaluationcontext)
      - [参考](#参考)
  - [JNDI 注入](#jndi-注入)
    - [限制协议](#限制协议)
      - [参考](#参考-1)
# 漏洞修复方案
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
### StandardEvaluationContext
使用`SimpleEvaluationContext`代替`StandardEvaluationContext` .  
[官方文档](https://docs.spring.io/spring-framework/docs/5.0.6.RELEASE/javadoc-api/org/springframework/expression/spel/support/SimpleEvaluationContext.html)  
#### 参考
[CVE-2022-22980](https://github.com/spring-projects/spring-data-mongodb/commit/5e241c6ea55939c9587fad5058a07d7b3f0ccbd3)
## JNDI 注入
### 限制协议
1. 只允许java协议
```java
        URI uri = new URI("ldap://127.0.0.1:1389");
        String scheme = uri.getScheme();
        assertTrue(scheme == null || scheme.equals("java"), "Unsupported JNDI URI: ");
        System.out.println(scheme);
```  
#### 参考  
[CVE-2022-25167](https://github.com/apache/flume/commit/dafb26c)  
