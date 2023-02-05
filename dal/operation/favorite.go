package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"gorm.io/gorm"
)

func IsFavorite(ctx context.Context, userId int64, videoId int64) (res bool, err error) {
	err = dal.DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoId).First(&TLikedVideo{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return
	}
	res = true
	return
}
