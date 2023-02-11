package user

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/kitex_gen/user/userservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
	"strconv"
)

var (
	Client      userservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("user")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitUser()
}

func InitUser() {
	//r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	//if err != nil {
	//	panic(err)
	//}
	//provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(constant.ServerServiceName),
	//	provider.WithExportEndpoint(constant.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	//c, err := userservice.NewClient(
	//	serviceName,
	//	client.WithResolver(r),
	//	client.WithMuxConnection(1),
	//	client.WithSuite(tracing.NewClientSuite()),
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.ServerServiceName}),
	//)
	c, err := userservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
	if err != nil {
		panic(err)
	}
	Client = c
}

func Register(ctx context.Context, req *tiktokapi.RegisterRequest) (list *user.RegisterResponse, err error) {
	var rpcReq = &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}
	list, err = Client.UserRegister(ctx, rpcReq)
	return
}
func Login(ctx context.Context, req *tiktokapi.LoginRequest) (login *user.LoginResponse, err error) {
	var rpcReq = &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	login, err = Client.UserLogin(ctx, rpcReq)
	return
}
func GetInfo(ctx context.Context, req *tiktokapi.UserInfoRequest, userId int64) (info *user.UserInfoResponse, err error) {
	queryId, _ := strconv.ParseInt(req.GetUserID(), 10, 64)
	var rpcReq = &user.UserInfoRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	info, err = Client.GetUserInfo(ctx, rpcReq)
	return
}
