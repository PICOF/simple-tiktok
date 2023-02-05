package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
)

func GetUserById(ctx context.Context, userId int64, queryId int64) (user *user.UserInfo, err error) {
	var userInfo = &TUserInfo{}
	err = dal.DB.WithContext(ctx).First(userInfo, queryId).Error
	if err != nil {
		return nil, err
	}
	user, err = ConvertUserInfo(ctx, userId, userInfo)
	return
}
func ConvertUserInfo(ctx context.Context, userId int64, info *TUserInfo) (res *user.UserInfo, err error) {
	isFollow, err := IsFollow(ctx, info.Id, userId)
	if err != nil {
		return nil, err
	}
	res = &user.UserInfo{
		Id:            info.Id,
		Name:          info.Username,
		FollowCount:   info.FollowCount,
		FollowerCount: info.FollowerCount,
		IsFollow:      isFollow,
	}
	return
}
