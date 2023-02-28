package geek

import (
	"geekdemo/model"
	"geekdemo/model/dto"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGeekCourseType(ctx *gin.Context) dto.ResponseResult {
	var dataList []model.GeekCourseType
	model.DB.Model(dataList).Find(&dataList)

	log.Println("1", dataList)
	if len(dataList) == 0 {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "没有查询到数据",
		// 	"code": 400,
		// 	"data": gin.H{},
		// })
		return dto.SetResponseFailure("没有查询到数据")

	} else {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "查询成功",
		// 	"code": 200,
		// 	"data": gin.H{
		// 		"list": dataList,
		// 	},
		// })
		return dto.SetResponseData(gin.H{
			"docs": dataList,
		})
	}
}

func GeekList(ctx *gin.Context) {
	var dataList []model.GeekProduct
	// 查询全部数据 or 查询分页数据
	// strconv.Atoi() 字符串转整型   ctx.Query("limit") 截取请求参数
	cc := ctx.PostForm("limit")
	log.Println("11111", cc)
	limit, _ := strconv.Atoi(ctx.PostForm("limit"))
	page, _ := strconv.Atoi(ctx.PostForm("page"))

	// 判断是否需要分页
	if limit == 0 {
		limit = -1
	}
	if page == 0 {
		page = -1
	}

	offsetVal := (page - 1) * limit // 固定写法 记住就行
	if page == -1 && limit == -1 {
		offsetVal = -1
	}

	// 返回一个总数
	var total int64

	// 查询数据库
	model.DB.Model(dataList).Count(&total).Limit(limit).Offset(offsetVal).Find(&dataList)

	// 插入
	// supplier := model.GeekProduct{Title: "全栈", Remark: "牛批"}
	// model.DB.Create(&supplier)

	if len(dataList) == 0 {
		ctx.JSON(200, gin.H{
			"msg":  "没有查询到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":  "查询成功",
			"code": 200,
			"data": gin.H{
				"list":  dataList,
				"total": total,
				"page":  page,
				"limit": limit,
			},
		})
	}

}
