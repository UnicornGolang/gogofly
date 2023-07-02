package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	path, err := os.Getwd()
	if err != nil {
		panic("get project path error")
	}
	viper.AddConfigPath(path + "/config/")

	if err = viper.ReadInConfig(); err != nil {
		panic("read config error" + err.Error())
	}
}
