package service

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
	"github.com/pandodao/tokenizer-go"
	"github.com/valyala/fasthttp"
)

// GetChats godoc
// @Summary  Claude3
// @Description 不支持stream模式
// @Tags   GPT
//
// @Produce  json
//
// @Router       /claude/getChats [post]
func GetChats(ctx *gin.Context) dto.ResponseResult {
	url := utils.ClaudeConfig.Url + `/v1/messages`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", utils.ClaudeConfig.ApiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.SetBody([]byte(`{"messages": [
		{"role": "user", "content": "Hello, world"}], "max_tokens": 2000,  "model": "claude-3-opus-20240229"}`))

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
	return dto.SetResponseData(obj)
}

type ImageModel struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

// GetImageGenerations godoc
// @Summary  根据文字描述生成图片
// @Description 暂时写死只能生成一张
// @Tags   GPT
//
// @Produce  json
//
// @Router       /openai/getImageGenerations [post]
func GetImageGenerations(ctx *gin.Context) dto.ResponseResult {
	url := utils.GptConfig.Url + `/v1/images/generations`

	data, _ := ctx.GetRawData()
	var m map[string]interface{}
	// 包装成json 数据
	_ = json.Unmarshal(data, &m)

	prompt := m["prompt"].(string)
	// n := m["n"].(int)
	// size := m["size"].(string)
	imageModel := ImageModel{
		Prompt: prompt,
		N:      10,
		Size:   "1024x1024",
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
	req.Header.Set("Authorization", "Bearer "+utils.GptConfig.ApiKey)
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
// @Summary  根据上传的语音转换为文字
// @Description 限制最大为25M
// @Tags   GPT
//
// @Produce  json
//
// @Router       /openai/getSpeechToText [post]
func GetSpeechToText(ctx *gin.Context) dto.ResponseResult {
	url := utils.GptConfig.Url + `/v1/audio/transcriptions`

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
	req.Header.Set("Authorization", "Bearer "+utils.GptConfig.ApiKey)
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
	Model          string                  `json:"model"`
	MaxTokens      int                     `json:"max_tokens"`
	Messages       []ChatCompletionMessage `json:"messages"`
	ResponseFormat ResponseFormat          `json:"response_format"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}
type ChatFunctionModel struct {
	Model         string                  `json:"model"`
	MaxTokens     int                     `json:"max_tokens"`
	Messages      []ChatCompletionMessage `json:"messages"`
	Functions     []FunctionSchema        `json:"functions"`
	Function_call string                  `json:"function_call"`
}

type FunctionSchema struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type Parameters struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Required   []string   `json:"required"`
}

type Properties struct {
	Location Location `json:"location"`
	Unit     Unit     `json:"unit"`
}

