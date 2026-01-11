package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
)

type UserMapper interface {
	FindUserByUsername(ctx context.Context, username string) (entity.User, error)
}
