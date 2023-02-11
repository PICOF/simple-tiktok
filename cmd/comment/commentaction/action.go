package commentaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func PublishComment(ctx context.Context, request *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	var code int64
	var msg string
	var input, res operation.TComment
	var comm *comment.CommentInfo
	input = operation.TComment{
		UserId:     request.UserId,
		VideoId:    request.VideoId,
		Content:    *request.CommentText,
		GmtCreated: time.Now(),
	}
	res, err = operation.PublishComment(ctx, input)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to make likes: %v", err)
		code, msg = constant.Failed.GetInfo()
	} else {
		code, msg = constant.Success.GetInfo()
		comm, err = ConvertCommentInfo(ctx, res, request.UserId, request.UserId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to make likes: %v", err)
			code, msg = constant.Failed.GetInfo()
		}
	}
	resp = &comment.CommentResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		Comment:    comm,
	}
	return
}
func DeleteComment(ctx context.Context, request *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	var code int64
	var msg string
	err = operation.DeleteComment(ctx, *request.CommentId, request.VideoId, request.UserId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to delete comment: %v", err)
		code, msg = constant.Failed.GetInfo()
	} else {
		code, msg = constant.Success.GetInfo()
	}
	resp = &comment.CommentResponse{
		StatusCode: code,
		StatusMsg:  &msg,
	}
	return
}
func ActionHandler(ctx context.Context, request *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	if request.ActionType {
		return PublishComment(ctx, request)
	} else {
		return DeleteComment(ctx, request)
	}
}
