package jwt

import (
	"fmt"
	"log"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := JWTUtil.CreateToken(114514)
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
