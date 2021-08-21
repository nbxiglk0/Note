# install 
上传nmap后无法运行，需要静默安装nmap文件夹中的vc运行库，对于nmap 7.8  
```cmd
vcredist_x86.exe /install /quiet
```
## FreeBSD ports不能用（低版本，源被破坏等）安装nmap
# pkg_add -r nmap
如果报错，根据发行版本或报错的URL手动搜索软件包，如：
http://ftp-archive.freebsd.org/pub/FreeBSD-Archive/ports/amd64/packages-9.1-release/Latest/nmap.tbz
```
wget <pakget url>
pkg_add nmap.tbz 
pkg_add: could not find package lua-5.1.5_4 !
pkg_add: could not find package pcre-8.31_1 !
wget http://ftp-archive.freebsd.org/pub/FreeBSD-Archive/ports/amd64/packages-9.1-release/Latest/lua51.tbz
wget http://ftp-archive.freebsd.org/pub/FreeBSD-Archive/ports/amd64/packages-9.1-release/Latest/pcre.tbz
```
pkg_add 安装这两个包，再安装nmap
# Tricks
跨网段扫描时 -PS 指定常见端口扫描存活,再-Pn -p0-65535扫描全端口
