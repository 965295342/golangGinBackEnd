package main

import (
	"main/db"
	"main/entity"
	"main/routes"
)

func main() {
	//连接数据库
	err := db.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer db.Close()
	//绑定模型
	db.GetDb().AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8081的项目
	r.Run(":8081")
}
