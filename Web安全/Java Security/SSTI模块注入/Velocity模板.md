# Velocity模板
官方文档:https://velocity.apache.org/engine/2.3/user-guide.html
## 基础语法
`#`:标记表示VTL语法开始.  
`set`:表示一个指令,后面跟一个(),里面为表达式.  
`$`:变量标识符.  
`$customer.Address`等于`$customer.getAddress()`.
### 属性
`$customer.Address`  
1. 可以指customer对象的Address属性值
2. 可以指customer对象的Address相关方法
### 方法
