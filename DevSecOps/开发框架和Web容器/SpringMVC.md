- [Spring MVC](#spring-mvc)
	- [SpringMVC Core Technologiess](#springmvc-core-technologiess)
		- [IOC Container](#ioc-container)
			- [ApplicationContext](#applicationcontext)
			- [XML配置](#xml配置)
			- [Groovy 定义](#groovy-定义)
			- [创建Bean的几种方式](#创建bean的几种方式)
				- [构造方法](#构造方法)
				- [静态工厂方法](#静态工厂方法)
				- [实例工厂方法](#实例工厂方法)
				- [Singleton 单例Bean](#singleton-单例bean)
				- [Prototype原型 多例Bean](#prototype原型-多例bean)
			- [嵌套类定义](#嵌套类定义)
			- [获取Bean](#获取bean)
		- [Dependency Injection](#dependency-injection)
			- [基于构造方法](#基于构造方法)
			- [基于Bean参数](#基于bean参数)
			- [基于setter方法](#基于setter方法)
			- [子类Bean](#子类bean)
	- [常用注解](#常用注解)
		- [@Autowired](#autowired)
	- [Code Demo](#code-demo)

# Spring MVC
主要是阅读官方文档时的一些摘要.  
文档版本: 6.011  
spring-webmvc包含:
```
spring-aop
spring-beans
spring-context
spring-core
spring-expression
spring-web
```
spring-web包括:
```
spring-core
spring-beans
```
## SpringMVC Core Technologiess
SpringMVC主要的核心技术:

* IOC 控制反转容器(IOC Container)
* 
### IOC Container
控制反转(Inversion of Control)(IOC)容器,IOC也称依赖注入(DI).  
依赖关系是指如A对象的实现常常会依赖于B对象,而B对象又依赖于C对象,这种相互形成的依赖关系.  
而依赖注入则是通过一个对象构造方法的参数,或者对应工厂方法的参数,或者这两种方法返回后设置的属性来定义该对象的依赖关系,也就是该对象需要依赖于哪些其它对象.  

而SpringMVC实现了一种容器来创建这些Bean对象,并根据依赖关系向该Bean对象注入该Bean对象需要的其它Bean对象.
Bean对象: 在 Spring 中，构成应用程序主干并由 Spring IoC 容器管理的对象称为 Bean. bean 是由 Spring IoC 容器实例化、组装和管理的对象。Bean 以及它们之间的依赖关系反映在容器使用的配置元数据中。

创建Bean的接口为`org.springframework.beans.factory.BeanFactory`,定义了对Bean的一系列操作,有几个比较关键的实现子接口.  
* ApplicationContext: 在BeanFactory的基础上补充了一些功能.
* WebApplicationContext: ApplicationContext的子接口,专门用于WEB应用的接口.  
#### ApplicationContext
org.springframework.context.ApplicationContext该接口代表 Spring IoC 容器并且负责实例化、配置和组装 Bean。容器通过 XML、Java 注释或 Java 代码来获取有关实例化、配置和装配哪些对象的元数据指令.  
而从XML中获取元数据指令的话,用的比较多的有两个实现类`ClassPathXmlApplicationContext`和`FileSystemXmlApplicationContext`,也就是从类路径或者系统路径加载xml格式的配置文件,这个xml配置文件就包含了相关Bean加载元数据命令.    
而整个IOC容器就会根据配置文件的元数据命令和对应的Bean对象在容器内进行配置组装.  
![](./img/2023-08-04-15-11-50.png)  
但IOC容器加载配置Bean的过程实际与XML配置数据无关,XML配置只是其中一种形式,也可以不使用XML文件的形式,也可以用JAVA的方式(基于注解),`@Configuration`, `@Bean`, `@Import`, `@DependsOn`.  
而Spring配置至少有一个或多个由容器管理的Bean进行定义.    
#### XML配置
示例XML形式的配置格式:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://www.springframework.org/schema/beans
		https://www.springframework.org/schema/beans/spring-beans.xsd">

	<bean id="petStore" class="org.springframework.samples.jpetstore.services.PetStoreServiceImpl">
		<property name="accountDao" ref="accountDao"/>
		<property name="itemDao" ref="itemDao"/>
	</bean>

	<bean id="accountDao"
		class="org.springframework.samples.jpetstore.dao.jpa.JpaAccountDao">

	</bean>

	<bean id="itemDao" class="org.springframework.samples.jpetstore.dao.jpa.JpaItemDao">
	</bean>
</beans>
```
* id属性定义该bean的名称,class属性指向该bean的定义类全路径名称,在标签内可以使用`property`初始化该bean的一些属性,其中ref属性可以指向其它bean id来讲其它bean加载到该bean中.  
* 可以不配置bean id,容器会自动生成一个名称,名称通常是将对应的类名首字母变为小写,如果没有手动配置id,那么无法使用在xml中使用ref来对该bean引用,因为自动生成名称是在xml加载后进行的.  
#### Groovy 定义
和XML形式类似,通常保存在.groovy文件中,
```Groovy
beans {
	dataSource(BasicDataSource) {
		driverClassName = "org.hsqldb.jdbcDriver"
		url = "jdbc:hsqldb:mem:grailsDB"
		username = "sa"
		password = ""
		settings = [mynew:"setting"]
	}
	sessionFactory(SessionFactory) {
		dataSource = dataSource
	}
	myService(MyService) {
		nestedBean = { AnotherBean bean ->
			dataSource = dataSource
		}
	}
}
```
#### 创建Bean的几种方式
##### 构造方法
```xml
<bean id="exampleBean" class="examples.ExampleBean"/>
<bean name="anotherExample" class="examples.ExampleBeanTwo"/>`
```
##### 静态工厂方法
* 通常class指向该bean的定义类,然后通过其构造方法来实例化该对象(类似于new),但也可以指向一个静态工厂方法来创建该实例对象.  
```java
<bean id="clientService"
	class="examples.ClientService"
	factory-method="createInstance"/>
```
```java
public class ClientService {
	private static ClientService clientService = new ClientService();
	private ClientService() {}

	public static ClientService createInstance() {
		return clientService;
	}
}
```

该bean的创建将会通过`examples.ClientService.createInstance`方法进行创建.
##### 实例工厂方法
通过实例工厂对象的方法创建Bean,不定义class属性而是使用factory-bean属性指向实例工厂的Bean id.
```xml
<bean id="serviceLocator" class="examples.DefaultServiceLocator">
</bean>
<bean id="clientService"
	factory-bean="serviceLocator"
	factory-method="createClientServiceInstance"/>
```
```java
public class DefaultServiceLocator {

	private static ClientService clientService = new ClientServiceImpl();

	public ClientService createClientServiceInstance() {
		return clientService;
	}
}
```
##### Singleton 单例Bean
Bean默认是单例模式,即在一个IOC容器内,只有一个Bean实例对象,后续所有从IOC容器中请求该Bean对象都是返回的同一个对象.  
![](./img/2023-08-04-17-45-09.png)
##### Prototype原型 多例Bean
指定`scope="prototype"`,对应的Prototyep模式在每次请求该Bean对象时,IOC容器都会创建一个新的Bean对象返回.  
![](./img/2023-08-04-17-47-25.png)
#### 嵌套类定义
如果需要指向一个类中的静态嵌套类,那么可以通过`.`或者`$`来分隔.如`com.example.Abc`类中有一个`Edf`类,则可以用`com.examlpe.Abc$Edf`或者`com.examlpe.Abc.Edf`


加载xml配置文件:  
```java
ApplicationContext context = new ClassPathXmlApplicationContext("services.xml");
```
#### 获取Bean
通过getBean就能获取到指定对象的实例.  
```java
ApplicationContext context = new ClassPathXmlApplicationContext("services.xml");

PetStoreService service = context.getBean("petStore", PetStoreService.class);

List<String> userList = service.getUsernameList();
```
### Dependency Injection
依赖注入,也就是在springFramework中,IOC容器会通过几种方式来识别一个A对象在创建需要哪些其它Bean对象,如果需要的Bean对象未创建则自动创建需要的Bean对象并将需要Bean对象传入需要A对象中然后完成实例化.  
解析依赖关系的流程:
* 创建 ApplicationContext 时，会使用描述所有 Bean 的配置元数据对其进行初始化。配置元数据可以通过 XML、Java 代码或注解指定。
* 对于每个 Bean 来说，其依赖关系以属性、构造函数参数或静态工厂方法参数）的形式表示。这些依赖关系会在实际创建 Bean 时提供给 Bean。
* 每个属性或构造函数参数都是要设置的值的实例对象，或者是对容器中另一个 Bean 的引用。
* 类型转换:每个属性或构造函数参数的值都会从其指定格式转换为该属性或构造函数参数的实际类型。
#### 基于构造方法
最常见的就是对象的构造方法的参数中需要一个其它对象,如SimpleMovieLister对象的创建需要一个MovieFinder对象,则IOC容器就会知道SimpleMovieLister对象依赖于MovieFinder对象.
```java
public class SimpleMovieLister {
	private final MovieFinder movieFinder;

	public SimpleMovieLister(MovieFinder movieFinder) {
		this.movieFinder = movieFinder;
	}
}
```
#### 基于Bean参数
通过在定义Bean时的constructor-arg属性来指定其需要依赖的对象.
```java
package x.y;

public class ThingOne {

	public ThingOne(ThingTwo thingTwo, ThingThree thingThree) {
	}
}
```
```xml
<beans>
	<bean id="beanOne" class="x.y.ThingOne">
		<constructor-arg ref="beanTwo"/>
		<constructor-arg ref="beanThree"/>
	</bean>
	<bean id="beanTwo" class="x.y.ThingTwo"/>
	<bean id="beanThree" class="x.y.ThingThree"/>
</beans>
```
#### 基于setter方法
在调用无参数构造函数或无参数静态工厂方法实例化 Bean 后，容器会调用 Bean 上的setter方法，从而实现基于setter的 DI。
#### 子类Bean
Bean在定义时也可以指定父类Bean Class,从而继承父类Class的属性或者覆盖.  
https://docs.spring.io/spring-framework/reference/core/beans/child-bean-definitions.html  
```java
<bean id="inheritedTestBean" abstract="true"
		class="org.springframework.beans.TestBean">
	<property name="name" value="parent"/>
	<property name="age" value="1"/>
</bean>

<bean id="inheritsWithDifferentClass"
		class="org.springframework.beans.DerivedTestBean"
		parent="inheritedTestBean" init-method="initialize">
	<property name="name" value="override"/>
	<!-- the age property value of 1 will be inherited from parent -->
</bean>
```
## 常用注解
### @Autowired
> JSR-330 standard 标准中可以用`@Inject`代替`@Autowired`.  
> 
可以用在方法和字段上面,ioc 容器会自动注入方法需要的Bean或者向该类注入该字段类型的Bean,如果用在数组或者set等类型的对象上,那么会自动注入所有该类型的Bean.
```java
	@Autowired
	private MovieCatalog[] movieCatalogs;
	//注入所有MovieCatalog类型的bean到movieCatalogs对象.
	
	@Autowired
	public void setMovieCatalogs(Set<MovieCatalog> movieCatalogs) {
		this.movieCatalogs = movieCatalogs;
	}
	//注入所有类型的MovieCatalog到movieCatalogs对象.
```

**原理**: 在启动spring IoC时，容器自动装载了一个AutowiredAnnotationBeanPostProcessor后置处理器，当容器扫描到@Autowied、@Resource或@Inject时，就会在IoC容器自动查找需要的bean，并装配给该对象的属性
## Code Demo
