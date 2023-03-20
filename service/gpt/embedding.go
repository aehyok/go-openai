package service

import (
	"encoding/json"
	"fmt"
	"geekdemo/model/dto"
	"geekdemo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type EmbeddingModel struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type EmbeddingRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// EmbeddingResponse is the response from a Create embeddings request.
type EmbeddingResponse struct {
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
	Model  string      `json:"model"`
	Usage  Usage       `json:"usage"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// GetEmbeddings godoc
// @Summary		数据转换，转换为向量数据
// @Description
// @Tags			GPT
//
//	@Produce		json
//
// @Router       /openai/getEmbeddings [post]
func GetEmbeddings(ctx *gin.Context) dto.ResponseResult {
	// 配置日志

	data, _ := ctx.GetRawData()
	var parameters map[string]interface{}
	// 包装成json 数据
	_ = json.Unmarshal(data, &parameters)

	input := parameters["input"].(string)
	// n := m["n"].(int)
	// size := m["size"].(string)
	var response = GetEmbeddingApi(input)
	var obj map[string]interface{}
	if err := json.Unmarshal(response, &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)
	return dto.SetResponseData(obj)
}

// func updateByJson(c *gin.Context) dto.ResponseResult {
// 	var json []map[string]string
// 	if err := c.Bind(&json); err != nil {
// 		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// return
// 		return dto.SetResponseFailure("error")
// 	}
// 	if len(json) == 0 {
// 		// c.JSON(http.StatusBadRequest, gin.H{"error": "json is empty"})
// 		// return
// 		return dto.SetResponseFailure("json is empty")
// 	}
// 	//数据向量化
// 	points := make([]qdrant.Point, 0)
// 	for _, v := range json {
// 		embeddingRequest := openai.EmbeddingRequest{
// 			Input: v["text"],
// 			Model: openai.TextEmbeddingAda002,
// 		}
// 		response, err := openai.SendEmbeddings(embeddingRequest)
// 		if err != nil {
// 			common.Logger.Error(err.Error())
// 			c.JSON(http.StatusOK, common.Error(err.Error()))
// 			return
// 		}
// 		points = append(points, qdrant.Point{
// 			ID:      uuid.New().String(),
// 			Payload: v,
// 			Vector:  response.Data[0].Embedding,
// 		})
// 	}
// 	pr := qdrant.PointRequest{
// 		Points: points,
// 	}

// 	//存储
// 	err := qdrant.CreatePoints(common.GlobalObject.Qdrant.CollectionName, pr)
// 	if err != nil {
// 		common.Logger.Error(err.Error())
// 		c.JSON(http.StatusOK, common.Error(err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, common.Success(nil))
// }

func GetEmbeddingApi(input string) []byte {
	embeddingModel := EmbeddingModel{
		Model: "text-embedding-ada-002",
		Input: input,
	}
	url := utils.OpenAIUrl + `/v1/embeddings`
	bytes, err := json.Marshal(embeddingModel)
	if err != nil {
		fmt.Println("Error:", err)
		// return dto.SetResponseFailure("调用openai发生错误")
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
		// return dto.SetResponseFailure("调用openai发生错误")
	}

	fmt.Println("Status:", resp.StatusCode())

	return resp.Body()
}

func UploadJsonData(c *gin.Context) dto.ResponseResult {
	var jsonData []map[string]string
	if err := c.Bind(&jsonData); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
		return dto.SetResponseFailure("error")
	}
	if len(jsonData) == 0 {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "json is empty"})
		// return
		return dto.SetResponseFailure("json is empty")
	}
	//数据向量化
	points := make([]Point, 0)
	for _, v := range jsonData {
		input := v["text"]
		response := GetEmbeddingApi(input)
		fmt.Println(response, "response----response")
		var embeddingResponse EmbeddingResponse
		json.Unmarshal(response, &embeddingResponse)
		points = append(points, Point{
			ID:      uuid.New().String(),
			Payload: v,
			Vector:  embeddingResponse.Data[0].Embedding,
		})
	}
	pr := PointRequest{
		Points: points,
	}

	//存储
	err := CreatePoints(utils.QdrantCollectName, pr)
	if err != nil {
		// common.Logger.Error(err.Error())
		// c.JSON(http.StatusOK, common.Error(err.Error()))
		// return
		return dto.SetResponseFailure("数据上传发生错误")
	}
	// c.JSON(http.StatusOK, common.Success(nil))
	return dto.SetResponseSuccess("数据上传成功")
}

type ChatMeMessage struct {
	Id   string `json:"id"`
	Text string `json:"content"`
}

func ChatMe(c *gin.Context) dto.ResponseResult {
	var message ChatMeMessage
	if err := c.Bind(&message); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}
	response := GetEmbeddingApi(message.Text)

	json.Unmarshal(response, &response)
	fmt.Println(response, "response----response")
	var embeddingResponse EmbeddingResponse
	json.Unmarshal(response, &embeddingResponse)
	params := make(map[string]interface{})
	params["exact"] = false
	params["hnsw_ef"] = 128

	sr := PointSearchRequest{
		Params:      params,
		Vector:      embeddingResponse.Data[0].Embedding,
		Limit:       3,
		WithPayload: true,
	}

	//查询相似的
	res, err := SearchPoints(utils.QdrantCollectName, sr)
	if err != nil {
		// common.Logger.Error(err.Error())
		// c.JSON(http.StatusOK, common.Error(err.Error()))
		// return
	}

	//组装本地数据
	localData := ""
	for i, v := range res {
		re := v.Payload.(map[string]interface{})
		localData += "\n"
		localData += strconv.Itoa(i)
		localData += "."
		localData += re["title"].(string)
		localData += ":"
		localData += re["text"].(string)
	}
	messages := make([]ChatCompletionMessage, 0)
	q := "使用以下段落来回答问题，如果段落内容不相关就返回未查到相关信息：\"" + message.Text + "\""
	q += localData

	system := ChatCompletionMessage{
		Role:    "system",
		Content: "你是一个医院问诊机器人",
	}
	demo_q := "使用以下段落来回答问题：\"成人头疼，流鼻涕是感冒还是过敏？\"\n1. 普通感冒：您会出现喉咙发痒或喉咙痛，流鼻涕，流清澈的稀鼻涕（液体），有时轻度发热。\n2. 常年过敏：症状包括鼻塞或流鼻涕，鼻、口或喉咙发痒，眼睛流泪、发红、发痒、肿胀，打喷嚏。"
	demo_a := "成人出现头痛和流鼻涕的症状，可能是由于普通感冒或常年过敏引起的。如果病人出现咽喉痛和咳嗽，感冒的可能性比较大；而如果出现口、喉咙发痒、眼睛肿胀等症状，常年过敏的可能性比较大。"
	user1 := ChatCompletionMessage{
		Role:    "user",
		Content: demo_q,
	}
	assistant := ChatCompletionMessage{
		Role:    "assistant",
		Content: demo_a,
	}
	user := ChatCompletionMessage{
		Role:    "user",
		Content: q,
	}

	messages = append(messages, system)
	messages = append(messages, user1)
	messages = append(messages, assistant)
	messages = append(messages, user)
	var chatResponse = GetChatCompletionsApi(messages)
	var obj map[string]interface{}
	if err := json.Unmarshal(chatResponse, &obj); err != nil {
		panic(err)
	}
	fmt.Println("Body:", obj)

	// 最后我通过一个方法进行统一返回参数处理
	return dto.SetResponseData(obj)
}
