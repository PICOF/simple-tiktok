package relationaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/relation"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FollowAction(ctx context.Context, request *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	var code int64
	var msg string
	code, msg = constant.Success.GetInfo()
	if request.GetActionType() {
		err = operation.FollowAction(ctx, request.UserId, request.ToUserId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to follow an user: %v", err)
			code, msg = constant.Failed.GetInfo()
		}
	} else {
		err = operation.UnfollowAction(ctx, request.UserId, request.ToUserId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to unfollow an user: %v", err)
			code, msg = constant.Failed.GetInfo()
		}
	}
	resp = &relation.FollowResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
	return
}
