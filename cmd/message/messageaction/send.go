package messageaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/message"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func SendMessage(ctx context.Context, request *message.MessageRequest) (resp *message.MessageResponse, err error) {
	var code int64
	var msg string
	code, msg = constant.Failed.GetInfo()
	if request.ActionType {
		t := time.UnixMilli(request.SendTime)
		err = operation.SendMessage(ctx, request.UserId, request.ToUserId, request.Content, t)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to send message: %v", err)
		} else {
			code, msg = constant.Success.GetInfo()
		}
	}
	resp = &message.MessageResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
	return
}
