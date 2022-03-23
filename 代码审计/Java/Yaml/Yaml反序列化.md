- [Yaml反序列化](#yaml反序列化)
  - [SnakeYaml基础](#snakeyaml基础)
# Yaml反序列化
一般来说使用的较多的是snakeyaml组件进行yaml解析.
## SnakeYaml基础
Yaml的主要方法
```java
String	dump(Object data)
将Java对象序列化为YAML字符串。
void	dump(Object data, Writer output)
将Java对象序列化为YAML流。
String	dumpAll(Iterator<? extends Object> data)
将一系列Java对象序列化为YAML字符串。
void	dumpAll(Iterator<? extends Object> data, Writer output)
将一系列Java对象序列化为YAML流。
String	dumpAs(Object data, Tag rootTag, DumperOptions.FlowStyle flowStyle)
将Java对象序列化为YAML字符串。
String	dumpAsMap(Object data)
将Java对象序列化为YAML字符串。
<T> T	load(InputStream io)
解析流中唯一的YAML文档，并生成相应的Java对象。
<T> T	load(Reader io)
解析流中唯一的YAML文档，并生成相应的Java对象。
<T> T	load(String yaml)
解析字符串中唯一的YAML文档，并生成相应的Java对象。
Iterable<Object>	loadAll(InputStream yaml)
解析流中的所有YAML文档，并生成相应的Java对象。
Iterable<Object>	loadAll(Reader yaml)
解析字符串中的所有YAML文档，并生成相应的Java对象。
Iterable<Object>	loadAll(String yaml)
解析字符串中的所有YAML文档，并生成相应的Java对象。
```

