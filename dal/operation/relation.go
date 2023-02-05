package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"gorm.io/gorm"
)

func IsFollow(ctx context.Context, userId int64, followerId int64) (res bool, err error) {
	err = dal.DB.WithContext(ctx).Where("user_id = ? AND follower_id = ?", userId, followerId).First(&TFollowList{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return
	}
	res = true
	return
}
