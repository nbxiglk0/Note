- [SercurityMannager](#sercuritymannager)
  - [启动SercurityMannager](#启动sercuritymannager)
  - [配置](#配置)

# SercurityMannager
JAVA安全管理器可以对运行的代码进行权限控制,即可以对读写,命令执行等类操作进行配置.  
默认配置文件路径: $JAVA_HOME/jre/lib/security/java.policy,即当未指定配置文件时将会使用该配置.  

## 启动SercurityMannager

1. 启动程序时通过附加参数启动(指定了配置文件)
   `-Djava.security.manage-Djava.security.policy="E:/java.policy"`
2. 编码方式启动

```java
System.setSecurityManager(new SecurityManager());
```

## 配置

示例配置文件:

```
// Standard extensions get all permissions by default

grant codeBase "file:${{java.ext.dirs}}/*" {
    permission java.security.AllPermission;
};

// default permissions granted to all domains

grant { 
    // Allows any thread to stop itself using the java.lang.Thread.stop()
    // method that takes no argument.
    // Note that this permission is granted by default only to remain
    // backwards compatible.
    // It is strongly recommended that you either remove this permission
    // from this policy file or further restrict it to code sources
    // that you specify, because Thread.stop() is potentially unsafe.
    // See the API specification of java.lang.Thread.stop() for more
        // information.
    permission java.lang.RuntimePermission "stopThread";

    // allows anyone to listen on un-privileged ports
    permission java.net.SocketPermission "localhost:1024-", "listen";

    // "standard" properies that can be read by anyone

    permission java.util.PropertyPermission "java.version", "read";
    permission java.util.PropertyPermission "java.vendor", "read";
    permission java.util.PropertyPermission "java.vendor.url", "read";
    permission java.util.PropertyPermission "java.class.version", "read";
    permission java.util.PropertyPermission "os.name", "read";
    permission java.util.PropertyPermission "os.version", "read";
    permission java.util.PropertyPermission "os.arch", "read";
    permission java.util.PropertyPermission "file.separator", "read";
    permission java.util.PropertyPermission "path.separator", "read";
    permission java.util.PropertyPermission "line.separator", "read";

    permission java.util.PropertyPermission "java.specification.version", "read";
    permission java.util.PropertyPermission "java.specification.vendor", "read";
    permission java.util.PropertyPermission "java.specification.name", "read";

    permission java.util.PropertyPermission "java.vm.specification.version", "read";
    permission java.util.PropertyPermission "java.vm.specification.vendor", "read";
    permission java.util.PropertyPermission "java.vm.specification.name", "read";
    permission java.util.PropertyPermission "java.vm.version", "read";
    permission java.util.PropertyPermission "java.vm.vendor", "read";
    permission java.util.PropertyPermission "java.vm.name", "read";
};
```

directory/ 表示directory目录下的所有.class文件，不包括.jar文件
directory/* 表示directory目录下的所有的.class及.jar文件
directory/- 表示directory目录下的所有的.class及.jar文件，包括子目录