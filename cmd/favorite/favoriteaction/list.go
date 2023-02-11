package favoriteaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/favorite"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FavoriteList(ctx context.Context, request *favorite.LikeListRequest) (resp *favorite.LikeListResponse, err error) {
	var list []operation.TVideoInfo
	list, err = operation.GetFavoriteList(ctx, request.GetQueryId())
	if err != nil {
		klog.CtxErrorf(ctx, "Failed get favorite video list: %v", err)
		code, msg := constant.Failed.GetInfo()
		resp = &favorite.LikeListResponse{
			StatusCode: code,
			StatusMsg:  &msg,
		}
		return
	}
	var info []*feed.VideoInfo
	info, err = ConvertAllVideoInfo(ctx, request.GetUserId(), list)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to convert favorite list: %v", err)
		code, msg := constant.Failed.GetInfo()
		resp = &favorite.LikeListResponse{
			StatusCode: code,
			StatusMsg:  &msg,
		}
		return
	}
	code, msg := constant.Success.GetInfo()
	resp = &favorite.LikeListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		VideoList:  info,
	}
	return
}
