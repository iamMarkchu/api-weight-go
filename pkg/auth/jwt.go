package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type ApiClaims struct {
	UserId uint64 `json:"user_id"`
	jwtv4.RegisteredClaims
}

func GetUserIdFromCtx(ctx context.Context) (userId uint64, ok bool) {
	cc, ok := jwt.FromContext(ctx)
	if !ok {
		return
	}
	ccm, ok := cc.(jwtv4.MapClaims)
	if !ok {
		return
	}
	ccc, ok := ccm["user_id"]
	if !ok {
		return
	}
	cccc, ok := ccc.(float64)
	if !ok {
		return
	}
	return uint64(cccc), ok
}
