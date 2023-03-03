package routes

import (
	"geekdemo/service/geek"

	"github.com/gin-gonic/gin"
)

func GeekApi(v1 *gin.RouterGroup) {
	v1.POST("/geek/Alllist", geek.GeekList)

	v1.GET("/geek/GetCourseType", Wrapper(geek.GetGeekCourseType))

	v1.POST("/geek/GetGeekCourse", Wrapper(geek.GetGeekCourse))

	v1.POST("/geek/GetGeekArticle", Wrapper(geek.GetGeekArticle))
}
