package routes

import (
	"geekdemo/model/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回数据处理
type WrapperFuncType func(c *gin.Context) dto.ResponseResult

func Wrapper(handle WrapperFuncType) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result dto.ResponseResult = handle(c)
		c.PureJSON(http.StatusOK, result)
	}
}
