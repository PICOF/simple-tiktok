package main

import (
	"context"
	user "github.com/PICOF/simple-tiktok/kitex_gen/user/userservice"
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
	userConfig  *viper.Viper
	address     string
	serviceName string
)

func init() {
	userConfig = config.GetConfig("user")
	address = userConfig.GetString("server.address")
	serviceName = userConfig.GetString("server.serviceName")
}

func main() {
	f, err := os.OpenFile("./user_output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	//r, err := etcd.NewEtcdRegistry(constant.ETCDAddress)
	//if err != nil {
	//	panic(err)
	//}
	//addr, err := net.ResolveTCPAddr("tcp", address)
	//if err != nil {
	//	panic(err)
	//}
	//provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(serviceName),
	//	provider.WithExportEndpoint(constant.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	//svr := user.NewServer(new(UserServiceImpl),
	//	server.WithServiceAddr(addr),
	//	server.WithRegistry(r),
	//	server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
	//	server.WithMuxTransport(),
	//	server.WithSuite(tracing.NewServerSuite()),
	//	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	//)
	addr, _ := net.ResolveTCPAddr("tcp", address)
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	//TODO 换成 ETCD 加其它中间件，但是这玩意只支持 netpoll……
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
