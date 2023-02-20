package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-sql-driver/mysql"
)

func GetUser(ctx context.Context, queryId int64) (user TUserInfo, err error) {
	err = dal.DB.WithContext(ctx).First(&user, queryId).Error
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while getting user information: %v", err)
		return
	}
	return
}
func WriteUserInfo(ctx context.Context, info *TUserInfo) (userId int64, err error) {
	err = dal.DB.WithContext(ctx).Create(info).Error
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			klog.CtxWarnf(ctx, "Duplicate user name causes write failure: username: %s", info.Username)
		} else {
			klog.CtxErrorf(ctx, "An error occurred while writing user data: %v", err)
		}
	}
	userId = info.Id
	return
}
func GetUserByName(ctx context.Context, username string) (user *TUserInfo, err error) {
	user = &TUserInfo{}
	err = dal.DB.WithContext(ctx).Where("username = ?", username).First(user).Error
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while getting user information: %v", err)
		return
	}
	return
}
func GetWorkCount(ctx context.Context, userId int64) (count int64, err error) {
	err = dal.DB.WithContext(ctx).Model(&TVideoInfo{}).Where("author_id = ?", userId).Count(&count).Error
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while getting user work_count: %v", err)
		return
	}
	return
}
