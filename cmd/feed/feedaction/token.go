package feedaction

import "github.com/PICOF/simple-tiktok/pkg/jwt"

func ParseToken(token string) (int64, error) {
	parseToken, err := jwt.JWTUtil.ParseToken(token)
	if err != nil {
		return 0, err
	}
	return parseToken.UserId, nil
}
