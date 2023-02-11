package mw

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/pkg/bcrypt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func BcryptMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		encrypted, err := bcrypt.EncryptPassword(ctx, c.Query("password"))
		if err != nil {
			code, msg := constant.Failed.GetInfo()
			c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
		}
		c.Set("password", string(encrypted))
		c.Next(ctx)
	}
}
