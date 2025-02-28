package utils

import (
	"context"
)

func GetUserIDFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserId)
	if userId == nil {
		return 0
	}
	return userId.(int32)
}
