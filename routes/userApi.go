package routes

import (
	"geekdemo/model/dto"
	"geekdemo/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserApi(v1 *gin.RouterGroup) {
	v1.GET("ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})

	v1.POST("/user/add", Wrapper(user.AddUser))

	// 删

	v1.GET("/user/delete/:id", Wrapper(user.DeleteUser))

	// 改
	v1.POST("/user/update/:id", Wrapper(user.UpdateUser))

	// 查
	// 第一种：条件查询，
	v1.GET("/user/list/:name", Wrapper(user.ListUserByName))

	// 第二种：全部查询 / 分页查询
	v1.GET("/user/list", Wrapper(user.ListUser))
}

// 返回数据处理
type WrapperFuncType func(c *gin.Context) dto.ResponseResult

func Wrapper(handle WrapperFuncType) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result dto.ResponseResult = handle(c)
		c.PureJSON(http.StatusOK, result)
	}
}
