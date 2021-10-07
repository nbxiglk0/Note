- [XML安全](#xml安全)
  - [0x01 XML语言](#0x01-xml语言)
    - [基本语法格式](#基本语法格式)
    - [DTD(document type definition)](#dtddocument-type-definition)
        - [引入DTD](#引入dtd)
      - [实体引用](#实体引用)
  - [0x02 XML注入](#0x02-xml注入)
  - [0x03 XML外部实体注入(XXE)](#0x03-xml外部实体注入xxe)
    - [外带通道(OOB)](#外带通道oob)
    - [防御](#防御)
  - [参考](#参考)
# XML安全

## 0x01 XML语言

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

文档类型定义,一个控制XML格式规范的文件,可以嵌入XML文档内部

也可以独立放在一个dtd文件中.

##### 引入DTD

内部引用,直接插入XML文档内部

```xml
<!DOCTYPE 根元素名称 [元素声明]>
```

外部引用

<!DOCTYPE 根元素名称 [元素声明]>

<!DOCTYPE 根元素名称 [元素声明]>

<!DOCTYPE 根元素名称 [元素声明]>

#### 实体引用

XML元素以形如 `<tag>foo</tag>` 的标签开始和结束,如果元素内部出现如`<` 的特殊字符,解析就会失败,为了避免这种情况,XML用实体引用（entity reference）替换特殊字符.XML预定义五个实体引用,即用`< > & ' "` 替换 `< > & ' "` .


## 0x02 XML注入





## 0x03 XML外部实体注入(XXE)

### 外带通道(OOB)

### 防御

1. 使用开发语言自带的方法禁用外部实体.
2. 过滤用户提交的XML数据.
3. 不允许用户提交自己DTD文件.

参考OWASP:https://cheatsheetseries.owasp.org/cheatsheets/XML_External_Entity_Prevention_Cheat_Sheet.html

## 参考

https://xz.aliyun.com/t/6887

https://xz.aliyun.com/t/7105

https://xz.aliyun.com/t/9519

https://bbs.ichunqiu.com/thread-44650-1-7.html