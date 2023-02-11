package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment/commentservice"
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
	commentConfig *viper.Viper
	address       string
	serviceName   string
)

func init() {
	commentConfig = config.GetConfig("comment")
	address = commentConfig.GetString("server.address")
	serviceName = commentConfig.GetString("server.serviceName")
}

func main() {
	f, err := os.OpenFile("./comment_output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	svr := commentservice.NewServer(new(CommentServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
