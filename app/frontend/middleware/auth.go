package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIDKey string

const SessionUserId SessionUserIDKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

// 鉴权逻辑
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userID := s.Get("user_id")
		if userID == nil {
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
