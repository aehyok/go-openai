package main

import (
	"goservices/model"
	"goservices/routes"
	// "gorm.io/driver/sqlite"
)

func main() {
	model.Database()
	// 接口
	r := routes.NewRouter()
	// 端口号
	PORT := "3001"
	r.Run(":" + PORT)
}
