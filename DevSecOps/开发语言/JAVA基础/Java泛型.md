- [泛型](#泛型)
  - [泛型类型](#泛型类型)
    - [泛型类](#泛型类)
    - [泛型接口](#泛型接口)
    - [泛型方法](#泛型方法)
  - [参考](#参考)

# 泛型
泛型的作用在于进行代码复用，当某个方法适用于多种数据类型时不需要对每种数据类型都编写对应的实现，使用泛型可以只实现一次即可。  
如实现加法操作，参数可以是int，float，double类型，正常情况下需要根据不同类型实现对应的方法。  
```java
private static int add(int a, int b) {
    System.out.println(a + "+" + b + "=" + (a + b));
    return a + b;
}

private static float add(float a, float b) {
    System.out.println(a + "+" + b + "=" + (a + b));
    return a + b;
}

private static double add(double a, double b) {
    System.out.println(a + "+" + b + "=" + (a + b));
    return a + b;
}
```  
使用泛型则只需要实现一次：  
```java
private static <T extends Number> double add(T a, T b) {
    System.out.println(a + "+" + b + "=" + (a.doubleValue() + b.doubleValue()));
    return a.doubleValue() + b.doubleValue();
}
```  
另一个作用是提供类型的约束，提供编译前的检查，泛型中的类型在使用时指定，不需要强制类型转换（类型安全，编译器会检查类型）。
```java
List<String> list = new ArrayList<String>();
```
编译器会检查放入类型是否为String。
## 泛型类型
### 泛型类
### 泛型接口
### 泛型方法
## 参考
https://pdai.tech/md/java/basic/java-basic-x-generic.html