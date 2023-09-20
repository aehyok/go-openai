package routes

import (
	"geekdemo/service/geek"

	"github.com/gin-gonic/gin"
)

func GeekApi(v1 *gin.RouterGroup) {
	v1.POST("/geek/alllist", geek.GeekList)

	v1.GET("/geek/getCourseType", Wrapper(geek.GetGeekCourseType))

	v1.POST("/geek/getCourse", Wrapper(geek.GetGeekCourse))

	v1.POST("/geek/getArticle", Wrapper(geek.GetGeekArticle))

	v1.GET("/geek/getArticleContent", Wrapper(geek.GetGeekArticleContent))

	v1.GET("/geek/getListByVersion", Wrapper(geek.GetListByVersion))
}
