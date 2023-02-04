// Code generated by hertz generator.

package tiktokapi

import (
	"context"

	tiktokapi "github.com/PICOF/simple-tiktok/biz/model/tiktokapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FollowAction .
// @router /douyin/relation/action/ [POST]
func FollowAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.FollowRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.FollowResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetFollowList .
// @router /douyin/relation/follow/list/ [GET]
func GetFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.RelationListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.RelationListResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func GetFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.RelationListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.RelationListResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetFriendList .
// @router /douyin/relation/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.RelationListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.RelationListResponse)

	c.JSON(consts.StatusOK, resp)
}