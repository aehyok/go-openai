package routes

import (
	"geekdemo/service/user"

	"github.com/gin-gonic/gin"
)

func UserApi(v1 *gin.RouterGroup) {
	v1.GET("ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})

	v1.POST("/user/add", Wrapper(user.AddUser))

	// 删

	v1.GET("/user/delete/:id", Wrapper(user.DeleteUser))

	// 改
	v1.POST("/user/update/:id", Wrapper(user.UpdateUser))

	// 查
	// 第一种：条件查询，
	v1.GET("/user/list/:name", Wrapper(user.ListUserByName))

	// 第二种：全部查询 / 分页查询
	v1.GET("/user/list", Wrapper(user.ListUser))

	v1.GET("/user/:id", Wrapper(user.GetUser))

}
