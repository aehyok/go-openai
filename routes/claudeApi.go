package routes

import (
	claude "geekdemo/service/claude"

	"github.com/gin-gonic/gin"
)

func ClaudeApi(v1 *gin.RouterGroup) {
	v11 := v1.Group("/claude")
	{
		v11.POST("/getChats", Wrapper(claude.GetChats))
	}
}
