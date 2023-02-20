package routes

import (
	_ "geekdemo/docs"
	"geekdemo/middleware"
	"geekdemo/service/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// apiHandler := gin.WrapH(v1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	// 增

	// Group创建一个新的路由器组。您应该添加所有具有公共中间件或相同路径前缀的路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/login", Wrapper(user.Login))
		v1.Use(middleware.JWT())
		{
			UserApi(v1)
		}
		GeekApi(v1)
	}

	return r
}
