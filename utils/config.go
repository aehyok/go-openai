package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

// var (
// 	Username          string
// 	Password          string
// 	Host              string
// 	Port              int
// 	OpenAIType        string
// 	Database          string
// 	OpenAIUrl         string
// 	OpenAIAuthToken   string
// 	QdrantUrl         string
// 	QdrantApiKey      string
// 	QdrantCollectName string
// )

type Database struct {
	UserName string
	Password string
	Host     string
	Port     int
	Database string
}

type Qdrant struct {
	Url         string
	CollectName string
	ApiKey      string
}

type GPT struct {
	Type   string
	Url    string
	ApiKey string
}

var (
	DatabaseConfig Database
	QdrantConfig   Qdrant
	GptConfig      GPT
)

func init() {
	// 设置配置文件名和路径
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	if err := viper.UnmarshalKey("databaseconfig", &DatabaseConfig); err != nil {
		fmt.Println("Error database config:", err)
		return
	}

	if err := viper.UnmarshalKey("qdrantconfig", &QdrantConfig); err != nil {
		fmt.Println("Error database config:", err)
		return
	}

	if err := viper.UnmarshalKey("gptConfig", &GptConfig); err != nil {
		fmt.Println("Error gpt config:", err)
		return
	}

	// // 将配置项赋值给全局变量
	// Username = viper.GetString("username")
	// Password = viper.GetString("password")
	// Host = viper.GetString("host")
	// Port = viper.GetInt("port")
	// Database = viper.GetString("database")

	// OpenAIType = viper.GetString("openaitype") // openai 或者 azure
	// OpenAIUrl = viper.GetString("openaiurl")
	// OpenAIAuthToken = viper.GetString("openaiauthtoken")
	// QdrantUrl = viper.GetString("qdranturl")
	// QdrantApiKey = viper.GetString("qdrantapikey")
	// QdrantCollectName = viper.GetString("qdrantcollectname")
}
