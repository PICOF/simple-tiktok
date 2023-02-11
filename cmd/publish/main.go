package main

import (
	"context"
	publish "github.com/PICOF/simple-tiktok/kitex_gen/publish/publishservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/PICOF/simple-tiktok/pkg/logger"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
)

var (
	feedConfig  *viper.Viper
	address     string
	serviceName string
)

func init() {
	feedConfig = config.GetConfig("publish")
	address = feedConfig.GetString("server.address")
	serviceName = feedConfig.GetString("server.serviceName")
}

func main() {
	f, err := os.OpenFile("./publish_output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			klog.CtxFatalf(context.Background(), "Error closing log file: %v", err)
		}
	}(f)
	logger.SetLogger()
	klog.SetOutput(f)
	addr, _ := net.ResolveTCPAddr("tcp", address)
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
