package operation

import (
	"context"
	"errors"
	"github.com/PICOF/simple-tiktok/dal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func IsFollow(ctx context.Context, userId int64, followerId int64) (res bool, err error) {
	if userId == followerId {
		res = false
		return
	}
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
func FollowAction(ctx context.Context, userId int64, toUserId int64) (err error) {
	err = dal.DB.Transaction(func(tx *gorm.DB) error {
		var follow TFollowList
		var status bool
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? AND follower_id = ?", userId, toUserId).
			First(&follow).
			Error
		if err != nil && err.Error() != "record not found" {
			return err
		}
		if err == nil {
			status = true
			tx.Model(&follow).Update("status", status)
		}
		err = tx.WithContext(ctx).Create(&TFollowList{
			UserId:     toUserId,
			FollowerId: userId,
			Status:     status,
		}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&TUserInfo{
			PublicAttr: PublicAttr{
				Id: userId,
			},
		}).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&TUserInfo{
			PublicAttr: PublicAttr{
				Id: toUserId,
			},
		}).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}
func UnfollowAction(ctx context.Context, userId int64, toUserId int64) (err error) {
	err = dal.DB.Transaction(func(tx *gorm.DB) error {
		var follow TFollowList
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? AND follower_id = ?", userId, toUserId).
			First(&follow).
			Error
		if err != nil && err.Error() != "record not found" {
			return err
		}
		if err == nil {
			tx.Model(&follow).Update("status", false)
		}
		result := tx.WithContext(ctx).Where("user_id = ? AND follower_id = ?", toUserId, userId).Delete(&TFollowList{})
		if result.Error != nil || result.RowsAffected == 0 {
			return errors.New("failed to find follow record")
		}
		err = tx.Model(&TUserInfo{
			PublicAttr: PublicAttr{
				Id: userId,
			},
		}).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&TUserInfo{
			PublicAttr: PublicAttr{
				Id: toUserId,
			},
		}).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func GetFollowList(ctx context.Context, userId int64) (list []TFollowList, err error) {
	err = dal.DB.WithContext(ctx).Where("follower_id = ?", userId).Find(&list).Error
	if err != nil {
		return
	}
	return
}

func GetFollowerList(ctx context.Context, userId int64) (list []TFollowList, err error) {
	err = dal.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&list).Error
	if err != nil {
		return
	}
	return
}

func GetFriendList(ctx context.Context, userId int64) (list []TFollowList, err error) {
	err = dal.DB.WithContext(ctx).Where("user_id = ? AND status = ?", userId, true).Find(&list).Error
	if err != nil {
		return
	}
	return
}
