package jwt

import (
	"errors"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/spf13/viper"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const ConfigName = "jwt"

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	JWTConfig        *viper.Viper
	expire           time.Duration
	JWTUtil          *JWT
)

type JWT struct {
	SignKey []byte
}

type CustomClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

func init() {
	JWTConfig = config.GetConfig(ConfigName)
	expire = JWTConfig.GetDuration("timeout")
	JWTUtil = NewJWT()
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(JWTConfig.GetString("secretKey")),
	}
}

// 创建 token
func (j *JWT) CreateToken(userId int64) (string, error) {
	startTime := time.Now().Add(-time.Second)
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(startTime),             // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(startTime.Add(expire)), // 过期时间 7天  配置文件
			Issuer:    JWTConfig.GetString("issuer"),             // 签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})
	if err != nil {
		// if ve, ok := err.(*jwt.ValidationError); ok {
		// 	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
		// 		return nil, TokenMalformed
		// 	} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
		// 		// Token is expired
		// 		return nil, TokenExpired
		// 	} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
		// 		return nil, TokenNotValidYet
		// 	} else {
		// 		return nil, TokenInvalid
		// 	}
		// }
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
