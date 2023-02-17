package dto

import "net/http"

type ResponseResult struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponseData(data interface{}) ResponseResult {
	return ResponseResult{http.StatusOK, "success", data}
}

func SetResponseFailure(message interface{}) ResponseResult {
	return ResponseResult{http.StatusBadRequest, message, map[string]interface{}{}}
}

func SetResponseSuccess(message interface{}) ResponseResult {
	return ResponseResult{http.StatusOK, message, map[string]interface{}{}}
}
