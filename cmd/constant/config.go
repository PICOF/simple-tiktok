package constant

import (
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/spf13/viper"
)

var (
	ServiceConfig     *viper.Viper
	ETCDAddress       []string
	ExportEndpoint    string
	ServerServiceName string
	ServerAddress     string
)

func init() {
	ServiceConfig = config.GetConfig("service")
	ETCDAddress = ServiceConfig.GetStringSlice("etcd.address")
	ExportEndpoint = ServiceConfig.GetString("exportEndpoint")
	ServerServiceName = ServiceConfig.GetString("server.name")
	ServerAddress = ServiceConfig.GetString("server.address")
}
