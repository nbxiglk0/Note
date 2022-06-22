- [漏洞修复方案](#漏洞修复方案)
  - [XSS](#xss)
    - [Referer XSS](#referer-xss)
      - [Referrer-Policy](#referrer-policy)
    - [UA XSS](#ua-xss)
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
### Referer XSS
#### Referrer-Policy
Referrer-Policy: no-referrer——不显示Referrer的任何信息在请求头中。  
Referrer-Policy: no-referrer-when-downgrade——这是默认值。当从https网站跳转到http网站或者请求其资源时（安全降级HTTPS→HTTP），不显示Referrer的信息，其他情况（安全同级HTTPS→HTTPS，或者HTTP→HTTP）则在Referrer中显示完整的源网站的URL信息。  
*Referrer-Policy: origin——表示浏览器在Referrer字段中只显示源网站的源地址（即协议、域名、端口），而不包括完整的路径。*  
*Referrer-Policy: origin-when-cross-origin——当发请求给同源网站时，浏览器会在Referrer中显示完整的URL信息，发个非同源网站时，则只显示源地址（协议、域名、端口）*  
*Referrer-Policy: same-origin——表示浏览器只会显示Referrer信息给同源网站，并且是完整的URL信息。所谓同源网站，是协议、域名、端口都相同的网站。*  
Referrer-Policy: strict-origin——该策略更为安全些，和origin策略相似，只是不允许Referrer信息显示在从https网站到http网站的请求中（安全降级）。  
Referrer-Policy: strict-origin-when-cross-origin——和origin-when-cross-origin相似，只是不允许Referrer信息显示在从https网站到http网站的请求中（安全降级）。  
Referrer-Policy: unsafe-url——浏览器总是会将完整的URL信息显示在Referrer字段中，无论请求发给任何网站。   
### UA XSS
好像无利用场景，无危害
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
