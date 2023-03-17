package gpt

import (
	"encoding/json"
	"fmt"
	"geekdemo/model/dto"
	"geekdemo/utils"

	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

type EmbeddingModel struct {
	Input string `json:"input"`
	Model string `json:"model"`
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
	url := utils.OpenAIUrl + `/v1/embeddings`

	// 配置日志

	data, _ := ctx.GetRawData()
	var parameters map[string]interface{}
	// 包装成json 数据
	_ = json.Unmarshal(data, &parameters)

	input := parameters["input"].(string)
	// n := m["n"].(int)
	// size := m["size"].(string)
	embeddingModel := EmbeddingModel{
		Model: "text-embedding-ada-002",
		Input: input,
	}
	bytes, err := json.Marshal(embeddingModel)
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
