// Code generated by hertz generator.

package tiktokapi

import (
	"context"

	tiktokapi "github.com/PICOF/simple-tiktok/biz/model/tiktokapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetVedioList .
// @router /douyin/feed/ [GET]
func GetVedioList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.FeedResponse)

	c.JSON(consts.StatusOK, resp)
}
