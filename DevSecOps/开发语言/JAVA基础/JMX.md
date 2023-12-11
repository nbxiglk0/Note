- [JMX (Java Management Extensions)](#jmx-java-management-extensions)
  - [JMX 结构](#jmx-结构)
    - [Instrumentation](#instrumentation)
    - [JMX agent](#jmx-agent)
    - [Remote management](#remote-management)
  - [参考](#参考)

# JMX (Java Management Extensions)
JMX 技术提供了一种简单、标准的方法来管理应用程序、设备和服务等资源。  
JMX 规范定义了 Java 编程语言的体系结构、设计模式、API 和服务，用于管理和监视应用程序和网络。  
使用 JMX 技术，给定的资源由一个或多个称为托管 Bean 或 MBean 的 Java 对象来管理。  
JMX 技术定义了标准连接器（称为 JMX 连接器），使您能够从远程管理应用程序访问 JMX 代理。 使用不同协议的 JMX 连接器提供相同的管理接口。
## JMX 结构
JMX技术可以分为三个层次:  
* Instrumentation
* JMX agent
* Remote management
### Instrumentation
第一步首先是通过MBeans Java对象来实现对目标资源情况的访问，如果MBean检测到对应的资源之后就可以通过JMX代理对其进行管理，其中MBean不需要了解其使用的JMX代理。JMX规范的检测级别还提供了通知机制。此机制使MBean能够生成通知事件并将其传播到其他级别的组件。
### JMX agent
JMX代理直接控制资源并使它们可供远程管理应用程序使用。JMX代理通常与其控制的资源位于同一台机器上，但不是必须。JMX代理的核心组件是MBean Server，其中注册了MBean对象，JMX 代理同时还包括一组管理MBean对象的服务和一个可以让管理程序访问的通信适配器(连接器)。
### Remote management
JMX Instrumentation可以通过多种不同的方式进行访问，可以通过简单网络管理协议 (SNMP) 等现有管理协议，也可以通过专有协议。每个适配器通过特定协议提供在 MBean 服务器中注册的所有 MBean 的视图。连接器提供管理器端接口，用于处理管理器和 JMX 代理之间的通信。 每个连接器通过不同的协议提供相同的远程管理界面。 当远程管理应用程序使用此接口时，它可以通过网络透明地连接到 JMX 代理，而不管协议如何。 JMX 技术提供了一种标准解决方案，用于将 JMX 技术工具导出到基于 Java 远程方法调用 (Java RMI) 的远程应用程序。 
## 参考
https://docs.oracle.com/javase/tutorial/jmx/overview/index.html