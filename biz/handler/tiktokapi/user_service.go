// Code generated by hertz generator.

package tiktokapi

import (
	"context"
	"fmt"
	"github.com/PICOF/simple-tiktok/util"

	tiktokapi "github.com/PICOF/simple-tiktok/biz/model/tiktokapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"golang.org/x/crypto/bcrypt"
)

// UserRegist .
// @router /douyin/user/register/ [POST]
func UserRegist(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.RegisterResponse)

	//以后还要加逻辑，查询是不是已经注册过

	encode_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	fmt.Print(encode_password)
	if err != nil {
		resp.StatusCode = 1
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	userInfo := tiktokapi.UserInfo{
		ID:            1,
		Name:          req.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	J := util.NewJWT()
	token, err := J.CreateToken(userInfo)
	if err != nil {
		resp.StatusCode = 1
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.UserID = int64(userInfo.ID)
	resp.Token = token
	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.LoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req tiktokapi.UserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(tiktokapi.UserInfoResponse)

	c.JSON(consts.StatusOK, resp)
}