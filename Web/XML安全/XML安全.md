# XXE安全

## XML语言

XML 指可扩展标记语言(EXtensible Markup Language)

### 基本语法格式

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?><!--xml文件的声明-->
<bookstore>                                                 <!--根元素-->
<book category="COOKING">        <!--bookstore的子元素，category为属性-->
<title>Everyday Italian</title>           <!--book的子元素，lang为属性-->
<author>Giada De Laurentiis</author>                  <!--book的子元素-->
<year>2005</year>                                     <!--book的子元素-->
<price>30.00</price>                                  <!--book的子元素-->
</book>                                                 <!--book的结束-->
</bookstore>                                       <!--bookstore的结束-->
```

standalone值是yes的时候表示DTD仅用于验证文档结构,从而外部实体将被禁用,但它的默认值是no,而且有些parser会直接忽略这一项.

格式:

- 所有 XML 元素都须有关闭标签。
- XML 标签对大小写敏感。
- XML 必须正确地嵌套。
- XML 文档必须有根元素。
- XML 的属性值须加引号。
- 若多个字符都需要转义，则可以将这些内容存放到CDATA里面 `<![CDATA[ 内容 ]]>`

### DTD(document type definition)

## XML注入





## XML外部实体注入(XXE)

### 外带通道(OOB)



## 参考

https://xz.aliyun.com/t/6887

## 