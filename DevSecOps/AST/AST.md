- [DAST(Dynamic App security Testing)](#dastdynamic-app-security-testing)
- [SAST(Static App security Testing)](#saststatic-app-security-testing)
  - [Codeql](#codeql)
- [IAST(Interactive Application Security Testing 交互式应用程序安全测试)](#iastinteractive-application-security-testing-交互式应用程序安全测试)
  - [DAST+SAST=IAST](#dastsastiast)
  - [检测模式](#检测模式)
    - [插桩检测模式](#插桩检测模式)
      - [动态污点追踪](#动态污点追踪)
      - [交互式缺陷定位](#交互式缺陷定位)
      - [实现思路](#实现思路)
  - [参考](#参考)

# DAST(Dynamic App security Testing)
DAST(Dynamic App security Testing)动态应用安全测试，其实就是web扫描器，AVWS，Burp这种，依赖爬虫对URL进行收集，然后对URL发送Payload，根据响应来测试漏洞是否存在。
# SAST(Static App security Testing)
* SAST(Static App security Testing)静态应用安全测试，就是白盒扫描，代码审计，最原始的SAST就是对代码中触发常见漏洞的函数进行扫描(正则表达式，关键字匹配)，然后人工审计，后面发展了污点追踪，传入的可控参数被称为source，触发漏洞的函数叫做sink，如CodeQL，可以自动追踪参数的传递过程，即source到sink的调用堆栈，然后人工对该参数的传递过程进行确认是否存在漏洞。 
 
SAST原理：  
1. 首先通过调用语言的编译器或者解释器把前端的语言代码（如JAVA，C/C++源代码）转换成一种中间代码，将其源代码之间的调用关系、执行环境、上下文等分析清楚。
2. 语义分析：分析程序中不安全的函数，方法的使用的安全问题。
3. 数据流分析：跟踪，记录并分析程序中的数据传递过程所产生的安全问题。
4. 控制流分析：分析程序特定时间，状态下执行操作指令的安全问题。
5. 配置分析：分析项目配置文件中的敏感信息和配置缺失的安全问题。
6. 结构分析：分析程序上下文环境，结构中的安全问题。
7. 结合2）-6）的结果，匹配所有规则库中的漏洞特征，一旦发现漏洞就抓取出来。
8.  最后形成包含详细漏洞信息的漏洞检测报告，包括漏洞的具体代码行数以及漏洞修复的建议。
## Codeql
# IAST(Interactive Application Security Testing 交互式应用程序安全测试)
## DAST+SAST=IAST
而IAST其实就是两者相结合，在使用客户端发送请求的同时，后端对该请求的调用过程进行动态Hook，获取到参数真实环境的传递调用过程，而不会扫描全代码，减少测试时间，减少SAST误报高等缺点。
## 检测模式
IAST主要有几种检测模式
1. 插桩检测模式
* 动态污点追踪（被动式插桩）
* 交互式缺陷定位（主动式插桩）

1. 流量检测模式
* 终端流量代理
* 主机流量嗅探
* 旁路流量镜像

类似burp做代理当扫描器。
1. 日志检测模式
* Web 日志分析
* 爬虫检测模式
检测原理和流量监测模式差不多，只是
而主流的則是插桩检测模式。
### 插桩检测模式
插桩检测类似于Hook语言的函数方法，在函数方法中加入拦截代码，这样在该函数被调用的时候就能得到调用时的实时参数和各种信息进行进一步分析。
针对不同语言的插桩的实现方式不同  
JAVA: JAVA自带了插桩的接口，通过instrumentation接口，可以以一种标准的方式，在启动应用时添加javaagent参数来加载插桩探针，从而实现动态数据流污点追踪。
PHP: PHP主要是通过替换内部函数，对原始函数进行包装，其中插入探针进行分析，分析记录后再放行请求。
#### 动态污点追踪
污点追踪就是将程序执行的数据流分成三个部分，source，sanitizers，sinks进行分析。
* source: 即程序从外部获得的输入源，用户输入，文件读取的数据等。
* sanitizers: sanitizers即一些安全处理逻辑，对source进行的一系列过滤，校验等操作。
* sinks: sinks即使可能引发漏洞的函数方法，执行SQL语句，读写文件等。  
  
污点追踪就是追踪source在程序的执行流程中是否经过了正确的sanitizers处理后来到Sinks，如果没有的话那么该条数据流則是有漏洞的。  
其中动态污点追踪則是在应用运行时进行分析，获取的数据更加贴近真实环境，降低误报，这种IAST在检测过程中不会重新发送请求的称为被动式插桩。
#### 交互式缺陷定位
而动态污点追踪会追踪每一条数据流，在性能上对应用有一定影响，而交互式缺陷定位则通过只Hook关键方法，获取到请求数据信息后发送给IAST服务器，然后由IAST对数据进行分析后重新构造一个新的请求发送到服务器来对漏洞进行检测，这种会主动发送请求的则被称为主动插桩。
#### 实现思路
实现思路，即对常见的source，Sink函数进行Hook，然后储存相关请求数据到一个全局属性中，关键的是追踪的过程中需要对常见的传播点函数进行Hook，也就是一些对会source进行处理赋值的函数，传播点Hook的越多，那么得到的数据传递过程自然就越详细，但通过如果传播点过于广泛那么得到的数据传递过程也会过于冗余，不利于人工排查。
## 参考
https://www.freebuf.com/articles/web/290863.html  
https://www.03sec.com/Ideas/qian-tan-bei-dong-shiiast-chan-pin-yu-ji-shu-shi-x-1.html#morphing