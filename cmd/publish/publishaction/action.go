package publishaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish"
	"github.com/PICOF/simple-tiktok/pkg/md5"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func PublishAction(ctx context.Context, request *publish.PublishRequest) (resp *publish.PublishResponse, err error) {
	var code int64
	var msg string
	now := time.Now()
	name, _ := md5.GetFileMd5(request.GetData())
	videoInfo := &operation.TVideoInfo{
		AuthorId:    request.GetUserId(),
		PublishTime: now,
		Title:       request.GetTitle(),
	}
	err = operation.PublishVideo(ctx, videoInfo, request.GetData(), name)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to publish video: %v", err)
		code, msg = constant.Failed.GetInfo()
	} else {
		code, msg = constant.Success.GetInfo()
	}
	resp = &publish.PublishResponse{
		StatusCode: code,
		StatusMsg:  &msg,
	}
	return
}
