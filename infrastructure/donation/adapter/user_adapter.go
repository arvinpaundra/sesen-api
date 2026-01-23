package adapter

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	"github.com/arvinpaundra/sesen-api/domain/user/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/user/service"
	infra "github.com/arvinpaundra/sesen-api/infrastructure/user"
	"gorm.io/gorm"
)

var _ repository.UserMapper = (*UserAdapter)(nil)

type UserAdapter struct {
	db *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{db: db}
}

func (u *UserAdapter) FindUserByUsername(ctx context.Context, username string) (entity.User, error) {
	payload := request.FindUserByUsernamePayload{
		Username: username,
	}

	svc := service.NewFindUserByUsername(infra.NewUserReaderRepository(u.db))

	user, err := svc.Execute(ctx, payload)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
