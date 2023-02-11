package mw

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/pkg/jwt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func JWTHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var token string
		form, _ := c.MultipartForm()
		if form != nil {
			v := form.Value["token"]
			if v != nil {
				token = v[0]
			} else {
				code, msg := constant.Failed.GetInfo()
				c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
			}
		} else {
			token = c.Query("token")
		}
		path := string(c.Path())
		if token != "" || (path != "/douyin/feed/" && path != "/douyin/user/login/" && path != "/douyin/user/register/") {
			parseToken, err := jwt.JWTUtil.ParseToken(token)
			if err != nil {
				code, msg := constant.Failed.GetInfo()
				c.AbortWithStatusJSON(http.StatusBadRequest, checkFailed{StatusCode: code, StatusMessage: msg})
			} else {
				c.Set("user_id", parseToken.UserId)
			}
		}
		c.Next(ctx)
	}
}

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
//	"github.com/PICOF/simple-tiktok/pkg/config"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/hertz-contrib/jwt"
//	"github.com/spf13/viper"
//)
//
//var (
//	JWTConfig      *viper.Viper
//	AuthMiddleware *jwt.HertzJWTMiddleware
//)
//
//const JWTConfigName = "jwt"
//const identityKey = "id"
//
//func init() {
//	initJWT()
//	JWTConfig = config.GetConfig(JWTConfigName)
//}
//
//func initJWT() {
//	var err error
//	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
//		Key:         []byte(JWTConfig.GetString("secretKey")),
//		Timeout:     JWTConfig.GetDuration("Timeout"),
//		MaxRefresh:  JWTConfig.GetDuration("maxRefresh"),
//		IdentityKey: identityKey,
//		PayloadFunc: func(data interface{}) jwt.MapClaims {
//			if v, ok := data.(int64); ok {
//				return jwt.MapClaims{
//					identityKey: v,
//				}
//			}
//			return jwt.MapClaims{}
//		},
//		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
//			claims := jwt.ExtractClaims(ctx, c)
//			userId, _ := claims[identityKey].(json.Number).Int64()
//			return &tiktokapi.UserInfo{
//				ID: userId,
//			}
//		},
//		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
//			var loginVals tiktokapi.LoginRequest
//			if err := c.BindAndValidate(&loginVals); err != nil {
//				return "", jwt.ErrMissingLoginValues
//			}
//			userID := loginVals.Username
//			password := loginVals.Password
//
//			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
//				return &User{
//					UserName:  userID,
//					LastName:  "Hertz",
//					FirstName: "CloudWeGo",
//				}, nil
//			}
//
//			return nil, jwt.ErrFailedAuthentication
//		},
//		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
//			if v, ok := data.(*User); ok && v.UserName == "admin" {
//				return true
//			}
//
//			return false
//		},
//		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
//			c.JSON(code, map[string]interface{}{
//				"code":    code,
//				"message": message,
//			})
//		},
//	})
//}
