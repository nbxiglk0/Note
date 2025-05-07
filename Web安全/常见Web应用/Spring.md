## Spring系列配置错误
### Actuator未授权访问
错误配置:
```xml
management.endpoints.web.exposure.include=*
```
### 安全配置
* 禁用接口 management.endpoints.enabled-by-default=false
* 使用spring security加个认证