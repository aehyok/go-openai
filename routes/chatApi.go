package routes

import (
	"geekdemo/service/gpt"

	"github.com/gin-gonic/gin"
)

func ChatApi(v1 *gin.RouterGroup) {
	v11 := v1.Group("/openai")
	{
		v11.POST("/getModels", Wrapper(gpt.GetModels))
		v11.POST("/getUsage", Wrapper(gpt.GetUsage))
		v11.POST("/getCompletions", Wrapper(gpt.GetCompletions))
		v11.POST("/getChatCompletions", Wrapper(gpt.GetChatCompletions))
		v11.POST("/getImageGenerations", Wrapper(gpt.GetImageGenerations))
		v11.POST("/getSpeechToText", Wrapper(gpt.GetSpeechToText))
		v11.POST("/getEmbeddings", Wrapper(gpt.GetEmbeddings))
		v11.POST("/tokenizer", Wrapper(gpt.GetTokenizer))
	}
}
