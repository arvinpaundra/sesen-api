package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/auth/entity"
)

type UserReader interface {
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	FindUserWithActiveSessionsById(ctx context.Context, id string) (*entity.User, error)
}
