package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"tiktok/server/util"
)

const FILEPATH = "config/config.yaml"

func LoadConfig() {
	viper.SetConfigFile(FILEPATH)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("获取配置文件失败,error:", err)
	}
}

// 在配置文件被修改后更新相应配置结构体
func renewConfig() {
	util.GetLoggerConfig()
}

func BootConfigMonitor() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		util.MsgLog.Println("配置文件发生变化:", e.String())
		renewConfig()
		util.MsgLog.Println("相关配置刷新成功！")
	})
}
