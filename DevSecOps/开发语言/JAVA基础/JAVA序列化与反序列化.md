- [反序列化](#反序列化)
  - [基于Bean属性](#基于bean属性)
  - [前置知识](#前置知识)
  - [JEP290](#jep290)

# 反序列化
在JAVA中反序列化主要有两种实现机制.
* Bean属性
* Field机制

## 基于Bean属性

## 前置知识
* 反序列化的类必须要显示声明**Serializable**接口.
* 反序列化数据的特征:前四个字节为`0xaced(Magic Number)0005(Version).
## JEP290