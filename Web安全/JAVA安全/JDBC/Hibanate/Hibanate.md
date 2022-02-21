# Hibernate安全
官方文档:https://hibernate.net.cn/column/1.html

## Native SQL查询
原生SQL查询,通过执行Session.createSQLQuery()获取这个接口.
原生查询支持位置参数和命名参数：
```JAVA
Query query = sess.createSQLQuery("SELECT * FROM CATS WHERE NAME like ?").addEntity(Cat.class);
List pusList = query.setString(0, "Pus%").list();
     
query = sess.createSQLQuery("SELECT * FROM CATS WHERE NAME like :name").addEntity(Cat.class);
List pusList = query.setString("name", "Pus%").list(); 
```