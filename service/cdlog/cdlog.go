package cdlog

import (
	"geekdemo/model"
	"geekdemo/model/dto"

	"github.com/gin-gonic/gin"
)

// getListByVersion godoc
// @Summary		CICD日志
// @Description	根据版本查询日志
// @Tags			log
// @Param			version query string true "版本号"
//
// @Accept			json
//
//	@Produce		json
//
// @Router       /log/getListByVersion [get]
func GetListByVersion(ctx *gin.Context) dto.ResponseResult {
	version := ctx.Query("version")
	var dataList []model.CicdLog
	model.DB.Order("createTime desc").Where("version like ? ", "%"+version+"%").Find(&dataList)
	if len(dataList) == 0 {
		return dto.SetResponseFailure("没有查询到数据")
	} else {
		return dto.SetResponseData(dataList)
	}
}
