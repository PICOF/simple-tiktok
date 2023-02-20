package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func SendMessage(ctx context.Context, userId int64, toUserId int64, context string, sendTime time.Time) (err error) {
	err = dal.DB.WithContext(ctx).Create(&TMessage{
		UserId:   userId,
		ToUserId: toUserId,
		Content:  context,
		SendTime: sendTime,
	}).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to create comment: %v", err)
		return
	}
	return
}

func GetMessage(ctx context.Context, userId int64, friendId int64, latestTime int64) (list []TMessage, err error) {
	t := time.UnixMilli(latestTime).Format("2006-01-02 15:04:05")
	err = dal.DB.WithContext(ctx).Where("user_id = ? AND to_user_id = ? AND send_time > ?", friendId, userId, t).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get message list: %v", err)
		return
	}
	return
}
