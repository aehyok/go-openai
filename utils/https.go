package utils

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func SendRequest(url string, body []byte) []byte {
	// 定义fasthttp请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 设置请求的url地址，请求头，以及通过SetBody设置请求的参数
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	if GptConfig.Type == "openai" {
		req.Header.Set("Authorization", "Bearer "+GptConfig.ApiKey)
	} else {
		req.Header.Set("api-key", GptConfig.ApiKey)
	}

	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 通过fasthttp.Do真正的发起对象
	// Client
	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		// return dto.SetResponseFailure("调用openai发生错误")
		// ctx.JSON(200, gin.H{"data": "调用openai发生错误"})

	}
	return resp.Body()
}
