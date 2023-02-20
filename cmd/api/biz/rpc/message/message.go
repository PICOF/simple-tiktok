package message

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/message"
	"github.com/PICOF/simple-tiktok/kitex_gen/message/messageservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

const sendType = "1"

var Client messageservice.Client

func init() {
	InitMessage()
}

func InitMessage() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := messageservice.NewClient(
		constant.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.ServerServiceName}),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func SendMessage(ctx context.Context, req *tiktokapi.MessageRequest, userId int64) (list *message.MessageResponse, err error) {
	var rpcReq = &message.MessageRequest{
		UserId:     userId,
		ToUserId:   req.GetToUserID(),
		ActionType: req.ActionType == sendType,
		Content:    req.Content,
		SendTime:   time.Now().UnixMilli(),
	}
	list, err = Client.SendMessage(ctx, rpcReq)
	return
}
func GetChatRecord(ctx context.Context, req *tiktokapi.ChatRecordRequest, userId int64) (resp *message.ChatRecordResponse, err error) {
	var rpcReq = &message.ChatRecordRequest{
		UserId:     userId,
		ToUserId:   req.GetToUserID(),
		LatestTime: req.GetPreMsgTime(),
	}
	resp, err = Client.GetChatRecord(ctx, rpcReq)
	return
}
