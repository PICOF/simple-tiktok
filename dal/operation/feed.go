package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func GetVideoList(ctx context.Context, latestTime int64) (list []TVideoInfo, err error) {
	t := time.UnixMilli(latestTime).Format("2006-01-02 15:04:05")
	err = dal.DB.WithContext(ctx).Limit(30).Order("publish_time desc").Where("publish_time < ?", t).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get video list: %v", err)
		return
	}
	return
}
