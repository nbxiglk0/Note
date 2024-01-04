- [Java脚本语言](#java脚本语言)
  - [Groovy](#groovy)
    - [版本区别](#版本区别)
    - [和Java的区别](#和java的区别)
    - [常用groovy语法](#常用groovy语法)
    - [在Java中使用groovy](#在java中使用groovy)
      - [Eval](#eval)
      - [GroovyShell#evaluate](#groovyshellevaluate)
    - [GroovyClassLoader](#groovyclassloader)
    - [ScriptEngine](#scriptengine)
  - [参考资料](#参考资料)

# Java脚本语言
## Groovy
Apache提供的基于JVM的对面对象编程语言，这门动态语言拥有类似Python、Ruby和Smalltalk中的一些特性，可以作为Java平台的脚本语言使用，Groovy代码动态地编译成运行于Java虚拟机（JVM）上的Java字节码，并与其他Java代码和库进行互操作。
### 版本区别
在4.x之前，groupId为`org.apache.groovy`，4.x开始groupId为`org.codehaus.groovy`。
```xml
    <dependency>
        <groupId>org.apache.groovy</groupId>
        <artifactId>groovy</artifactId>
        <version>4.0.17</version>
    </dependency>
```
### 和Java的区别
文档: https://groovy-lang.org/differences.html
其中默认导入以下包，不需要Import。
```
java.io。
java.lang.*
java.math.BigDecimal
java.math.BigInteger
java.net。
java.util.*
groovy.lang.*
groovy.util.*
```
### 常用groovy语法
使用execute方法可以直接将字符串当作命令执行。
```
"calc.exe".execute()
"${"calc.exe".execute()}"
```

### 在Java中使用groovy
在java中提供了几种方式运行groovy代码。
#### Eval
语法：`Eval.me(exp)`
```java
import groovy.util.Eval;
...
    public static void main(String[] args) {
        Eval.me("Runtime.getRuntime().exec('calc.exe')");
    }
```  
因为默认导入java.lang包，所以直接可以调用相关类方法。  
其中Eval方法支持传入参数，方法分别为x(),xy(),xyz()。
```java
        System.out.println(((int) Eval.x(4, "2*x")));
        System.out.println(((int) Eval.xy(4, 5, "x*y")));
        System.out.println(((int) Eval.xyz(4, 5, 6, "x*y+z")));
```
最终其实都是调用的都是`me(final String symbol, final Object object, final String expression)`方法。  
![](img/11-31-52.png)  
其中无参的me方法只是相关参数为null。
#### GroovyShell#evaluate
从上面可以看到Eval最后其实调用的是GroovyShell对象的evaluate方法，所以我们也可以直接使用GroovyShell#evaluate方法。
```java
import groovy.lang.GroovyShell;
public class Code {
    public static void main(String[] args) {
        GroovyShell sh = new GroovyShell();
        sh.evaluate("Runtime.getRuntime().exec('calc.exe')");
    }
}
```
### GroovyClassLoader
GroovyClassLoader用于在加载groovy 类并调用。
```groovy
import groovy.lang.GroovyClassLoader

def gcl = new GroovyClassLoader()
def clazz = gcl.parseClass('class Foo { void doIt() { println "ok" } }')
assert clazz.name == 'Foo'
def o = clazz.newInstance()
o.doIt()        
```
在java中使用
```java
import groovy.lang.GroovyClassLoader;
import groovy.lang.GroovyObject;
public class Code {
    public static void main(String[] args) throws ClassNotFoundException, InstantiationException, IllegalAccessException {
        GroovyClassLoader classLoader = new GroovyClassLoader();
        Class c = classLoader.parseClass(
                "class Run {\n" +
                "    public static void main(String[] args) {\n" +
                "        GroovyShell shell = new GroovyShell();\n" +
                "        shell.evaluate(\"\\\"calc.exe\\\".execute()\");\n" +
                "    }\n" +
                "\n" +
                "}\n");
    GroovyObject object = (GroovyObject) c.newInstance();
    object.invokeMethod("main","");
    }
}

```
### ScriptEngine
```java
public class Code {
    public static void main(String[] args) throws ScriptException {
        ScriptEngine scriptEngine = new ScriptEngineManager().getEngineByName("groovy");
        scriptEngine.eval("\"calc.exe\".execute()");
    }
}

```
## 参考资料
https://groovy-lang.org/