package routes

import (
	_ "geekdemo/docs"
	"geekdemo/service/geek"
	"geekdemo/service/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.POST("/user/add", user.AddUser)

		// 删

		v1.GET("/user/delete/:id", user.DeleteUser)

		// 改
		v1.POST("/user/update/:id", user.UpdateUser)

		// 查
		// 第一种：条件查询，
		v1.GET("/user/list/:name", user.ListUserByName)

		// 第二种：全部查询 / 分页查询
		v1.GET("/user/list", user.ListUser)
	}

	// apiHandler := gin.WrapH(v1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	// 增

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
