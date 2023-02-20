package useraction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/pkg/jwt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ReqToTUser(req *user.RegisterRequest) (info *operation.TUserInfo) {
	info = &operation.TUserInfo{
		Username: req.Username,
		Password: req.Password,
	}
	return
}

func DataToResp(ctx context.Context, userId int64) (resp *user.RegisterResponse, err error) {
	var token string
	token, err = jwt.JWTUtil.CreateToken(userId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to create token: %v", err)
		return
	}
	resp = PackRegisterResponse(userId, token, constant.Success)
	return
}
func ConvertUserInfo(ctx context.Context, userId int64, info operation.TUserInfo) (res *user.UserInfo, err error) {
	isFollow, err := operation.IsFollow(ctx, info.Id, userId)
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while converting user information: %v", err)
		return
	}
	count, err := operation.GetWorkCount(ctx, info.Id)
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while converting user information: %v", err)
		return nil, err
	}
	favoriteCount, err := operation.GetFavoriteCount(ctx, info.Id)
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while converting user information: %v", err)
		return nil, err
	}
	res = &user.UserInfo{
		Id:            info.Id,
		Name:          info.Username,
		FollowCount:   info.FollowCount,
		FollowerCount: info.FollowerCount,
		IsFollow:      isFollow,
		WorkCount:     &count,
		FavoriteCount: &favoriteCount,
	}
	return
}
