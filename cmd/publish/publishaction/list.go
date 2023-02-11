package publishaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/cmd/publish/rpc"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish"
)

func GetPublishList(ctx context.Context, request *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	info, err := rpc.GetVideoListById(ctx, request.GetUserId(), request.GetQueryId())
	if err != nil {
		resp = PackListResponse(nil, constant.Failed)
	} else {
		resp = PackListResponse(info, constant.Success)
	}
	return
}
