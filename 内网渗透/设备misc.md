# 中标麒麟 kylincloud
中标麒麟 kylincloud  
可能的默认账号密码： root/123123  
WEB界面访问端口:9090  
WEB界面默认带shell执行终端  
# 华为iBMC
HTTPS:443  
V3服务器的缺省用户为“root”，缺省密码为“Huawei12#$”  
V5服务器G530 V5和G560 V5的缺省用户为“Administrator”，缺省密码为“Admin@9000”  
登录后可通过面板“远程控制”启动的java客户端连接kvm虚拟机控制台，通过重启修改单用户模式添加用户进入虚拟机  
单用户模式可通过百度找到大量资料  
# 达梦数据库
达梦数据库 DMDBMS 基于Oracle魔改  
端口5236 12345  
SQL管理工具: /disql SYSDBA/SYSDBA:5236  
默认账号密码: SYSDBA/SYSDBA  
默认文件目录: /opt/dmdbms/data/  
会创建系统账号dmdba，密码：PiPYaTFhPUyBo (不清楚是否是随机生成）  
SQL语句示例，表明是Oracle魔改  
```
select distinct owner from all_tables;
select table_name from all_tables where owner='xxx';
select owner,table_name from all_tab_columns where column_name like '%PASS%';
```
