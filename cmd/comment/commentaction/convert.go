package commentaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/comment/rpc"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ConvertCommentInfo(ctx context.Context, input operation.TComment, userId int64, queryId int64) (res *comment.CommentInfo, err error) {
	var userInfo *user.UserInfo
	userInfo, err = rpc.GetUserInfo(ctx, userId, queryId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get user info: %v", err)
		return
	}
	res = &comment.CommentInfo{
		Id:         input.Id,
		User:       userInfo,
		Content:    input.Content,
		CreateDate: input.GmtCreated.Format("1-2"),
	}
	return
}
func ConvertAllCommentInfo(ctx context.Context, input []operation.TComment, userId int64) (res []*comment.CommentInfo, err error) {
	var info *comment.CommentInfo
	for _, v := range input {
		info, err = ConvertCommentInfo(ctx, v, userId, v.UserId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to convert comment list: %v", err)
			return
		}
		res = append(res, info)
	}
	return
}
