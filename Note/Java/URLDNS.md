# URLDNS
该链无实际利用价值,主要用于无回显的情况下确定是否有反序列化漏洞,该链最后会向指定URL发送一次dns请求,因此用来确定是否被成功反序列化
## 分析
在HashMap类中重写了readObject()方法,其中在最后调用了hash()方法
```java
private void readObject(java.io.ObjectInputStream s)
             //
            ...
            ...
            // Read the keys and values, and put the mappings in the HashMap
            for (int i = 0; i < mappings; i++) {
                @SuppressWarnings("unchecked")
                    K key = (K) s.readObject();
                @SuppressWarnings("unchecked")
                    V value = (V) s.readObject();
                putVal(hash(key), key, value, false, false);//调用hash()
            }
        }
    }
```
在hash()的返回值中调用了key的hashcode()方法,而key的类型为一个Object，通过HashMap类的put方法,将key设置为一个java.net.url对象,如`URL u = new URL(null, url, handler);`那么这里调用的就是URL的hashcode方法
```java   
    static final int hash(Object key) {
        int h;
        return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16);//调用key的hashCode()方法
    }
```
在URL类的hashcode方法中，如果hashcode等于-1的话就会调用handler.hashcode方法,而hanlder的值则是new URL时候传入的第三个参数,且该handler为URLStreamHandler类型,`    transient URLStreamHandler handler;`,即调用URLStreamHandler类的hashcode()
```java
    public synchronized int hashCode() {
        if (hashCode != -1)
            return hashCode;

        hashCode = handler.hashCode(this);
        return hashCode;
    }
```
在URLStreamHandler类的hashcode()方法中，调用了getHostAddress()来得到addr的值
```java
    protected int hashCode(URL u) {
        int h = 0;

        // Generate the protocol part.
        String protocol = u.getProtocol();
        if (protocol != null)
            h += protocol.hashCode();

        // Generate the host part.
        InetAddress addr = getHostAddress(u);//使用getHostAddress（）
        if (addr != null) {
            h += addr.hashCode();
        } else {
            String host = u.getHost();
            if (host != null)
                h += host.toLowerCase().hashCode();
        }
        //...
        //...
        //...
    }
```
在getHostAddress()中则调用了InetAddress.getByName(host),而该方法会向host进行一次dns请求
```java
   protected synchronized InetAddress getHostAddress(URL u) {
        if (u.hostAddress != null)
            return u.hostAddress;

        String host = u.getHost();//从传入的url中获取host值
        if (host == null || host.equals("")) {
            return null;
        } else {
            try {
                u.hostAddress = InetAddress.getByName(host);//发送DNS请求
            } catch (UnknownHostException ex) {
            //...
            //...
```
## 利用
所以最后只要反序列化HashMap类时hashcode的值为-1即可,利用反射机制即可获取该属性并调用set设值即可

![](/Code-Aduit/JAVA/pic/1.png)
**利用链**:
1. HashMap->readObject()
2. HashMap->hash()
3. URL->hashCode()
4. URLStreamHandler->hashCode()
5. URLStreamHandler->getHostAddress()
6. InetAddress->getByName()
