# CSRF

## 原理

原理在于服务端无法确定请求是否是由正常用户发起的.

## 防御

1. 在敏感页面的请求中加入唯一的Token,后端对敏感请求进行Token校验,而正常情况下(如果存在XSS可获得Token)攻击者无法获得该Token,则无法冒充用户进行敏感操作.

2. 验证 HTTP Referer 字段,该方法可轻易绕过.

3. 在 HTTP 头中自定义属性并验证,类似于加Token,只是加在HTTP头中,即每个页面都会带上该Token校验头,而不必在每个相关页面代码中都加上Token.

### 参考

https://blog.csdn.net/zl834205311/article/details/81773511