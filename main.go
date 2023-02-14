package main

import (
	"geekdemo/model"
	"geekdemo/routes"
)

func main() {
	// 数据库初始化
	model.Database()
	// 接口路由
	r := routes.NewRouter()
	// 端口号
	PORT := "3001"
	r.Run(":" + PORT)
}
