package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PublishComment(ctx context.Context, input TComment) (comment TComment, err error) {
	err = dal.DB.WithContext(ctx).Create(&input).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to create comment: %v", err)
		return
	}
	comment = input
	return
}
