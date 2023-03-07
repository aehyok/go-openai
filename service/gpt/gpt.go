package gpt

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func GptText(ctx *gin.Context) {
	config := openai.DefaultConfig("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")
	proxyUrl, err := url.Parse("https://service-o0cr0c30-1253646855.hk.apigw.tencentcs.com/v1")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "go语言实现hello world 并解析",
				},
			},
		},
	)

	if err != nil {
		ctx.JSON(200, gin.H{"data": "error"})
	}
	fmt.Println(resp.Choices[0].Message.Content)
	ctx.JSON(200, gin.H{"data": resp.Choices[0]})
}
