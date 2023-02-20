package user

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var Client userservice.Client

func init() {
	InitUser()
}

func InitUser() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		constant.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.ServerServiceName}),
	)
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
func GetInfo(ctx context.Context, userId int64) (info *user.UserInfoResponse, err error) {
	var rpcReq = &user.UserInfoRequest{
		UserId:  userId,
		QueryId: userId,
	}
	info, err = Client.GetUserInfo(ctx, rpcReq)
	return
}
