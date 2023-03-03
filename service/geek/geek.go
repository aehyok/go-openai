package geek

import (
	"geekdemo/model"
	"geekdemo/model/dto"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GeekCourseModel struct {
	// 页数
	Page int `json:"page"`
	// 条数
	Limit int `json:"limit"`
	// 类型
	TypeId string `json:"typeId"`
}

// GetGeekCourseType godoc
// @Summary		课程大分类
// @Description	查看大课程的类别
// @Tags			geek
// @Param			Authorization header string true "token"
//
// @Accept			json
//
//	@Produce		json
//
// @Router       /geek/GetCourseType [get]
func GetGeekCourseType(ctx *gin.Context) dto.ResponseResult {
	var dataList []model.GeekCourseType
	model.DB.Model(dataList).Find(&dataList)
	// var aa []model.GeekCourseType
	// model.DB.Model(aa).First(&aa)
	// log.Println("1", dataList, aa)
	if len(dataList) == 0 {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "没有查询到数据",
		// 	"code": 400,
		// 	"data": gin.H{},
		// })
		// return dto.SetResponseData(gin.H{
		// 	"docs": dataList,
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

// GetGeekCourse godoc
// @Summary		课程查看
// @Description	查看大分类下的课程
// @Tags			geek
// @Param			Authorization header string true "token"
// @Param			GeekCourseModel body GeekCourseModel true "页数"
//
// @Accept			json
// @Produce		    json
//
// @Router       /geek/GetGeekCourse [post]
func GetGeekCourse(ctx *gin.Context) dto.ResponseResult {
	var dataList []model.GeekCourse
	var geekCourseModel GeekCourseModel
	if err := ctx.ShouldBindJSON(&geekCourseModel); err != nil {
		return dto.SetResponseFailure("err--err--err--err")
	}

	limit := geekCourseModel.Limit
	page := geekCourseModel.Page
	typeId, _ := strconv.Atoi(geekCourseModel.TypeId)

	// limit, _ := strconv.Atoi(ctx.PostForm("limit"))
	// page, _ := strconv.Atoi(ctx.PostForm("page"))
	// typeId, _ := strconv.Atoi(ctx.PostForm("typeId"))

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
	if typeId == 0 {
		model.DB.Model(dataList).Count(&total).Limit(limit).Offset(offsetVal).Find(&dataList)
	} else {
		model.DB.Model(dataList).Where("typeId = ? ", typeId).Count(&total).Limit(limit).Offset(offsetVal).Find(&dataList)
	}
	if len(dataList) == 0 {
		return dto.SetResponseFailure("没有查询到数据")
	} else {
		return dto.SetResponseData(gin.H{
			"docs":  dataList,
			"total": total,
			"page":  page,
			"limit": limit,
		})
	}
}

func GeekList(ctx *gin.Context) {
	var dataList []model.GeekProduct
	// 查询全部数据 or 查询分页数据
	// strconv.Atoi() 字符串转整型
	// ctx.Query("limit") url截取请求参数
	// ctx.PostForm("limit") 请求体截取请求参数
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
