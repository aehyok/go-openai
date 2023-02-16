package routes

import (
	_ "geekdemo/docs"
	"geekdemo/service/geek"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// apiHandler := gin.WrapH(v1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	// 增

	v1 := r.Group("/api/v1")
	{
		UserApi(v1)
	}

	// 获取productList
	r.POST("/user/Alllist", geek.ListProduct)

	return r
}
