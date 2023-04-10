package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"geekdemo/model/dto"
	"geekdemo/utils"
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

func Send(httpMethod string, suffix string, reqBytes []byte) (body []byte, err error) {
	req, err := http.NewRequest(httpMethod, utils.QdrantConfig.Url+suffix, bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", utils.QdrantConfig.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	return
}

func SendMessage(httpMethod string, suffix string) dto.ResponseResult {
	url := utils.QdrantConfig.Url + `/dashboard/billing/credit_grants`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod(httpMethod)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+utils.QdrantConfig.ApiKey)

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
