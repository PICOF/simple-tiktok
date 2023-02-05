package main

import (
	feed "github.com/PICOF/simple-tiktok/kitex_gen/feed/feedservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/spf13/viper"
	"log"
	"net"
)

var (
	feedConfig  *viper.Viper
	address     string
	serviceName string
)

func init() {
	feedConfig = config.GetConfig("feed")
	address = feedConfig.GetString("server.address")
	serviceName = feedConfig.GetString("server.serviceName")
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", address)
	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
