<a name="g3ItM"></a>
# 无passwd修改密码
替换shadow生成MD5 UNIX密码
```bash
openssl passwd -1 -salt <salt> <password>  
Useradd -g groupid -s /bin/bash -u uid -p 'slat-passwd' username
```
<a name="KXgiy"></a>
# 计划任务
```bash
格式: * * * * * command 
A、cron配置文件路径
#vi /etc/crontab
B、重启cron的方法
#/etc/rc.d/init.d/crond restart
Usage: /etc/rc.d/init.d/crond {start|stop|status|reload|restart|condrestart}
```
# ssh登录记录
who /var/log/wtmp
