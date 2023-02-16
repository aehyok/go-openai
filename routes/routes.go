package routes

import (
	"geekdemo/service/user"
	"geekdemo/service/geek"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 增
	r.POST("/user/add", user.AddUser)

	// 删

	r.DELETE("/user/delete/:id", user.DeleteUser)

	// 改
	r.PUT("/user/update/:id", user.UpdateUser)

	// 查
	// 第一种：条件查询，
	r.GET("/user/list/:name", user.ListUserByName)

	// 第二种：全部查询 / 分页查询
	r.GET("/user/list", user.ListUser)

	// 获取productList
	r.POST("/user/Alllist", geek.ListProduct)

	return r
}

