package user

import (
	"errors"
	"fmt"
	"geekdemo/middleware"
	"geekdemo/model"
	"geekdemo/model/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginModel struct {
	// 账号
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
}

// AddUser godoc
//
//	@Summary		添加一个用户
//	@Description	添加一个用户
//	@Tags			user
//	@Accept			json
//	@Produce		json
//
// @Router       /user/add [post]
func AddUser(ctx *gin.Context) dto.ResponseResult {
	// 定义一个变量指向结构体
	var data model.BasicUser
	// 绑定方法
	err := ctx.ShouldBindJSON(&data)
	// 判断绑定是否有错误
	if err != nil {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "添加失败",
		// 	"data": gin.H{},
		// 	"code": "400",
		// })
		return dto.SetResponseFailure("添加失败")
	} else {
		// 数据库的操作
		model.DB.Create(&data) // 创建一条数据
		// ctx.JSON(200, gin.H{
		// 	"msg":  "添加成功",
		// 	"data": data,
		// 	"code": "200",
		// })
		return dto.SetResponseData(data)
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
// @Router       /user/delete/{id} [get]
func DeleteUser(ctx *gin.Context) dto.ResponseResult {
	var data []model.BasicUser
	// 接收id
	id := ctx.Param("id") // 如果有键值对形式的话用Query()
	// 判断id是否存在
	model.DB.Where("id = ? ", id).Find(&data)
	if len(data) == 0 {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "id没有找到，删除失败",
		// 	"code": 400,
		// })
		return dto.SetResponseFailure("id没有找到，删除失败")
	} else {
		// 操作数据库删除（删除id所对应的那一条）
		// db.Where("id = ? ", id).Delete(&data) <- 其实不需要这样写，因为查到的data里面就是要删除的数据
		model.DB.Delete(&data)

		// ctx.JSON(200, gin.H{
		// 	"msg":  "删除成功",
		// 	"code": 200,
		// })
		return dto.SetResponseSuccess("删除成功")
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
// @Router       /user/update [post]
func UpdateUser(ctx *gin.Context) dto.ResponseResult {
	// 1. 找到对应的id所对应的条目
	// 2. 判断id是否存在
	// 3. 修改对应条目 or 返回id没有找到
	var data model.BasicUser
	id := ctx.Param("id")
	// db.Where("id = ?", id).Find(&data) 可以这样写，也可以写成下面那样
	// 还可以再Where后面加上Count函数，可以查出来这个条件对应的条数
	model.DB.Select("id").Where("id = ? ", id).Find(&data)
	if data.Id == 0 {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "用户id没有找到",
		// 	"code": 400,
		// })
		return dto.SetResponseFailure("用户id没有找到")
	} else {
		// 绑定一下
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			// ctx.JSON(200, gin.H{
			// 	"msg":  "修改失败",
			// 	"code": 400,
			// })
			return dto.SetResponseSuccess("修改失败")
		} else {
			// db修改数据库内容
			model.DB.Where("id = ?", id).Updates(&data)
			// ctx.JSON(200, gin.H{
			// 	"msg":  "修改成功",
			// 	"code": 200,
			// })
			return dto.SetResponseSuccess("修改成功")
		}
	}
}

// GetUser godoc
// @Summary		查询用户
// @Description	根据传递的name查找来用户，可能返回一个或多个用户数据
// @Tags			user
//
// @Param id path int true "int valid"
//
//	@Accept			json
//	@Produce		json
//
// @Router       /user/{id} [get]
func GetUser(ctx *gin.Context) dto.ResponseResult {
	// 获取路径参数
	id := ctx.Param("id")
	var user model.BasicUser
	// 查询数据库

	if err := model.DB.First(&user, id).Error; err != nil {
		fmt.Println(err, "err111")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			fmt.Println(err, "true")
			return dto.SetResponseFailure("此用户不存在")
		} else {
			// 处理其他错误
			fmt.Println(err, "false")
			return dto.SetResponseFailure("发生错误，请重新输入")
		}
	} else {
		return dto.SetResponseData(user)
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
// @Router       /user/list/{name} [get]
func ListUserByName(ctx *gin.Context) dto.ResponseResult {
	// 获取路径参数
	name := ctx.Param("name")
	var dataList []model.BasicUser
	// 查询数据库
	model.DB.Where("name = ? ", name).Find(&dataList)
	// 判断是否查询到数据
	if len(dataList) == 0 {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "没有查询到数据",
		// 	"code": "400",
		// 	"data": gin.H{},
		// })
		return dto.SetResponseFailure("没有查询到数据")
	} else {
		// ctx.JSON(200, gin.H{
		// 	"msg":  "查询成功",
		// 	"code": "200",
		// 	"data": dataList,
		// })
		return dto.SetResponseData(dataList)
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
// @Router       /user/list [get]
func ListUser(ctx *gin.Context) dto.ResponseResult {
	var dataList []model.BasicUser
	// 查询全部数据 or 查询分页数据
	limit, _ := strconv.Atoi(ctx.Query("pageSize"))
	page, _ := strconv.Atoi(ctx.Query("pageNum"))

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
		// 		"list":     dataList,
		// 		"total":    total,
		// 		"pageNum":  pageNum,
		// 		"pageSize": pageSize,
		// 	},
		// })
		return dto.SetResponseData(gin.H{
			"docs":  dataList,
			"total": total,
			"page":  page,
			"limit": limit,
		})
	}
}

// Login godoc
// @Summary		登录
// @Description	根据用户的账号和密码
// @Tags			user
//
// @Param loginModel body LoginModel true "User information"
//
// @Accept json
// @Produce json
//
// @Router       /user/login [post]
func Login(ctx *gin.Context) dto.ResponseResult {
	var loginModel LoginModel
	if err := ctx.ShouldBindJSON(&loginModel); err != nil {
		return dto.SetResponseFailure("err--err--err--err")
	}

	fmt.Println(loginModel.Account, loginModel.Password, "login")
	var user model.BasicUser

	if err := model.DB.Where("account = ? AND password = ?", loginModel.Account, loginModel.Password).First(&user).Error; err != nil {
		fmt.Println(err, "err111")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			fmt.Println(err, "true")
			return dto.SetResponseFailure("账号和密码有误，请重新输入")
		} else {
			// 处理其他错误
			fmt.Println(err, "false")
			return dto.SetResponseFailure("发生错误，请重新输入")
		}
	} else {
		// 处理查询到的记录
		fmt.Println(err, "err--- false", user.Id, user.Account)
		token, _ := middleware.GenerateToken(user.Id, user.Account)
		var data = make(map[string]any)
		data["token"] = token
		data["id"] = user.Id
		return dto.SetResponseSuccess(data)
	}
}
