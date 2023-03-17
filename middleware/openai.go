package middleware

import (
	"encoding/json"
	"fmt"
	"geekdemo/model/dto"
	"geekdemo/utils"

	"github.com/valyala/fasthttp"
)

func SendMessage(parameters map[string]interface{}) dto.ResponseResult {
	url := utils.OpenAIUrl + `/dashboard/billing/credit_grants`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod(parameters["HttpMethod"].(string))
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
