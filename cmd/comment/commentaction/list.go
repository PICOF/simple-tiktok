package commentaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetCommentList(ctx context.Context, request *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	var code int64
	var msg string
	var cList []operation.TComment
	cList, err = operation.GetCommentList(ctx, request.VideoId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get comment list: %v", err)
		code, msg = constant.Failed.GetInfo()
		resp = &comment.CommentListResponse{
			StatusCode: code,
			StatusMsg:  &msg,
		}
		return
	}
	var info []*comment.CommentInfo
	info, err = ConvertAllCommentInfo(ctx, cList, request.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get comment list: %v", err)
		code, msg = constant.Failed.GetInfo()
		resp = &comment.CommentListResponse{
			StatusCode: code,
			StatusMsg:  &msg,
		}
		return
	}
	code, msg = constant.Success.GetInfo()
	resp = &comment.CommentListResponse{
		StatusCode:  code,
		StatusMsg:   &msg,
		CommentList: info,
	}
	return
}
