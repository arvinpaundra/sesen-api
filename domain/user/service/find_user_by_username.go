package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/user/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/user/dto/response"
	"github.com/arvinpaundra/sesen-api/domain/user/repository"
)

type FindUserByUsername struct {
	userReader repository.UserReader
}

func NewFindUserByUsername(
	userReader repository.UserReader,
) FindUserByUsername {
	return FindUserByUsername{
		userReader: userReader,
	}
}

func (s *FindUserByUsername) Execute(ctx context.Context, payload request.FindUserByUsernamePayload) (response.User, error) {
	user, err := s.userReader.FindUserByUsername(ctx, payload.Username)
	if err != nil {
		return response.User{}, err
	}

	return response.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.Fullname,
	}, nil
}
