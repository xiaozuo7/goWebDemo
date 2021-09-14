package utils

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("F://project//goWebDemo")
	err := viper.ReadInConfig()
	if err != nil {
		panic("初始化配置文件发生错误" + err.Error())
	}
}
