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
// @Param			cmdStr query string true "命令行"
//
// @Accept			json
//
//	@Produce		json
//
// @Router       /log/cmd [get]
func Cmd(ctx *gin.Context) dto.ResponseResult {
	cmdStr := ctx.Query("cmdStr")
	fmt.Println(cmdStr)
	//window cmd
	// strings := "cd /E/work/git-refactor/mp-h5 && yarn build"
	// strings := "cd /data/work/git-refactor/mp-h5 && yarn build" // (yarn build success)

	//npm yarn pnpm success
	// cmd.Env = append(os.Environ(), "PATH=/usr/local/lib/nodejs/bin:$PATH")

	// git version git tag
	// cmd.Env = append(os.Environ(), "PATH=/usr/bin:$PATH")

	// 增加安全性，可以设置执行命令的列表，不在其中的则不允许执行
	strings := "cd /data/work/github/zx-deploy && " + cmdStr
	fmt.Println(strings)
	cmd := exec.Command("bash", "-c", strings)
	cmd.Env = append(os.Environ(), "PATH=/usr/local/lib/nodejs/bin:$PATH")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	error := cmd.Run()
	if error != nil {
		fmt.Println(error)
	}

	return dto.SetResponseFailure("没有查询到数据")
}
