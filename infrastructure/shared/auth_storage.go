package shared

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/shared/entity"
)

func NewAuthStorage(ctx context.Context) *entity.UserAuth {
	auth, ok := ctx.Value("session").(*entity.UserAuth)
	if !ok {
		return nil
	}

	return auth
}
