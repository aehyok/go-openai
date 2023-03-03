package main

import (
	"fmt"
	"geekdemo/model"
	"geekdemo/routes"
	"geekdemo/utils"
)

// @title 极客时间 API
// @version 0.0.1
// @description geek time
// @name aehyok
// @BasePath /api/v1
func main() {
	// 数据库初始化
	model.Database()

	fmt.Println("token==", utils.Username)

	// 接口路由
	r := routes.NewRouter()
	// 端口号
	PORT := "3001"
	r.Run(":" + PORT)
}
