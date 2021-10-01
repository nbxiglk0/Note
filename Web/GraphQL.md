# GraphQL 安全

## Json --> form-urlencoded --> CSRF

正常GraphQL查询的content-type为application/json
```
POST /graphql HTTP/1.1
Host: redacted
Connection: close
Content-Length: 100
accept: */*
User-Agent: ...
content-type: application/json
Referer: https://redacted/
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.9
Cookie: ...

{"operationName":null,"variables":{},"query":"{\n  user {\n    firstName\n    __typename\n  }\n}\n"}
```
而由于中间件转换的原因,有时候服务器也会接受同样以form-urlencoded格式的post请求
```
POST /graphql HTTP/1.1
Host: redacted
Connection: close
Content-Length: 72
accept: */*
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36
Content-Type: application/x-www-form-urlencoded
Referer: https://redacted
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.9
Cookie: ...

query=%7B%0A++user+%7B%0A++++firstName%0A++++__typename%0A++%7D%0A%7D%0A
```
有时候该post请求则可能会发生csrf漏洞,因为开发者认为GraphQL的格式为Json格式,无法从表单创建,没有办法从urlencoded请求到json,而对GraphQL请求放弃做CSRF的检查
