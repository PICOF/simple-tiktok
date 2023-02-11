package useraction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/pkg/bcrypt"
	"github.com/PICOF/simple-tiktok/pkg/jwt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func VerifyUserInfo(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {
	var info *operation.TUserInfo
	info, err = operation.GetUserByName(ctx, request.Username)
	if err != nil {
		klog.CtxErrorf(ctx, "Error getting user information: %v", err)
		return
	}
	if bcrypt.ComparePassword(info.Password, request.Password) {
		var token string
		token, err = jwt.JWTUtil.CreateToken(info.Id)
		if err != nil {
			klog.CtxErrorf(ctx, "Error creating token: %v", err)
		}
		code, msg := constant.Success.GetInfo()
		resp = &user.LoginResponse{
			StatusCode: code,
			StatusMsg:  &msg,
			UserId:     &info.Id,
			Token:      &token,
		}
	} else {
		klog.CtxWarnf(ctx, "Authentication failed")
		code, _ := constant.Failed.GetInfo()
		resp = &user.LoginResponse{
			StatusCode: code,
		}
	}
	return
}
