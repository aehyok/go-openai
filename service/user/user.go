package user

import (
	"geekdemo/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddUser godoc
//
//	@Summary		添加一个用户
//	@Description	添加一个用户
//	@Tags			user
//	@Accept			json
//	@Produce		json
//
// @Router       /api/v1/user/add [post]
func AddUser(ctx *gin.Context) {
	// 定义一个变量指向结构体
	var data model.List
	// 绑定方法
	err := ctx.ShouldBindJSON(&data)
	// 判断绑定是否有错误
	if err != nil {
		ctx.JSON(200, gin.H{
			"msg":  "添加失败",
			"data": gin.H{},
			"code": "400",
		})
	} else {
		// 数据库的操作
		model.DB.Create(&data) // 创建一条数据
		ctx.JSON(200, gin.H{
			"msg":  "添加成功",
			"data": data,
			"code": "200",
		})
	}
}

// DeleteUser godoc
// @Summary		删除用户
// @Description	根据传递的id查找来删除用户
// @Tags			user
//
//	@Accept			json
//	@Produce		json
//
// @Router       /api/v1/user/delete/{id} [get]
func DeleteUser(ctx *gin.Context) {
	var data []model.List
	// 接收id
	id := ctx.Param("id") // 如果有键值对形式的话用Query()
	// 判断id是否存在
	model.DB.Where("id = ? ", id).Find(&data)
	if len(data) == 0 {
		ctx.JSON(200, gin.H{
			"msg":  "id没有找到，删除失败",
			"code": 400,
		})
	} else {
		// 操作数据库删除（删除id所对应的那一条）
		// db.Where("id = ? ", id).Delete(&data) <- 其实不需要这样写，因为查到的data里面就是要删除的数据
		model.DB.Delete(&data)

		ctx.JSON(200, gin.H{
			"msg":  "删除成功",
			"code": 200,
		})
	}
}

// UpdateUser godoc
// @Summary		修改用户
// @Description	根据传递的用户信息进行更新，必须要传递已存在的id
// @Tags			user
//
//	@Accept			json
//	@Produce		json
//
// @Router       /api/v1/user/update [post]
func UpdateUser(ctx *gin.Context) {
	// 1. 找到对应的id所对应的条目
	// 2. 判断id是否存在
	// 3. 修改对应条目 or 返回id没有找到
	var data model.List
	id := ctx.Param("id")
	// db.Where("id = ?", id).Find(&data) 可以这样写，也可以写成下面那样
	// 还可以再Where后面加上Count函数，可以查出来这个条件对应的条数
	model.DB.Select("id").Where("id = ? ", id).Find(&data)
	if data.ID == 0 {
		ctx.JSON(200, gin.H{
			"msg":  "用户id没有找到",
			"code": 400,
		})
	} else {
		// 绑定一下
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg":  "修改失败",
				"code": 400,
			})
		} else {
			// db修改数据库内容
			model.DB.Where("id = ?", id).Updates(&data)
			ctx.JSON(200, gin.H{
				"msg":  "修改成功",
				"code": 200,
			})
		}
	}
}

// ListUserByName godoc
// @Summary		查询用户
// @Description	根据传递的name查找来用户，可能返回一个或多个用户数据
// @Tags			user
//
//	@Accept			json
//	@Produce		json
//
// @Router       /api/v1/user/list/{name} [get]
func ListUserByName(ctx *gin.Context) {
	// 获取路径参数
	name := ctx.Param("name")
	var dataList []model.List
	// 查询数据库
	model.DB.Where("name = ? ", name).Find(&dataList)
	// 判断是否查询到数据
	if len(dataList) == 0 {
		ctx.JSON(200, gin.H{
			"msg":  "没有查询到数据",
			"code": "400",
			"data": gin.H{},
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":  "查询成功",
			"code": "200",
			"data": dataList,
		})
	}
}

// ListUser godoc
// @Summary		用户列表
// @Description	根据传分页信息查询用户列表
// @Tags			user
//
//	@Accept			json
//	@Produce		json
//
// @Router       /api/v1/user/list [get]
func ListUser(ctx *gin.Context) {
	var dataList []model.List
	// 查询全部数据 or 查询分页数据
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	// 判断是否需要分页
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	offsetVal := (pageNum - 1) * pageSize // 固定写法 记住就行
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}

	// 返回一个总数
	var total int64

	// 查询数据库
	model.DB.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)

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
				"list":     dataList,
				"total":    total,
				"pageNum":  pageNum,
				"pageSize": pageSize,
			},
		})
	}
}