- [JAVA权限校验安全](#java权限校验安全)
  - [Interceptor 拦截器](#interceptor-拦截器)
    - [绕过案例](#绕过案例)
  - [Filter 过滤器](#filter-过滤器)
  - [第三方组件](#第三方组件)
    - [Shiro](#shiro)
    - [Spring Security](#spring-security)
  - [JWT](#jwt)
  - [自研](#自研)
  - [参考](#参考)

# JAVA权限校验安全
## Interceptor 拦截器
SpringMVC提供的组件，用户在请求处理的过程中（前后）扩展自定义功能逻辑，常用于在拦截器中实现权限校验，通过在用户请求进入Controller逻辑之前拦截该请求，对请求的内容进行身份验证。
实现拦截器主要需要实现三个方法
```java
@Component
//定义拦截器类，实现HandlerInterceptor接口
//注意当前类必须受Spring容器控制
public class ProjectInterceptor implements HandlerInterceptor {
    @Override
    //原始方法调用前执行的内容
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        //实现请求的身份校验
        return true;
        //为true则继续执行后续业务逻辑，为false则中止后续执行。
    }
​
    @Override
    //原始方法调用后执行的内容
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, ModelAndView modelAndView) throws Exception {
        System.out.println("postHandle...");
        //业务逻辑处理完成后执行的逻辑，视图还未渲染。
    }
​
    @Override
    //原始方法调用完成后执行的内容
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) throws Exception {
        System.out.println("afterCompletion...");
        //整个请求完成后，视图结果渲染完成后调用。
    }
}
```
默认情况下Interceptor 拦截器执行顺序为注册的顺序，也可以通过InterceptorRegistry#order(int) 方法指定执行顺序。
### 绕过案例

## Filter 过滤器
## 第三方组件
### Shiro
### Spring Security
## JWT
## 自研
## 参考
https://mp.weixin.qq.com/s/pK5sWtZV6919qR5cMr_x0g
https://mp.weixin.qq.com/s/hFVlFcGPaWUl7whwL6RRkQ
https://www.cnblogs.com/xiaoyh/p/16444681.html