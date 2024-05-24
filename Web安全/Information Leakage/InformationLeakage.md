# InformationLeakage
## .git
.git目录：使用git init初始化git仓库的时候，生成的隐藏目录，git会将所有的文件，目录，提交等转化为git对象，压缩存储在这个文件夹当中。每个git对象都有个哈希值代表这个对象，也就是上面所说的键值对的对应形式。

在使用git时，各种行为和文件都被标记了一个hash，然后记录在了.git/objects文件夹下，这个hash前两位是所在文件夹，后38位是文件夹内的文件。

利用.git泄露来恢复源码的本质就是去这个.git/objects文件夹下下载源码对应的hash文件并恢复成源文件。  

可以通过git泄露的固定地址找到commit类型的hash，再由commit类型文件找到tree类型的文件hash，之后通过tree类型找到blob类型hash，最后从blob类型文件中恢复源码。