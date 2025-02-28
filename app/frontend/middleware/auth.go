package middleware

import (
	"context"

	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

// 鉴权逻辑
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userID := s.Get("user_id")

		if userID == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
