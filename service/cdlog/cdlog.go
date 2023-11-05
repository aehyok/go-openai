package cdlog

import (
	"fmt"
	"geekdemo/model"
	"geekdemo/model/dto"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// getListByVersion godoc
// @Summary		CICD日志
// @Description	根据版本查询日志
// @Tags			log
// @Param			version query string true "版本号"
// @Param			project query string true "项目"
//
// @Accept			json
//
//	@Produce		json
//
// @Router       /log/getListByVersion [get]
func GetListByVersion(ctx *gin.Context) dto.ResponseResult {
	version := ctx.Query("version")
	project := ctx.Query("project")
	var dataList []model.CicdLog
	sql := model.DB.Order("createTime desc").Where("version like ? ", "%"+version+"%")
	if project != "" {
		sql = sql.Where("project= ?", project)
	}
	sql.Find(&dataList)
	if len(dataList) == 0 {
		return dto.SetResponseFailure("没有查询到数据")
	} else {
		return dto.SetResponseData(dataList)
	}
}

// cmd godoc
// @Summary		根据参数执行命令
// @Description	执行命令
// @Tags			log
//
// @Accept			json
//
//	@Produce		json
//
// @Router       /log/cmd [get]
func Cmd(ctx *gin.Context) dto.ResponseResult {

	// cmdStr := `cd /E/work/git-refactor/mp-h5 && yarn build`

	// cmd := exec.Command("bash", "-c", cmdStr)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// // 获取输出
	// err := cmd.Run() // cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	strings := "cd /E/work/git-refactor/mp-h5 && yarn build"
	cmd := exec.Command("bash", "-c", strings)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	error := cmd.Run()
	if error != nil {
		fmt.Println(error)
	}

	return dto.SetResponseFailure("没有查询到数据")
}
