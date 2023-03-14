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

// GetUsage godoc
// @Summary		用户授信额度使用
// @Description	当前用户的总额度和使用额度
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getUsage [post]
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

// GetModels godoc
// @Summary		openai 所有开放的模型model
// @Description	列出所有模型
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getModels [post]
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

// GetCompletions godoc
// @Summary		GPT-3.0聊天对话模式
// @Description	不支持stream模式
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getCompletions [post]
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

// GetImageGenerations godoc
// @Summary		根据文字描述生成图片
// @Description	暂时写死只能生成一张
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getImageGenerations [post]
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

// GetSpeechToText godoc
// @Summary		根据上传的语音转换为文字
// @Description	限制最大为25M
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getSpeechToText [post]
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

// GetChatCompletions godoc
// @Summary		GPT-3.5模型聊天对话
// @Description	暂时不支持上下文
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getChatCompletions [post]
func GetChatCompletions(ctx *gin.Context) dto.ResponseResult {

	// 通过GetRawData获取前端传递的JSON数据结构
	data, _ := ctx.GetRawData()

	// 将data数据 包装成json数据
	var m map[string]interface{}
	_ = json.Unmarshal(data, &m)

	// 这里我定义的参数是content获取传入的参数
	content := m["content"].(string)

	// 组装openai 接口的参数实体
	chatModel := ChatModel{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "user", Content: content},
		},
	}

	// 将实体结构转换为byte数组
	bytes, err := json.Marshal(chatModel)

	fmt.Println(string(bytes), "bytes")
	if err != nil {
		fmt.Println("error:", err)
		return dto.SetResponseFailure("数据转换错误")
	}

	// openai接口地址，可通过代理处理
	url := utils.OpenAIUrl + `/v1/chat/completions`

	// 定义fasthttp请求对象
	req := fasthttp.AcquireRequest()

	// 使用defer关键字可以确保在函数返回之前，即使出现了错误，也会释放请求对象的内存，从而避免内存泄漏和浪费。
	// 当请求处理完成时，应该调用fasthttp.ReleaseRequest(req)来将请求对象返回给对象池。
	defer fasthttp.ReleaseRequest(req)

	// 设置请求的url地址，请求头，以及通过SetBody设置请求的参数
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Authorization", "Bearer "+utils.OpenAIAuthToken)
	//gpt-3.5-turbo-0301
	req.SetBody(bytes)

	// 这里跟上面AcquireRequest 类似的，一个是请求对象，一个是返回对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 通过fasthttp.Do真正的发起对象
	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		// return dto.SetResponseFailure("调用openai发生错误")
		ctx.JSON(200, gin.H{"data": "调用openai发生错误"})
	}

	fmt.Println("Status:", resp.StatusCode())

	// 将返回对象中的body数据转换为json数据
	var obj map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)

	// 最后我通过一个方法进行统一返回参数处理
	return dto.SetResponseData(obj)
}
