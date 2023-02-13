package messageaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/dal/redis"
	"github.com/PICOF/simple-tiktok/kitex_gen/message"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

func GetRecord(ctx context.Context, request *message.ChatRecordRequest) (resp *message.ChatRecordResponse, err error) {
	var code int64
	var msg string
	var list []operation.TMessage
	code, msg = constant.Failed.GetInfo()
	list, err = operation.GetMessage(ctx, request.UserId, request.ToUserId, request.GetLatestTime())
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get chat record: %v", err)
	} else {
		code, msg = constant.Success.GetInfo()
	}
	if length := len(list); length > 0 {
		redis.Redis.Set("chat_latestTime_"+strconv.FormatInt(request.UserId, 10)+strconv.FormatInt(request.ToUserId, 10), list[length-1].SendTime.UnixMilli(), time.Hour*24)
	}
	resp = &message.ChatRecordResponse{
		StatusCode:  code,
		StatusMsg:   &msg,
		MessageList: MessageFormat(list),
	}
	return
}

func MessageFormat(list []operation.TMessage) (ret []*message.MessageInfo) {
	for _, v := range list {
		info := &message.MessageInfo{
			Id:         v.Id,
			Content:    v.Content,
			CreateTime: v.SendTime.UnixMilli(),
		}
		ret = append(ret, info)
	}
	return
}
