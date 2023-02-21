package dto

import "net/http"

type ResponseResult struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponseData(data interface{}) ResponseResult {
	return ResponseResult{Code: http.StatusOK, Message: "success", Data: data}
}

func SetResponseFailure(message interface{}) ResponseResult {
	return ResponseResult{Code: http.StatusBadRequest, Message: message, Data: map[string]interface{}{}}
}

func SetResponseSuccess(data interface{}) ResponseResult {
	return ResponseResult{Code: http.StatusOK, Message: "success", Data: data}
}
