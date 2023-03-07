package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Username        string
	Password        string
	Host            string
	Port            int
	Database        string
	OpenAIUrl       string
	OpenAIAuthToken string
)

func init() {
	// 设置配置文件名和路径
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	// 将配置项赋值给全局变量
	Username = viper.GetString("username")
	Password = viper.GetString("password")
	Host = viper.GetString("host")
	Port = viper.GetInt("port")
	Database = viper.GetString("database")
	OpenAIUrl = viper.GetString("openaiurl")
	OpenAIAuthToken = viper.GetString("openaiauthtoken")
}
