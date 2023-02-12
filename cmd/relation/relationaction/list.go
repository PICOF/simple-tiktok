package relationaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/cmd/relation/rpc"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/relation"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetFollowerList(ctx context.Context, request *relation.RelationListRequest) (resp *relation.RelationListResponse, err error) {
	var code int64
	var msg string
	var userInfo []*user.UserInfo
	var list []operation.TFollowList
	code, msg = constant.Failed.GetInfo()
	resp = &relation.RelationListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
	}
	list, err = operation.GetFollowerList(ctx, request.QueryId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get follower list: %v", err)
		return
	}
	for _, v := range list {
		var info *user.UserInfo
		info, err = rpc.GetUserInfo(ctx, request.UserId, v.FollowerId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to get follower info: %v", err)
			return
		}
		userInfo = append(userInfo, info)
	}
	code, msg = constant.Success.GetInfo()
	resp = &relation.RelationListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		UserList:   userInfo,
	}
	return
}

func GetFollowList(ctx context.Context, request *relation.RelationListRequest) (resp *relation.RelationListResponse, err error) {
	var code int64
	var msg string
	var userInfo []*user.UserInfo
	var list []operation.TFollowList
	code, msg = constant.Failed.GetInfo()
	resp = &relation.RelationListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
	}
	list, err = operation.GetFollowList(ctx, request.QueryId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get follower list: %v", err)
		return
	}
	for _, v := range list {
		var info *user.UserInfo
		info, err = rpc.GetUserInfo(ctx, request.UserId, v.UserId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to get follower info: %v", err)
			return
		}
		userInfo = append(userInfo, info)
	}
	code, msg = constant.Success.GetInfo()
	resp = &relation.RelationListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		UserList:   userInfo,
	}
	return
}
func GetFriendList(ctx context.Context, request *relation.RelationListRequest) (resp *relation.RelationListResponse, err error) {
	var code int64
	var msg string
	var userInfo []*user.UserInfo
	var list []operation.TFollowList
	code, msg = constant.Failed.GetInfo()
	resp = &relation.RelationListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
	}
	list, err = operation.GetFriendList(ctx, request.QueryId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get follower list: %v", err)
		return
	}
	for _, v := range list {
		var info *user.UserInfo
		info, err = rpc.GetUserInfo(ctx, request.UserId, v.FollowerId)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to get follower info: %v", err)
			return
		}
		userInfo = append(userInfo, info)
	}
	code, msg = constant.Success.GetInfo()
	resp = &relation.RelationListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		UserList:   userInfo,
	}
	return
}
