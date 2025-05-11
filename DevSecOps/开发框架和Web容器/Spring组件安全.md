- [Spring 组件安全](#spring-组件安全)
  - [Spring Boot Actuator 未授权访问](#spring-boot-actuator-未授权访问)
    - [配置](#配置)
    - [版本区别](#版本区别)
    - [漏洞利用](#漏洞利用)
      - [/env](#env)
      - [/refresh](#refresh)
      - [/trace(httptrace)](#tracehttptrace)
      - [/mappings](#mappings)
      - [/shutdown](#shutdown)
      - [/heapdump](#heapdump)
      - [修改配置](#修改配置)
  - [参考](#参考)

# Spring 组件安全
## Spring Boot Actuator 未授权访问
### 配置
不安全配置:
```yaml
management.security.enabled=false #所有端点可未授权访问
endpoints.env.sensitive=false #只针对env端点
```
安全配置:
1. 开启HTTP basic认证  

添加依赖
```xml
<dependency>
<groupId>org.springframework.boot</groupId>
<artifactId>spring-boot-starter-security</artifactId>
</dependency>
```

application.properties 添加用户名和密码
```yaml
security.user.name=admin
security.user.password=123456
management.security.enabled=true
management.security.role=ADMIN
```

Spring Boot Actuator 针对于所有 endpoint 都提供了两种状态的配置
* enabled 启用状态： 默认情况下除了 shutdown 之外，其他 endpoint 默认是启用状态。
* exposure 暴露状态： endpoint 的 enabled 设置为 true 后，还需要暴露一次，才能够被访问，默认情况下只有 health 和 info 是暴露的。
### 版本区别
* SpringBoot <= 1.5.x: 不需要任何配置的，直接就可以访问到端点。
* 1.5.x<=SpringBoot<=2.x: 默认只能访问到health和info端点。
* SpringBoot>=2.x: 默认只能访问到health和info端点的,访问多了个前缀/actutator。
### 漏洞利用
#### /env
Spring 的 ConfigurableEnvironment 公开属性，其中包括系统版本，环境变量信息、内网地址等信息，但是一些敏感信息会被关键词匹配，做隐藏*处理.  
* 可以通过post请求来新增系统的全局变量，或者修改全局变量的值.
#### /refresh 
* 用于配置修改后的刷新，常用于结合/env+其他依赖用来触发漏洞。
#### /trace(httptrace)
* 请求这个站点时的完整的http包，其中就可能包括用户的session.
#### /mappings
* @RequestMapping 路由信息
#### /shutdown 
关闭程序
#### /heapdump
* 堆栈信息，里面可能包括数据库连接凭证等敏感内容，可以用"Eclipse Memory Analyzer"内存分析工具。
#### 修改配置
## 参考
https://xz.aliyun.com/news/9218