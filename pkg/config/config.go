package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func GetConfig(configName string) (v *viper.Viper) {
	v = viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")
	err := v.ReadInConfig() // 查找并读取配置文件
	if err != nil {         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		klog.CtxInfof(context.Background(), "config file changes: %s", e.String())
	})
	return
}
