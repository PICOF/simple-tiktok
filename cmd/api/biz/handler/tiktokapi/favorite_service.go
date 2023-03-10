// Code generated by hertz generator.

package tiktokapi

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/rpc/favorite"

	tiktokapi "github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// LikeAction .
// @router /douyin/favorite/action/ [POST]
func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.LikeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userId := c.GetInt64("user_id")
	resp, err := favorite.Action(ctx, &req, userId)

	if err != nil && resp == nil {
		msg := "请求远程服务时出错"
		c.JSON(consts.StatusInternalServerError, tiktokapi.LoginResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		})
	} else {
		c.JSON(consts.StatusOK, resp)
	}
}

// GetLikeList .
// @router /douyin/favorite/list/ [GET]
func GetLikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.LikeListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userId := c.GetInt64("user_id")
	resp, err := favorite.GetFavoriteList(ctx, &req, userId)

	if err != nil && resp == nil {
		msg := "请求远程服务时出错"
		c.JSON(consts.StatusInternalServerError, tiktokapi.LoginResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		})
	} else {
		c.JSON(consts.StatusOK, resp)
	}
}
