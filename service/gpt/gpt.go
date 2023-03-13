package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"geekdemo/model/dto"
	"geekdemo/utils"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

func GetUsage(ctx *gin.Context) dto.ResponseResult {
	url := utils.OpenAIUrl + `/dashboard/billing/credit_grants`

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
	return dto.SetResponseData(obj)
}

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

type ImageModel struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

func GetImageGenerations(ctx *gin.Context) dto.ResponseResult {
	url := utils.OpenAIUrl + `/v1/images/generations`

	data, _ := ctx.GetRawData()
	var m map[string]interface{}
	// 包装成json 数据
	_ = json.Unmarshal(data, &m)

	prompt := m["prompt"].(string)
	// n := m["n"].(int)
	// size := m["size"].(string)
	imageModel := ImageModel{
		Prompt: prompt,
		N:      1,
		Size:   "512x512",
	}
	bytes, err := json.Marshal(imageModel)
	if err != nil {
		fmt.Println("Error:", err)
		return dto.SetResponseFailure("调用openai发生错误")
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+utils.OpenAIAuthToken)
	req.SetBody(bytes)
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
	return dto.SetResponseData(obj)
}

func GetSpeechToText(ctx *gin.Context) dto.ResponseResult {
	url := utils.OpenAIUrl + `/v1/audio/transcriptions`

	// 1. 创建 multipart.Writer 对象
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// 2. 添加表单数据和文件
	writer.WriteField("model", "whisper-1")
	file, err := os.Open("test.m4a")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	part, err := writer.CreateFormFile("file", "test.m4a")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	// 3. 关闭 multipart.Writer
	writer.Close()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetContentType(writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+utils.OpenAIAuthToken)
	req.SetRequestURI(url)
	req.SetBody(body.Bytes())

	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("failed to send request:", err)
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)

	// 释放资源
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return dto.SetResponseData(obj)
}

type ChatModel struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func GetChatCompletions(ctx *gin.Context) dto.ResponseResult {
	data, _ := ctx.GetRawData()
	var m map[string]interface{}
	// 包装成json 数据
	_ = json.Unmarshal(data, &m)

	content := m["content"].(string)
	chatModel := ChatModel{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "user", Content: content},
		},
	}
	bytes, err := json.Marshal(chatModel)

	fmt.Println(string(bytes), "bytes")
	if err != nil {
		fmt.Println("error:", err)
		return dto.SetResponseFailure("数据转换错误")
	}

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
	req.SetBody(bytes)

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
	return dto.SetResponseData(obj)

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
