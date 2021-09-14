package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var BasePath string

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(BasePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic("初始化配置文件发生错误" + err.Error())
	}
}

func init() {
	// 1.初始化程序根目录
	if curPath, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		fmt.Println("初始化程序根目录错误: " + err.Error())

	}
}
