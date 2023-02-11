package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func PublishComment(ctx context.Context, input TComment) (comment TComment, err error) {
	err = dal.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Create(&input).Error
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to create comment: %v", err)
			return err
		}
		err = tx.Model(&TVideoInfo{
			PublicAttr: PublicAttr{
				Id: input.VideoId,
			},
		}).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to decrease comment count: %v", err)
			return err
		}
		return nil
	})
	comment = input
	return
}
func DeleteComment(ctx context.Context, commentId int64, videoId int64, userId int64) (err error) {
	err = dal.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Delete(&TComment{
			UserId:     userId,
			PublicAttr: PublicAttr{Id: commentId},
		}).Error
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to delete comment: %v", err)
			return err
		}
		err = tx.Model(&TVideoInfo{
			PublicAttr: PublicAttr{
				Id: videoId,
			},
		}).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to decrease comment count: %v", err)
			return err
		}
		return nil
	})
	return
}
func GetCommentList(ctx context.Context, videoId int64) (comments []TComment, err error) {
	err = dal.DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get comment list: %v", err)
		return
	}
	return
}
