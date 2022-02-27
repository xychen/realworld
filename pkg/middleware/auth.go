package middleware

import (
	"errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"realworld/pkg/jwt"
	"strings"
)
import "context"

var currentUserKey struct{}

type CurrentUser struct {
	Username string
}

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				//do something on entering
				// Token xxxxx
				tokenStr := tr.RequestHeader().Get("Authorization")
				auths := strings.SplitN(tokenStr, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
					return nil, errors.New("jwt token missing")
				}
				// 获取user info
				userinfo, err := jwt.ParseToken(auths[1])
				if err != nil {
					//	鉴权失败
					return nil, err
				}
				ctx = WithContext(ctx, &CurrentUser{Username: userinfo.UserName})
				defer func() {
					//	do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}

func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}
