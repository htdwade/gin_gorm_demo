# gin_gorm_demo
基于gin+gorm开发的todo清单项目

## 配置MySQL
* 在数据库中执行以下命令，创建本项目所用的数据库：
```sql
CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;
```
* 在`gin_gorm_demo/config/config.ini`文件中按如下提示配置数据库连接信息。
```ini
[mysql]
host=数据库ip
port=你的数据库端口
user=你的数据库用户名
password=你的数据库密码
dbname=bubble
```

## 编译运行
```bash
go build
./gin_gorm_demo.exe
```
