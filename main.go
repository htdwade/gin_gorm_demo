package main

import (
	"fmt"
	"gin_gorm_demo/config"
	"gin_gorm_demo/dao"
	"gin_gorm_demo/models"
	"gin_gorm_demo/routers"
	"gopkg.in/ini.v1"
)

func main() {
	cfg := new(config.AppConfig)
	// 映射ini文件到结构体cfg
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("load config failed, err:", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;
	// 连接数据库
	dsn := config.GetDSN(&cfg.MySQLConfig)
	err = dao.InitMySQL(dsn)
	if err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}
	defer dao.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetupRouter()
	_ = r.Run()
}
