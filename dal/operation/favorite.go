package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-sql-driver/mysql"
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
func FavoriteAction(ctx context.Context, userId int64, videoId int64, isFavorite bool) (err error) {
	var expr string
	if isFavorite {
		expr = "favorite_count + ?"
	} else {
		expr = "favorite_count - ?"
	}
	err = dal.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&TVideoInfo{
			PublicAttr: PublicAttr{
				Id: videoId,
			},
		}).UpdateColumn("favorite_count", gorm.Expr(expr, 1)).Error
		if err != nil {
			return err
		}
		if isFavorite {
			result := tx.WithContext(ctx).Create(&TLikedVideo{
				UserId:  userId,
				VideoId: videoId,
			})
			err = result.Error
			if err != nil {
				if err.(*mysql.MySQLError).Number == 1062 {
					klog.CtxWarnf(ctx, "Duplicate favorite causes write failure: video_id: %d", videoId)
				} else {
					klog.CtxErrorf(ctx, "An error occurred while writing user data: %v", err)
				}
				return err
			}
		} else {
			err = tx.WithContext(ctx).Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&TLikedVideo{}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}
func GetFavoriteList(ctx context.Context, userId int64) (list []TVideoInfo, err error) {
	var like []TLikedVideo
	err = dal.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&like).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get favorite video list: %v", err)
		return
	}
	var video TVideoInfo
	for _, v := range like {
		video.Id = v.VideoId
		err = dal.DB.WithContext(ctx).First(&video).Error
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to get favorite video info: %v", err)
			return
		}
		list = append(list, video)
	}
	return
}
func GetFavoriteCount(ctx context.Context, userId int64) (count int64, err error) {
	err = dal.DB.WithContext(ctx).Model(&TLikedVideo{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while getting user favorite_count: %v", err)
		return
	}
	return
}
