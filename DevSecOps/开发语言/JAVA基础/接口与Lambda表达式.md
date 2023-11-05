- [Lambda表达式](#lambda表达式)
  - [函数式接口](#函数式接口)
    - [BiFunction](#bifunction)
- [接口](#接口)
  - [static](#static)

# Lambda表达式
Lambda一般指匿名函数，在java中一般通过函数式接口来生成匿名类和方法。
## 函数式接口
只有一个抽象方法的接口，就是函数式接口。可以通过Lambda表达式来创建函数式接口的对象。在接口上标注了一个@FuncationInterface注解，此注解就是Java 8新增的注解，用来标识一个函数式接口。
### BiFunction
# 接口
## static
接口中可以定义static方法和default方法,static方法不会被实现类和子类基础,只能通过直接引用接口调用。
```java
public interface sta {
	static void show(){
		System.out.println("interface static method");
	}
}
//
public class test implements sta {

	public static void main(String[] args) {
		test test =new test();
		sta.show();
	}
}
```