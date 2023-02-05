package jwt

import (
	"fmt"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"log"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := JWTUtil.CreateToken(tiktokapi.UserInfo{ID: 123, Name: "", FollowCount: 1, FollowerCount: 1, IsFollow: false})
	if err != nil {
		log.Println("Error", err)
		return
	}
	parseToken, err := JWTUtil.ParseToken(token)
	if err != nil {
		log.Println("Error", err)
		return
	}
	fmt.Println(*parseToken)
}
