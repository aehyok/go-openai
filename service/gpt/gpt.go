package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"geekdemo/model/dto"
	"geekdemo/utils"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/valyala/fasthttp"
)

func GetModels(ctx *gin.Context) dto.ResponseResult {
	url := utils.OpenAIUrl + `/v1/models`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+utils.OpenAIAuthToken)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return dto.SetResponseFailure("调用openai发生错误")
	}

	fmt.Println("Status:", resp.StatusCode())
	var obj map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)
	return dto.SetResponseData(obj["data"])
}

func GetCompletions(ctx *gin.Context) dto.ResponseResult {
	url := utils.OpenAIUrl + `/v1/completions`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+utils.OpenAIAuthToken)
	req.SetBody([]byte(`{"prompt": "go语言实现hello world 并解析一下", "max_tokens": 2000,  "model": "text-davinci-003", "suffix": "欢迎再次体验" }`))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return dto.SetResponseFailure("调用openai发生错误")
	}

	fmt.Println("Status:", resp.StatusCode())
	fmt.Println("Status:", resp.StatusCode())
	var obj map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)
	return dto.SetResponseData(obj["choices"])
}

func GetChatCompletions(ctx *gin.Context) {
	url := utils.OpenAIUrl + `/v1/chat/completions`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Transfer-Encoding", "chunked")
	// req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Authorization", "Bearer "+utils.OpenAIAuthToken)
	//gpt-3.5-turbo-0301
	req.SetBody([]byte(`{"model": "gpt-3.5-turbo","max_tokens": 4096, "messages": [{"role": "user", "content": "go语言实现hello world 并解析一下"}] }`))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		// return dto.SetResponseFailure("调用openai发生错误")
		ctx.JSON(200, gin.H{"data": "调用openai发生错误"})
	}

	fmt.Println("Status:", resp.StatusCode())
	var obj map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)
	// return dto.SetResponseData(obj["choices"])
	ctx.JSON(200, gin.H{"data": obj})
	// for i := 0; i < int(resp.Body())(); i++ {
	// 	chunk := resp.BodyBuffer().Bytes()[i : i+1]
	// 	if _, err := ctx.Write(chunk); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	// // 释放请求和响应对象内存
	// fasthttp.ReleaseRequest(req)
	// fasthttp.ReleaseResponse(resp)
	// return dto.SetResponseData(obj["choices"])
}

func GptText(ctx *gin.Context) {
	config := openai.DefaultConfig(utils.OpenAIAuthToken)
	proxyUrl, err := url.Parse(utils.OpenAIUrl)
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
