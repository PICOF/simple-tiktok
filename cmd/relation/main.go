package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	relation "github.com/PICOF/simple-tiktok/kitex_gen/relation/relationservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/PICOF/simple-tiktok/pkg/logger"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
)

var (
	relationConfig *viper.Viper
	address        string
)

func init() {
	relationConfig = config.GetConfig("service")
	address = relationConfig.GetString("service.relation.address")
}

func main() {
	f, err := os.OpenFile("./relation_output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	r, err := etcd.NewEtcdRegistry(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.RelationServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.RelationServiceName}),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
