- [Web漏洞修复方案](#web漏洞修复方案)
  - [XSS](#xss)
    - [过滤](#过滤)
    - [配置](#配置)
    - [Referer XSS](#referer-xss)
      - [Referrer-Policy](#referrer-policy)
  - [CSRF](#csrf)
  - [SSRF](#ssrf)
  - [CORS](#cors)
    - [JSONP](#jsonp)
  - [XXE](#xxe)
  - [SQL注入](#sql注入)
    - [使用预编译](#使用预编译)
    - [强制转换类型](#强制转换类型)
    - [黑(白)名单](#黑白名单)
  - [RCE](#rce)
    - [系统命令执行](#系统命令执行)
    - [动态调用](#动态调用)
    - [表达式注入](#表达式注入)
      - [SPEL](#spel)
      - [OGNL\&EL](#ognlel)
    - [SSTI](#ssti)
      - [Freemarker](#freemarker)
      - [Velocity](#velocity)
    - [脚本语言](#脚本语言)
      - [Groovy](#groovy)
    - [组件漏洞](#组件漏洞)
  - [JNDI 注入](#jndi-注入)
# Web漏洞修复方案
## XSS
### 过滤
在Web的Filter或者拦截器设置过滤或者使用工具包(xssProtect)
* 黑名单对常见标签和关键字过滤,正则匹配.
* 转义编码(HTML实体,URL编码)常见特殊字符,如`'`,`"`,`<>`.
### 配置
* 正确设置响应包的Content-Type.
* Cookie-only.
### Referer XSS
#### Referrer-Policy
```
Referrer-Policy: no-referrer——不显示Referrer的任何信息在请求头中。  
Referrer-Policy: no-referrer-when-downgrade——这是默认值。当从https网站跳转到http网站或者请求其资源时（安全降级HTTPS→HTTP）,不显示Referrer的信息,其他情况（安全同级HTTPS→HTTPS,或者HTTP→HTTP）则在Referrer中显示完整的源网站的URL信息。  
Referrer-Policy: origin——表示浏览器在Referrer字段中只显示源网站的源地址（即协议、域名、端口）,而不包括完整的路径。
Referrer-Policy: origin-when-cross-origin——当发请求给同源网站时,浏览器会在Referrer中显示完整的URL信息,发个非同源网站时,则只显示源地址（协议、域名、端口）  
Referrer-Policy: same-origin——表示浏览器只会显示Referrer信息给同源网站,并且是完整的URL信息。所谓同源网站,是协议、域名、端口都相同的网站。 
Referrer-Policy: strict-origin——该策略更为安全些,和origin策略相似,只是不允许Referrer信息显示在从https网站到http网站的请求中（安全降级）。  
Referrer-Policy: strict-origin-when-cross-origin——和origin-when-cross-origin相似,只是不允许Referrer信息显示在从https网站到http网站的请求中（安全降级）。  
Referrer-Policy: unsafe-url——浏览器总是会将完整的URL信息显示在Referrer字段中,无论请求发给任何网站。 
```  
## CSRF
* 在敏感页面的请求中加入唯一的Token,后端对敏感请求进行Token校验,而正常情况下(如果存在XSS可获得Token)攻击者无法获得该Token,则无法冒充用户进行敏感操作.

* 验证 HTTP Referer 字段,该方法可轻易绕过.

* 在 HTTP 头中自定义属性并验证,类似于加Token,只是加在HTTP头中,即每个页面都会带上该Token校验头,而不必在每个相关页面代码中都加上Token.
## SSRF
* 统一错误信息,避免用户可以根据错误信息来判断远端服务器的端口状态。
* 限制请求的端口为http常用的端口,比如,80,443,8080,8090等。
* 过滤返回信息,验证远程服务器对请求的响应是比较容易的方法。如果web应用是去获取某一种类型的文件。那么在把返回结果展示给用户之前先验证返回的信息是否符合标准
* 禁用不需要的协议,仅仅允许http和https请求。可以防止类似于file:///,gopher://,ftp:// 等引起的问题。
* 根据业务需求,判定所需的域名是否是常用的几个,若是,将这几个特定的域名加入到白名单,拒绝白名单域名之外的请求。
* 根据请求来源,判定请求地址是否是固定请求来源,若是,将这几个特定的域名/IP加入到白名单,拒绝白名单域名/IP之外的请求。
* 黑名单内网ip。避免应用被用来获取获取内网数据,攻击内网。
* 若业务需求和请求来源并非固定,写一个过滤函数手动过滤。
## CORS
1. 正确配置跨域请求
如果 Web 资源包含敏感信息,则应在标头中正确指定源。Access-Control-Allow-Origin  
2. 仅允许受信任的站点  
标头中指定的来源应仅是受信任的站点。特别是,动态反映来自跨源请求的源而无需验证是很容易被利用的,应避免使用。Access-Control-Allow-Origin  
3. 避免将空列入白名单
避免使用标头。Access-Control-Allow-Origin: null
4. 避免在内部网络中使用通配符  
当内部浏览器可以访问不受信任的外部域时,仅信任网络配置来保护内部资源是不够的。
1. CORS 不能替代服务器端安全策略  
CORS定义了浏览器行为,绝不能替代服务器端对敏感数据的保护 - 攻击者可以直接伪造来自任何受信任来源的请求。因此,除了正确配置的 CORS 之外,Web 服务器还应继续对敏感数据（如身份验证和会话管理）应用保护。
### JSONP
1. 接受请求时检查referer来源；
2. 在请求中添加token并在后端进行验证；
3. 严格过滤 callback 函数名及 JSON 里数据的输出。
4. 设置SameSite: (https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Set-Cookie/SameSite)  

SameSite 是 HTTP 响应头 Set-Cookie 的属性之一。它允许您声明该 Cookie 是否仅限于第一方或者同一站点上下文。
SameSite 接受下面三个值:
* Lax:Cookies 允许与顶级导航一起发送,并将与第三方网站发起的 GET 请求一起发送。这是浏览器中的默认值。

* Strict:Cookies 只会在第一方上下文中发送,不会与第三方网站发起的请求一起发送。

* None:Cookie 将在所有上下文中发送,即允许跨站发送。
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
### 使用预编译
使用预编译功能够在提供运行效率的同时预防大多数得SQL注入, 预编译可以提前将要执行的SQL语句逻辑进行编译,用占位符对参数进行占位,在使用预编译时数据库只会将占位符上传入的参数当作数据进行计算而不是会改变原有的SQL逻辑和结构.  
JAVA:
```java
String sql = "select * from userinfo where id = ？";
ps = conn.prepareStatement(sql);
ps.setInt(1,id);
rs = ps.executeQuery();
```
php:
```php
$sql = "select * from userinfo where id = ？";
$stmt = $pdo->prepare($sql);
$stmt->bindValue(1,"test");
$result = $stmt->execute();
```
但是order by后面的语句无法进行预编译,只能进行拼接,需要进行手动过滤,同时需要正确使用占位符,就算使用了`prepareStatement`但还是进行的拼接执行则还是存在SQL注入.  
### 强制转换类型
对参数的类型进行严格转换,如int类型的参数则只接受int类型的数据.
### 黑(白)名单
对常见的SQL注入关键字建立黑名单,手动转义单双引号等或者在不影响业务的情况下会传入的数据进行白名单校验.
## RCE
### 系统命令执行
1. 不可信数据仅作为执行命令的参数而非命令。
不要将参数命令和执行命令做拼接,如图所示,将命令和参数分开传入。
![](1.png)
1. 对外部传入数据进行过滤。可通过白名单限制字符类型,仅允许字符、数字、下划线；或过滤转义以下符号:|;&$><`（反引号）!
### 动态调用
有时候代码中没有直接执行命令的函数,但支持从输入中的参数中获取数据动态执行.比如php中的call_user_function,JAVA的反射等,需要对用户输入的数据作过滤
### 表达式注入
#### SPEL
使用`SimpleEvaluationContext`代替`StandardEvaluationContext` .  
[官方文档](https://docs.spring.io/spring-framework/docs/5.0.6.RELEASE/javadoc-api/org/springframework/expression/spel/support/SimpleEvaluationContext.html)  
参考:  
[CVE-2022-22980](https://github.com/spring-projects/spring-data-mongodb/commit/5e241c6ea55939c9587fad5058a07d7b3f0ccbd3)
#### OGNL&EL
使用黑名单,对用户输入做严格过滤在不影响业务的情况下使用白名单.
### SSTI
#### Freemarker
`${...}`FreeMarker将会输出真实的值来替换大括号内的表达式,这样的表达式被称为interpolation(插值).
ftl文件中FTL标签的内容会被解析,在用户输入传递到这些标签中时需要做严格过滤.  

利用方式有两种方式  
API: 通过它可以访问底层Java Api Freemarker的 BeanWrappers.该配置在2.3.22版本之后默认不开启,但通过Configurable.setAPIBuiltinEnabled可以开启它。

new: new()函数可创建任意实现了TemplateModel接口的Java对象，同时还可以触发没有实现 TemplateModel接口的类的静态初始化块。官方提供的一种限制方式,使用Configuration.setNewBuiltinClassResolver(TemplateClassResolver)或设置 new_builtin_class_resolver 来限制这个内建函数对类的访问。  
此处官方提供了三个预定义的解析器(从2.3.17版开始)。  
* UNRESTRICTED_RESOLVER: 简单地调用 ClassUtil.forName(String).
* SAFER_RESOLVER: 和第一个类似,但禁止解析 ObjectConstructor,Execute和freemarker.template.utility.JythonRuntime.
* ALLOWS_NOTHING_RESOLVER: 禁止解析任何类。
#### Velocity
`#`关键字Velocity关键字都是使用#开头的，如#set、#if、#else、#end、#foreach等,Velocity变量都是使用`$`开头的，如:$name、$msg  
通过`$`获取变量的的class再加载恶意类.  
对传入到vm模板文件`$`中的内容做严格过滤.  
### 脚本语言
#### Groovy
对传入的Groovy代码需要做有效过滤.  
关键类函数特征
```java
groovy.lang.GroovyShell	evaluate
groovy.util.GroovyScriptEngine	run
groovy.lang.GroovyClassLoader	parseClass
javax.script.ScriptEngine	eval
```
### 组件漏洞
更新安全版本,根据POC对传入的数据做过滤.
## JNDI 注入
* 使用高版本JDK.
* 限制出网环境,限制服务器能访问公网(白名单)的范围.
* 过滤本地恶意Facotry.对高版本的JNDI利用可以在内部对Factory类做黑名单过滤.
* 限制协议,只允许java协议
```java
        URI uri = new URI("ldap://127.0.0.1:1389");
        String scheme = uri.getScheme();
        assertTrue(scheme == null || scheme.equals("java"), "Unsupported JNDI URI: ");
        System.out.println(scheme);
```  
参考:    
[CVE-2022-25167](https://github.com/apache/flume/commit/dafb26c)  
