package shared

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/shared/entity"
)

// TxKey is the shared context key for storing database transactions across bounded contexts.
type TxKey struct{}

// NewAuthStorage retrieves UserAuth from context, returns nil if not found
func NewAuthStorage(ctx context.Context) *entity.UserAuth {
	auth, ok := ctx.Value("session").(*entity.UserAuth)
	if !ok {
		return nil
	}

	return auth
}
