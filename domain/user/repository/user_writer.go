package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/user/entity"
)

type UserWriter interface {
	Save(ctx context.Context, user *entity.User) error
}
