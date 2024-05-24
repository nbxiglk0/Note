- [URL解析差异](#url解析差异)
  - [WEB应用中的权限校验](#web应用中的权限校验)
    - [常见URL获取方式](#常见url获取方式)
    - [常见匹配模式](#常见匹配模式)
      - [Ant匹配](#ant匹配)
      - [Regex匹配](#regex匹配)
    - [Tomcat](#tomcat)
      - [对分号的处理](#对分号的处理)
      - [Normalization](#normalization)
      - [总结Bypass](#总结bypass)
    - [SpringMVC](#springmvc)
      - [HandlerExecutionChain中对分号的处理](#handlerexecutionchain中对分号的处理)
      - [HandlerAdapter](#handleradapter)
      - [handle](#handle)
      - [总结Bypass](#总结bypass-1)
    - [SpringSecurity](#springsecurity)
      - [Spring Security对URI的统一处理](#spring-security对uri的统一处理)
        - [Request Method](#request-method)
        - [URL检查](#url检查)
        - [URL规范化](#url规范化)
        - [不可打印字符](#不可打印字符)
        - [手动修改配置](#手动修改配置)
      - [认证URI检查](#认证uri检查)
    - [Shiro](#shiro)
      - [对URI的处理](#对uri的处理)
    - [总结Bypass](#总结bypass-2)
  - [参考](#参考)

# URL解析差异

## WEB应用中的权限校验
WEB框架中往往是基于请求的URL进行权限校验，而由于容器和框架之间，t或者鉴权逻辑和后端路由转发逻辑之间对同一个URL的解析不一致经常会导致权限绕过的漏洞出现。  
通过构造畸形的URL请求导致在鉴权步骤中将其识别为无需认证的请求但在后端路由转发时又能转发到需要认证的Servlet中去。
### 常见URL获取方式
* getRequestURI: 原始 HTTP 请求路径部分在请求参数(即?)之前部分的内容，这是最原始的数据，包含原样的用户输入；
* getContextPath: 对应 Context 的路径，即 Web 容器挂载对应应用的前缀；
* getServletPath: 调用对应 Servlet 时的(相对)路径，不包含 Context 部分，且删除了路径参数；
* getPathInfo: 额外设置的路径信息，默认为 null；
### 常见匹配模式
#### Ant匹配
Ant风格的路径匹配规则包括以下几种模式：  
* ?：匹配任意单个字符。
* *：匹配任意数量（包括零个）的字符。
* **：匹配任意数量（包括零个）的目录路径。  

最长匹配规则（has more characters）：  
即越精确的模式越会被优先匹配到。例如，URL请求/app/dir/file.jsp，现在存在两个路径匹配模式/**/.jsp和/app/dir/.jsp，那么会根据模式/app/dir/*.jsp来匹配。
#### Regex匹配
正则表达式匹配，需要注意的是`.`并不包括`\r,\n`，这可能会导致绕过。
### Tomcat
Tomcat对URI的规范化处理流程主要是在`org/apache/catalina/connector/CoyoteAdapter#service`中，其中会调用`postParseRequest`来接请求进行解析，其中对URI中一些特殊字符的处理往往会产生绕过。  
以Tocmat 8.5.64，请求servlet`/user/service`URI为例。
#### 对分号的处理
首先会被特殊处理的字符就是`;`，当tomcat在URI遇到`;`时，会将第一个`;`后面截止到下一个`;`或者`/`之前的内容当作路径参数进行处理，相关处理逻辑位于`org/apache/catalina/connector/CoyoteAdapter#parsePathParameters`方法中。  
![](img/15-54-21.png)  
当URI中含有分号时，tomcat会调用ByteChunk.findBytes查找从`;`开始到URI结尾之间的第一个`;`或者`/`,然后tomcat会将该段字符从URI中删除，这之间的字符串如果含有`=`的话，并将该段字符转为请求的pathParamters键值对。  
![](img/16-05-12.png)
例如请求`/user;a=b/service`，tomcat会提取出`;`和`/`之间的`a=b`,然后将`a=b`添加到该请求的`pathParamters`，并在原来的URI中删除`;a=b/`。 
![](img/16-06-13.png)
在最后还会再检测uri中是否含有`;`，如果还有`;`,那么会再次进入while循环，也就是会递归处理`;`字符直到请求中不包含`;`为止。  
#### Normalization
在经过`parsePathParameters`处理之后然后会正常进行一次URLDecode，然后就会进入Normalization规范化的处理流程。  
![](img/16-31-35.png)  
相关处理逻辑位于`org/apache/catalina/connector/CoyoteAdapter#normalize`中，主要是对```"\", "//", "/./","/../```这个几个字符的处理。  
![](img/16-33-20.png)   
依次是将`\`替换为`/`。  
![](img/16-40-16.png)  
该处理流程需要`ALLOW_BACKSLASH`为true才行，不然会直接返回false。
将`//`替换`/`。
![](img/16-41-29.png)   
如果以`/.`或者`/..`结尾的话那么自动在后面再添加一个`/`。
![](img/16-44-32.png)
递归将`/.`/和`/../`规范为真实路径。    
![](img/16-46-50.png)  
中当`/../`超出根目录了时会返回400。
![](img/16-57-58.png)
#### 总结Bypass
以访问URI：`/user/service`为例。
通过以上处理流程可以得到以下URL也同样可以访问到`/user/service`对应的servlet：
* /user;a=b/service
* /user/service;favicon.ico(一些检测后缀如果为静态文件则放行的绕过)
* /user\service(需要ALLOW_BACKSLASH=true)
* /user//service或者//user/service
* /user/./service
* /user/abc/../service
* /u%73%65r/s%65%72vice  

这在用nginx对请求URI的关键字进行匹配或者Filter中使用getRequestURI进行URI匹配的权限校验可以轻松绕过。因为`HttpServletRequest.getRequestURI()`取得得URI是原生得URI，并没有经过tomcat处理，如果使用该方法获取的URI进行权限校验那么就会产生绕过。  
正确的URI获取应该是使用`HttpServletRequest.getServletPath()`+`HttpServletRequest.getPathinfo()`,这种方法返回的URI是经过处理后的URI。
### SpringMVC
在Spring中路由分发主要是在`org/springframework/web/servlet/DispatcherServlet#doDispatch`中进行的，继承于FrameworkServlet，也是实现Java EE Servlet 接口的一个Servlet。  
![](img/15-51-46.png)  
请求在经过容器处理后便会统一来到`DispatcherServlet`中，由Spring开始真正的路由转发。以spring-boot为例，内嵌tomcat。
在`DispatcherServlet`主要通过以下流程来调用目标controller。
#### HandlerExecutionChain中对分号的处理
第一步是寻找handler，这个handler其实是一个HandlerExecutionChain，也就是处理执行链，其中会包含后续要执行的controller和一些Interpertor：
```java
mappedHandler = getHandler(processedRequest);
```
![](img/16-13-36.png)  
可以看到默认含有六个handler，然后依次从每个mapping中去寻找请求对应的HandlerExecutionChain，其中查找使用 @Controller 修饰的类中的 @RequestMapping方法的mapping对应`RequestMappingHandlerMapping`，其中mapping的getHandler都是执行的父类`org/springframework/web/servlet/handler/AbstractHandlerMapping.java#getHandler`方法。
主要的寻找逻辑在`getHandlerInternal`方法中，首先通过initLookupPath()方法来获取请求的路径，然后再调用`lookupHandlerMethod()`得到对应路径的Handler方法。  
![](img/16-30-26.png)  
其中`initLookupPath()`处理如下  
![](img/16-32-08.png)  
可以看到最后会有一个`removeSemicolonContent()`方法，也就是和tomcat一样，会删除`;`开头的路径参数。`/service;ax=/user`会转变为`/service/user`。
![](img/16-34-18.png)  
然后就会根据处理后的URI从所有Method Map中找到匹配的Method，其中会通过三种方式分别来寻找匹配的Method，最后如果寻找到有多个匹配的话会进行比较选出最匹配的一个。  
第一步就是直接从`mappingRegistry#getMappingsByDirectPath()`中寻找相同URI的放入匹配List中。  
![](img/16-38-40.png)  
然后第一步如果有匹配的话会调用`addMatchingMappings`继续匹配。在其中会使用`org/springframework/web/util/pattern/PathPattern.java#match来进行匹配`。 
![](img/17-34-51.png)  
其中使用的DefaultPathContainer会在请求Dispatch之前对URI进行一次URL解码并将解码后的值放在valueToMatch变量中。  
![](img/17-33-55.png)    
在后续PathPattern的match时使用的就是valueToMatch变量。  
![](img/17-38-31.png)  
然后将该Method Handler和注册的Interpertor组合成一个HandlerExecutionChain返回。
#### HandlerAdapter
第二步得到该请求对应的HandlerExecutionChain之后根据该HandlerExecutionChain寻找对应的HandlerAdapter处理适配器。
```java
HandlerAdapter ha = getHandlerAdapter(mappedHandler.getHandler());
```
这个地方根据传入HandlerExecutionChain返回的是`RequestMappingHandlerAdapter`。
![](img/16-58-25.png)  

#### handle
最后通过HandlerAdapter去调用路由真正对应的Method。
```java
mav = invokeHandlerMethod(request, response, handlerMethod);
```
#### 总结Bypass
根据常规通过@RequestMapping路由的情况，可以有以下两种方式进行变形
以/service/user为例
* /ser%76%69%63e/us%65r(URL编码)
* /service;ax=/user(添加路径参数)  

针对其它路由定义模式，由于其对应的Mapping不同，在对应handler会有一些差异，可能会有其它的绕过方式。
### SpringSecurity
SpringSecurity作为权限校验框架，当其对同一个URI的解析和后端MVC框架路由时对URI的解析不一致时就容易遭成绕过。   

Spring Security 6.13，测试配置如下： 
```java
@Configuration
@EnableWebSecurity
public class Auth  {
@Bean
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception{
    http.authorizeHttpRequests((requests) -> requests
            .requestMatchers("/service/user").hasRole("USER")
            .requestMatchers("/service/login").permitAll()
            .requestMatchers("/**").permitAll()
    )
            .formLogin((form) -> form
                    .loginPage("/service/login").permitAll()
            )
            .logout(LogoutConfigurer::permitAll);
    return http.build();
    }
```  
这个时候再使用路径参数的方式就会得到400 ERROR，但URL编码还是可以正确请求。
![](img/14-45-26.png)  
![](img/14-45-36.png)  
原因在于Spring Security默认情况下会URI中中的一些可能的恶意字符进行拦截。相关处理逻辑主要位于`org/springframework/security/web/FilterChainProxy.java#doFilterInternal`中。  
![](img/15-12-44.png)
第一步调用`this.firewall.getFirewalledRequest()`来对请求进行一个防护墙似的过滤，查看其实现主要是检查几个地方，请求方法是否允许，URL是否在黑名单，Host是否是不可信来源，URL是否规范化，是否含有不可打印字符。  
![](img/15-15-27.png)  
#### Spring Security对URI的统一处理
##### Request Method
默认允许的请求方式如下：  
![](img/15-25-05.png)  
##### URL检查
而在检查URL是否在黑名单的`rejectedBlocklistedUrls`方法中，可以看到其默认有一个黑名单。 
![](img/15-21-43.png)  
![](img/15-22-06.png)  
这里encodedUrlContains主要检查`contextPath`和`getRequestURI()`。 
![](img/15-36-06.png)  
而decodedUrlContains主要检查`request.getServletPath()`和`getPathInfo()`。  
![](img/15-37-27.png)  
包括如下字符被拦截  
encodedUrlBlocklist:` //, \x00, %2F2F, %00, %25, \n, \r, %5C, %3B, %0A, %2E, %2F, %0D, ;, \`  
decodedUrlBlocklist:` //, \x00, %2F2F, %00, %, \u2028, \u2029 \n, \r, %5C, %3B, %0A, %2F, %0D, ;, \`  
所以在这里将会抛出异常导致路径参数绕过的方式失效。 
##### URL规范化
然后后面还有一个是否规范化，会依次检查几种请求路径的结果。
![](img/15-27-39.png)  
![](img/15-27-50.png)  
主要就是检测是否含有路径穿越符，如`./，/../，/.`.  
![](img/15-28-39.png)  
##### 不可打印字符
最后检测是否含有不可打印字符，主要检查字符是否小于`\u0020`或大于`\u007e`。
![](img/15-29-43.png)  
##### 手动修改配置
而从StrictHttpFirewall的接口可以看到，其中有一些检查可以关闭：
![](img/15-44-03.png)  
除了规范化和不可打印字符不能关闭外，其它检查都是可以手动设置修改的。  
#### 认证URI检查
以上是第一步对URI的统一检查，如开头设置中的`requestMatchers("/service/user").hasRole("USER")`。即对`/service/user`是需要认证才能访问的，而对URI一个URI是否需要认证的判断主要是在`org/springframework/security/web/access/intercept/AuthorizationFilter.java#doFilter`中。
![](img/16-29-57.png)  
其中调用`this.authorizationManager.check(this::getAuthentication, request);`来判断请求是否需要验证。 主要实现逻辑位于`org/springframework/security/web/access/intercept/RequestMatcherDelegatingAuthorizationManager.java#check`中，可以看到主要逻辑是从已经注册的mapping中循环取出对应的匹配器(RequestMatcher)，然后对请求进行匹配，如果匹配成功则取得对应的认证manager，由mannager进行权限校验。  
![](img/16-32-49.png)  
而RequestMatcher是一个接口，其对应的实现有很多。
![](img/16-38-32.png)  
以/user/service为例，其使用的匹配器是MvcRequestMatcher。  
![](img/16-44-12.png)  
该匹配器直接使用Spring MVC 的HandlerMapping来对路径进行匹配，从而避免了因为权限框架和SpringMVC对URL解析不一致的绕过问题。  
![](img/16-45-28.png)  
而可以看到Spring Security还支持其它匹配器，如AntMatcher，regexMatchers等等，而由于这些matcher和SpringMVC的路由转发时对URL的解析差异也导致过一些漏洞的产生。在Spring Security5.8之后，antMatchers、mvcMatchers 和 regexMatchers 方法已被弃用，取而代之的是新的 requestMatchers 方法。  
![](img/16-51-18.png)  
requestMatchers方法会自动选择最适合的匹配方式，如果类路径中含有SpringMVC并且在使用中，那么就会自动选择MvcRequestMatcher进行匹配。
### Shiro
和SpringSecurity一样，也存在因为URI解析差异性导致的绕过。
#### 对URI的处理
在Shiro中对URI的处理从`org/apache/shiro/web/filter/mgt/PathMatchingFilterChainResolver.java#getChain`开始，这里会从请求中获取requestURI然后根据requestURI进行路径匹配，然后得到该请求的后续执行链。  
![](img/15-32-57.png)  
主要逻辑位于`org/apache/shiro/web/util/WebUtils.java#getPathWithinApplication`中，可以看到在调用normalize之前先调用了`removeSemicolon`和之前一样来处理路径参数，可以看到处理的URI正是`ServletPath + getPathInfo`。  
![](img/15-52-24.png)  
区别在于shiro中是只处理最后一个分号，直接将最后一个分号之后的字符直接删除。  
![](img/15-58-22.png)  
在normalize时会传入`ALLOW_BACKSLASH`,默认为false，如果为true，则会用`\\`替换`/`。    
![](img/16-02-13.png)  
然后就是对常见路径字符的处理。  
![](img/16-11-02.png)  
依次处理顺序如下：
1. `\` -> `/`(如果ALLOW_BACKSLASH为true)。
2. 如果路径为`/.` 直接回 `/`。
3. 如果开头没有`/`，自动加上`/`。
4. `//` -> `/`
5. `/./` -> `/`
6. `/../` 递归解析到真实目录。

完成后返回，然后还会调用removeTrailingSlash删除尾部的`/`得到一个`requestURINoTrailingSlash`。  
![](img/16-24-50.png)  
![](img/16-22-50.png)  
从注释可以看到这是对Spring的特殊处理，因为在Spring中/a/b和/a/b/是等价的，而shiro的pathPattern当匹配/a/b时是匹配不到/a/b/的，这就会导致在URI后面添加一个`/`就可以绕过，所以在这里单独删除了最后的结尾的`/`。
### 总结Bypass
以访问URI：`/user/service`为例。
通过以上处理流程可以得到以下URL也同样可以访问到`/user/service`对应的代码，不考虑容器和其它框架的情况下：
1. /user\service (如果ALLOW_BACKSLASH为true)
2. /user/./service
3. user/service(可能容器会拦截)
4. /user//service
5. /user/abc/../service


## 参考
https://www.blackhat.com/docs/us-17/thursday/us-17-Tsai-A-New-Era-Of-SSRF-Exploiting-URL-Parser-In-Trending-Programming-Languages.pdf  
https://evilpan.com/2023/07/29/url-gotchas/  
https://evilpan.com/2023/08/19/url-gotchas-spring/  
https://docs.spring.io/spring-security/site/docs/current/api/org/springframework/security/web/firewall/StrictHttpFirewall.html
