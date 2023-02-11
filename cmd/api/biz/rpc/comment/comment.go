package comment

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment/commentservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
	"strconv"
)

var (
	Client      commentservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

const publishType = "1"

func init() {
	Config = config.GetConfig("comment")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitComment()
}

func InitComment() {
	c, err := commentservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func Action(ctx context.Context, req *tiktokapi.CommentRequest, userId int64) (resp *comment.CommentResponse, err error) {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, err
	}
	var rpcReq *comment.CommentRequest
	if isPublish := req.ActionType == publishType; isPublish {
		rpcReq = &comment.CommentRequest{
			UserId:      userId,
			VideoId:     videoId,
			ActionType:  isPublish,
			CommentText: req.CommentText,
		}
	} else {
		commentId, err := strconv.ParseInt(*req.CommentID, 10, 64)
		if err != nil {
			return nil, err
		}
		rpcReq = &comment.CommentRequest{
			UserId:     userId,
			VideoId:    videoId,
			ActionType: isPublish,
			CommentId:  &commentId,
		}
	}
	resp, err = Client.CommentAction(ctx, rpcReq)
	return
}

func GetCommentList(ctx context.Context, req *tiktokapi.CommentListRequest, userId int64) (resp *comment.CommentListResponse, err error) {
	var videoId int64
	videoId, err = strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return
	}
	var rpcReq *comment.CommentListRequest
	rpcReq = &comment.CommentListRequest{
		UserId:  userId,
		VideoId: videoId,
	}
	resp, err = Client.GetCommentList(ctx, rpcReq)
	return
}
