// Code generated by hertz generator.

package tiktokapi

import (
	"context"

	tiktokapi "github.com/PICOF/simple-tiktok/biz/model/tiktokapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendMessage .
// @router /douyin/message/action/ [POST]
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.MessageRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.MessageResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetChatRecord .
// @router /douyin/message/chat/ [GET]
func GetChatRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.ChatRecordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.ChatRecordResponse)

	c.JSON(consts.StatusOK, resp)
}