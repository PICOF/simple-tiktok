package favoriteaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/favorite"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FavoriteAction(ctx context.Context, request *favorite.LikeRequest) (resp *favorite.LikeResponse, err error) {
	var code int64
	var msg string
	err = operation.FavoriteAction(ctx, request.UserId, request.VideoId, request.ActionType)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to make likes: %v", err)
		code, msg = constant.Failed.GetInfo()
	} else {
		code, msg = constant.Success.GetInfo()
	}
	resp = &favorite.LikeResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
	return
}
