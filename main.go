package main

import (
	"fmt"
	"geekdemo/model"
	"geekdemo/routes"
	"geekdemo/utils"
)

// @title CICD自动化 API
// @version 0.0.1
// @description cicd time
// @name aehyok
// @BasePath /api/v1
func main() {
	// 定义要执行的命令列表
	// commands := []*exec.Cmd{
	// 	exec.Command("echo", "第一个命令"),
	// 	exec.Command("ls", "-l"),
	// 	exec.Command("echo", "第三个命令"),
	// 	exec.Command("", "", "pwd"),
	// }

	// 遍历并执行每个命令
	// for i, cmd := range commands {
	// 	if i == 3 {
	// 		//E:\work\git-refactor\mp-h5
	// 		cmd.Dir = "/e/work/git-refactor/mp-h5"
	// 	}
	// 	// 运行命令
	// 	out, err := cmd.CombinedOutput()
	// 	if err != nil {
	// 		fmt.Printf("命令%d执行出错: %s\n", i, err)
	// 		continue
	// 	}
	// 	// 打印命令的输出
	// 	fmt.Printf("命令%d输出:\n%s\n", i, string(out))
	// }

	// strings := "cd /E/work/git-refactor/mp-h5 && yarn build"
	// cmd := exec.Command("bash", "-c", strings)

	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// error := cmd.Run()
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// 数据库初始化
	model.Database()

	fmt.Println("token==", utils.DatabaseConfig.UserName)

	// 接口路由
	r := routes.NewRouter()
	// 端口号
	PORT := "3001"
	r.Run(":" + PORT)
}
