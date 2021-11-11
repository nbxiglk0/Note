# JAVA-AduitBasic
- [JAVA-AduitBasic](#java-aduitbasic)
  - [反序列化](#反序列化)
    - [前置知识](#前置知识)
  - [反射](#反射)
  - [JAVA代理机制](#java代理机制)
    - [代理类](#代理类)
    - [调用处理器(InvocationHanlder)](#调用处理器invocationhanlder)
    - [创建代理对象](#创建代理对象)
    - [EventHanlder](#eventhanlder)
## 反序列化
### 前置知识
* 反序列化的类必须要显示声明**Serializable**接口.
* 反序列化数据的特征:前四个字节为`0xaced(Magic Number)0005(Version).
## 反射

## JAVA代理机制
### 代理类
代理类可以在运行时创建全新的类,能够实现指定的接口,具有以下方法:
* 指定接口所需要的全部方法.
* Object类的全部方法(toString,equals).
### 调用处理器(InvocationHanlder)
调用处理器是实现了`InvocationHandlder`接口的类对象,该接口只有一个`invoke`方法,无论何时调用代理对象的方法,调用处理器的`invoke`方法都会被调用,并向其传递Method对象和原始的调用参数.
```java
Object invoke(Object proxy,Method method,Object[] args)
```
### 创建代理对象
使用`Proxy`类的`newProxyInstance`方法创建代理对象.
```java
public static Object newProxyInstance(ClassLoader loader,Class<?>[] interfaces,InvocationHandler h)
```   
Demo:
```java
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class Handler {

    public static void main(String[] args) {
        realperson test = new realperson();
        InvocationHandler handler = new personone(test);
        person person = (person)(Proxy.newProxyInstance(handler.getClass().getClassLoader(),test.getClass().getInterfaces(),handler));
        System.out.println(person.getClass().getName());
        person.execute();
        person.sucess();
        }
    }
    interface person{
        void execute();
        void sucess();
    }
    class realperson implements person{
        public void execute(){
            System.out.println("exec");
        }
        public void sucess()
        {
            System.out.println("sucess");
        }
    }
    class personone implements InvocationHandler {
        private Object person;
        personone(Object person){
            this.person = person;
        }
        public Object invoke(Object object, Method method, Object[] args) throws InvocationTargetException, IllegalAccessException {
                System.out.println("Invoke call");
                method.invoke(person,args);
                return null;
        }
}

```
结果
```
Invoke call 1
exec
Invoke call 2
Invoke call 1
sucess
Invoke call 2
```
### EventHanlder
EventHandler是一个内置的实现了InvocationHandler的动态代理类,EventHanlder能够监控接口中的方法被调用了之后执行EventHanlder中成员的变量和方法.