type Location struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Unit struct {
	Type string `json:"type"`
	// Enums string `json:"enum"`
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`

	// This property isn't in the official documentation, but it's in
	// the documentation for the official library for python:
	// - https://github.com/openai/openai-python/blob/main/chatml.md
	// - https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
	Name string `json:"name,omitempty"`
}

// GetChatCompletions godoc
// @Summary  GPT-3.5模型聊天对话(支持openai和azure两种接口)
// @Description 暂时不支持上下文
// @Tags   GPT
//
// @Produce  json
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

	messages := make([]ChatCompletionMessage, 0)

	user := ChatCompletionMessage{
		Role:    "user",
		Content: content,
	}

	messages = append(messages, user)

	resp := GetChatCompletionsApi(messages, "gpt-4-0613")
	// resp := GetChatCompletionsApi(messages, "gpt-3.5-turbo-0613")
	// 组装openai 接口的参数实体
	// gpt-4   gpt-3.5-turbo

	// 将返回对象中的body数据转换为json数据
	var obj map[string]interface{}
	if err := json.Unmarshal(resp, &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)

	// 最后我通过一个方法进行统一返回参数处理
	return dto.SetResponseData(obj)
}

func GetChatCompletionWithFunctions(ctx *gin.Context) dto.ResponseResult {

	// 通过GetRawData获取前端传递的JSON数据结构
	data, _ := ctx.GetRawData()

	// 将data数据 包装成json数据
	var m map[string]interface{}
	_ = json.Unmarshal(data, &m)

	// 这里我定义的参数是content获取传入的参数
	content := m["content"].(string)
	if content == "" {
		content = "中国的首都是那个城市，以及这个城市的天气情况"
	}
	messages := make([]ChatCompletionMessage, 0)

	user := ChatCompletionMessage{
		Role:    "user",
		Content: content,
	}

	messages = append(messages, user)

	//gpt-3.5-turbo-0613
	resp := GetChatCompletionsApi_WithFunction(messages, "gpt-4-1106-preview")
	// 组装openai 接口的参数实体
	// gpt-4   gpt-3.5-turbo

	// 将返回对象中的body数据转换为json数据
	var obj map[string]interface{}
	if err := json.Unmarshal(resp, &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)

	// 最后我通过一个方法进行统一返回参数处理
	return dto.SetResponseData(obj)
}

type WeatherInfo struct {
	Location    string   `json:"location"`
	Temperature string   `json:"temperature"`
	Unit        string   `json:"unit"`
	Forecast    []string `json:"forecast"`
}

func get_current_weather(location string, unit string) (string, error) {
	weatherInfo := WeatherInfo{
		Location:    location,
		Temperature: "72",
		Unit:        unit,
		Forecast:    []string{"sunny", "windy"},
	}

	weatherInfoJson, err := json.Marshal(weatherInfo)
	if err != nil {
		return "", err
	}

	return string(weatherInfoJson), nil
}

func GetChatCompletionsApi(messages []ChatCompletionMessage, apiModel string) []byte {
	var model string
	if utils.GptConfig.Type == "openai" {
		model = "gpt-4-1106-preview"
	} else {
		model = "gpt-35-turbo"
	}

	responseFormat := ResponseFormat{
		Type: "json_object",
	}

	chatModel := ChatModel{
		Model:          model,
		MaxTokens:      2000,
		Messages:       messages,
		ResponseFormat: responseFormat,
	}

	// 将实体结构转换为byte数组
	bytes, err := json.Marshal(chatModel)

	fmt.Println(string(bytes), "bytes")
	if err != nil {
		fmt.Println("error:", err)
		// return dto.SetResponseFailure("数据转换错误")
	}

	// openai接口地址，可通过代理处理
	var url string
	if utils.GptConfig.Type == "openai" {
		url = utils.GptConfig.Url + `/v1/chat/completions`
	} else {
		url = utils.GptConfig.Url + `/openai/deployments/ChatGPT/chat/completions?api-version=2023-03-15-preview`
	}

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
	if utils.GptConfig.Type == "openai" {
		req.Header.Set("Authorization", "Bearer "+utils.GptConfig.ApiKey)
	} else {
		req.Header.Set("api-key", utils.GptConfig.ApiKey)
	}

	//gpt-3.5-turbo-0301
	req.SetBody(bytes)

	// 这里跟上面AcquireRequest 类似的，一个是请求对象，一个是返回对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 通过fasthttp.Do真正的发起对象
	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		// return dto.SetResponseFailure("调用openai发生错误")
		// ctx.JSON(200, gin.H{"data": "调用openai发生错误"})

	}
	return resp.Body()
}

func GetChatCompletionsApi_WithFunction(messages []ChatCompletionMessage, apiModel string) []byte {
	var model string
	if utils.GptConfig.Type == "openai" {
		// model = "gpt-3.5-turbo"
		model = apiModel
	} else {
		model = "gpt-35-turbo"
	}

	var functions []FunctionSchema

	funcByte := []byte(`
		[
			{
				"name": "get_current_weather",
				"description": "Get the current weather in a given location",
				"parameters": {
						"type": "object",
						"properties": {
								"location": {
										"type": "string",
										"description": "The city and state, e.g. San Francisco, CA"
								},
								"unit": {"type": "string" }
						},
						"required": ["location"]
				}
			}
		]
	`)

	err := json.Unmarshal(funcByte, &functions)
	if err != nil {
		fmt.Println("error:", err)
	}

	chatModel := ChatFunctionModel{
		Model:         model,
		MaxTokens:     2000,
		Messages:      messages,
		Functions:     functions,
		Function_call: "auto",
	}

	// 将实体结构转换为byte数组
	bytes, err := json.Marshal(chatModel)

	fmt.Println(string(bytes), "bytes")
	if err != nil {
		fmt.Println("error:", err)
		// return dto.SetResponseFailure("数据转换错误")
	}

	// openai接口地址，可通过代理处理
	var url string
	if utils.GptConfig.Type == "openai" {
		url = utils.GptConfig.Url + `/v1/chat/completions`
	} else {
		url = utils.GptConfig.Url + `/openai/deployments/ChatGPT/chat/completions?api-version=2023-03-15-preview`
	}

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
	if utils.GptConfig.Type == "openai" {
		req.Header.Set("Authorization", "Bearer "+utils.GptConfig.ApiKey)
	} else {
		req.Header.Set("api-key", utils.GptConfig.ApiKey)
	}

	//gpt-3.5-turbo-0301
	req.SetBody(bytes)

	// 这里跟上面AcquireRequest 类似的，一个是请求对象，一个是返回对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 通过fasthttp.Do真正的发起对象
	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		// return dto.SetResponseFailure("调用openai发生错误")
		// ctx.JSON(200, gin.H{"data": "调用openai发生错误"})

	}
	return resp.Body()
}
func GetTokenizer(ctx *gin.Context) dto.ResponseResult {
	// 通过GetRawData获取前端传递的JSON数据结构
	data, _ := ctx.GetRawData()

	// 将data数据 包装成json数据
	var m map[string]interface{}
	_ = json.Unmarshal(data, &m)

	// 这里我定义的参数是content获取传入的参数
	content := m["content"].(string)
	t := tokenizer.MustCalToken(content)
	return dto.SetResponseData(t)
}
