package logger

import (
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/spf13/viper"
)

var LoggerConfig *viper.Viper

const ConfigName = "logger"

func initLogger() {
	LoggerConfig = config.GetConfig(ConfigName)
}

func SetLogger() {
	logger := logrus.NewLogger()
	klog.SetLogger(logger)
	switch LoggerConfig.GetString("logLevel") {
	case "trace":
		klog.SetLevel(klog.LevelTrace)
	case "debug":
		klog.SetLevel(klog.LevelDebug)
	case "info":
		klog.SetLevel(klog.LevelInfo)
	case "warn":
		klog.SetLevel(klog.LevelWarn)
	case "error":
		klog.SetLevel(klog.LevelError)
	case "fatal":
		klog.SetLevel(klog.LevelFatal)
	default:
		klog.SetLevel(klog.LevelWarn)
	}
}
