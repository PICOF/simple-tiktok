package rpc

import (
	"context"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/kitex_gen/user/userservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
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
	c, err := userservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func GetUserInfo(ctx context.Context, userId int64, queryId int64) (userInfo *user.UserInfo, err error) {
	var rpcReq = &user.UserInfoRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	var resp *user.UserInfoResponse
	resp, err = Client.GetUserInfo(ctx, rpcReq)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get user info: %v", err)
		return
	}
	userInfo = resp.User
	return
}
