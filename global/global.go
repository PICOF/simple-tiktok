package global

import (
	"github.com/PICOF/simple-tiktok/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DOUYIN_VIPER  *viper.Viper
	DOUYIN_CONFIG config.Config
	DOUYIN_DB     *gorm.DB
)
