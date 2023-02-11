package mw

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type checkFailed struct {
	StatusCode    int64  `json:"status_code"`
	StatusMessage string `json:"status_msg"`
}

func LengthCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		password := c.Query("password")
		username := c.Query("username")
		lp := len(password)
		ln := len([]rune(username))
		if ln == 0 || lp == 0 || ln > 32 || lp > 32 {
			code, msg := constant.Failed.GetInfo()
			c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
		}
		c.Next(ctx)
	}
}

func PublishCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		form, err := c.MultipartForm()
		file, err := c.FormFile("data")
		if err != nil || form == nil || len([]rune(form.Value["title"][0])) > 32 || file == nil {
			code, msg := constant.Failed.GetInfo()
			c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
		}
		open, err := file.Open()
		if err != nil {
			code, msg := constant.Failed.GetInfo()
			c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
		}
		data := make([]byte, file.Size)
		_, err = open.Read(data)
		if err != nil {
			code, msg := constant.Failed.GetInfo()
			c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
		}
		c.Set("data", data)
		c.Next(ctx)
	}
}
func CommentCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

	}
}
