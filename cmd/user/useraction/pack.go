package useraction

import (
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
)

func PackRegisterResponse(userId int64, token string, status constant.Status) (response *user.RegisterResponse) {
	code, msg := status.GetInfo()
	response = &user.RegisterResponse{
		StatusCode: code,
		StatusMsg:  msg,
		UserId:     userId,
		Token:      token,
	}
	return
}
func PackUserInfoResponse(info *user.UserInfo) (response *user.UserInfoResponse) {
	code, msg := constant.Success.GetInfo()
	response = &user.UserInfoResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		User:       info,
	}
	return
}
