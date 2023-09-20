package routes

import (
	"geekdemo/service/cdlog"

	"github.com/gin-gonic/gin"
)

func LogApi(v1 *gin.RouterGroup) {
	v1.GET("/log/getListByVersion", Wrapper(cdlog.GetListByVersion))
}
