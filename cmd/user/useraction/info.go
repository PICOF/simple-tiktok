package useraction

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetUserInfo(ctx context.Context, request *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	var info operation.TUserInfo
	info, err = operation.GetUser(ctx, request.GetQueryId())
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get userinfo: %v", err)
		return
	}
	var userinfo *user.UserInfo
	userinfo, err = ConvertUserInfo(ctx, request.GetUserId(), info)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to convert userinfo: %v", err)
		return
	}
	resp = PackUserInfoResponse(userinfo)
	return
}
