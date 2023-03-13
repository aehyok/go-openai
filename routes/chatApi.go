package routes

import (
	"geekdemo/service/gpt"

	"github.com/gin-gonic/gin"
)

func ChatApi(v1 *gin.RouterGroup) {
	v11 := v1.Group("/openai")
	{
		v11.GET("/getModels", Wrapper(gpt.GetModels))
		v11.GET("/getUsage", Wrapper(gpt.GetUsage))
		v11.GET("/getCompletions", Wrapper(gpt.GetCompletions))
		v11.GET("/getImageGenerations", Wrapper(gpt.GetImageGenerations))
		v11.GET("/getChatCompletions", gpt.GetChatCompletions)
		v11.GET("/getSpeechToText", Wrapper(gpt.GetSpeechToText))
	}
}
