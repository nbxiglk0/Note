- [Clickjacking 点击劫持](#clickjacking-点击劫持)
	- [原理](#原理)
	- [防御措施](#防御措施)
		- [浏览器端](#浏览器端)
		- [X-Frame-Options](#x-frame-options)
		- [CSP](#csp)

# Clickjacking 点击劫持
## 原理
靠iframe标签将网站的敏感页面以不可见的形式嵌入到恶意页面中,诱导用户在恶意页面进行点击,但其实点击的是敏感页面的操作.
示例代码:
```css
<head>
	<style>
		#target_website {
			position:relative;
			width:128px;
			height:128px;
			opacity:0.00001;
			z-index:2;
			}
		#decoy_website {
			position:absolute;
			width:300px;
			height:400px;
			z-index:1;
			}
	</style>
</head>
...
<body>
	<div id="decoy_website">
	...decoy web content here...
	</div>
	<iframe id="target_website" src="https://vulnerable-website.com">
	</iframe>
</body>
```
实际情况需要根据页面大小调整长宽和位置来将敏感页面的按钮和恶意页面重合.
## 防御措施
### 浏览器端
1. 检查并强制执行当前应用程序窗口是主窗口或顶层窗口.
2. 使所有框架可见.
3. 防止点击不可见的框架.
4. 拦截并标记对用户的潜在点击劫持攻击.
### X-Frame-Options 
通过指定X-Frame-Options头来限制当前页面是否可以被加载到其它域中.  
```
X-Frame-Options: deny //拒绝被加载
X-Frame-Options: sameorigin //只能被同源网站加载
X-Frame-Options: allow-from https://normal-website.com //只能被指定网站加载
```
### CSP
通过CSP来限制其它网站加载当前网站.  
Content-Security-Policy: frame-ancestors 'self';
Content-Security-Policy: frame-ancestors normal-website.com;
Content-Security-Policy: frame-ancestors deny;