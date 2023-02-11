package useraction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/pkg/bcrypt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func WriteUserInfo(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	var encrypt []byte
	encrypt, err = bcrypt.EncryptPassword(ctx, request.Password)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to encrypt password: %v", err)
		return
	}
	request.Password = string(encrypt)
	tUser := ReqToTUser(request)
	var info int64
	info, err = operation.WriteUserInfo(ctx, tUser)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to write user info: %v", err)
		return
	}
	resp, err = DataToResp(ctx, info)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to return registration response: %v", err)
		return
	}
	return
}
func WriteUserInfoHandler(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp, err = WriteUserInfo(ctx, request)
	if err != nil {
		resp = PackRegisterResponse(0, "", constant.Failed)
	}
	return
}
