package routes

import (
	"geekdemo/service/geek"

	"github.com/gin-gonic/gin"
)

func GeekApi(v1 *gin.RouterGroup) {
	v1.GET("/geek/GetCourseType", Wrapper(geek.GetGeekCourseType))

	v1.POST("/geek/GetGeekCourse", Wrapper(geek.GetGeekCourse))

	v1.GET("/geek/GetGeekArticle", Wrapper(geek.GetGeekArticle))

	v1.GET("/geek/GetArticleContent", Wrapper(geek.GetGeekArticleContent))
}
