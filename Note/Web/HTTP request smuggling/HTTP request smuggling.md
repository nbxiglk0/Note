# Http Request Smuggling

## 原理
### 处理差异
链式网络中,前一个服务器于后一个服务器的处理差异导致对相同的内容进行了不同的处理  
如 Front-end(反代服务器) -> Back-end(业务处理服务器)

![](pic/2020711.jpg)
### Keep-Alive&Pipeline

**Keep-Alive:**  
添加请求头Connection: Keep-Alive,服务端在处理本次请求后不关闭tcp链接,后续对相同服务器的请求重用本次TCP链接从而减少服务器开销，HTTP1.1默认开启  
**Pipeline:**
客户端不必等待服务器响应直接像流水线按顺序发送请求,服务器按请求顺序依此返回响应包
### TCP链接重用
因为nginx地址与后端服务器地址一般固定,反代nginx服务器与后端服务器为了处理效率一般会重用一个tcp链接
## 