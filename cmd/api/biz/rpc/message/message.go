package message

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/message"
	"github.com/PICOF/simple-tiktok/kitex_gen/message/messageservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

const sendType = "1"

var (
	Client      messageservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("message")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitMessage()
}

func InitMessage() {
	c, err := messageservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func SendMessage(ctx context.Context, req *tiktokapi.MessageRequest, userId int64) (list *message.MessageResponse, err error) {
	toUserId, _ := strconv.ParseInt(req.GetToUserID(), 10, 64)
	var rpcReq = &message.MessageRequest{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: req.ActionType == sendType,
		Content:    req.Content,
		SendTime:   time.Now().UnixMilli(),
	}
	list, err = Client.SendMessage(ctx, rpcReq)
	return
}
func GetChatRecord(ctx context.Context, req *tiktokapi.ChatRecordRequest, userId int64) (resp *message.ChatRecordResponse, err error) {
	toUserId, _ := strconv.ParseInt(req.GetToUserID(), 10, 64)
	var rpcReq = &message.ChatRecordRequest{
		UserId:     userId,
		ToUserId:   toUserId,
		LatestTime: time.Now().UnixMilli(),
	}
	resp, err = Client.GetChatRecord(ctx, rpcReq)
	return
}
