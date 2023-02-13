package comment

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment/commentservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"strconv"
)

var Client commentservice.Client

const publishType = "1"

func init() {
	InitComment()
}

func InitComment() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := commentservice.NewClient(
		constant.CommentServiceName,
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
