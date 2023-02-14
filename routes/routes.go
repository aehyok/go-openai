package routes

import (
	"geekdemo/service/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 测试
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "请求成功",
	// 	})
	// })

	// 业务码约定：正确200，错误400

	// 增
	r.POST("/user/add", user.AddUser)

	// 删
	// 1. 找到对应的id对应的条目
	// 2. 判断id是否存在
	// 3. 从数据库中删除 or 返回id没有找到

	// Restful编码规范
	r.DELETE("/user/delete/:id", user.DeleteUser)

	// 改
	r.PUT("/user/update/:id", user.UpdateUser)

	// 查
	// 第一种：条件查询，
	r.GET("/user/list/:name", user.ListUserByName)

	// 第二种：全部查询 / 分页查询
	r.GET("/user/list", user.ListUser)
	return r
}
