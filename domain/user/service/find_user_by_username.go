package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/user/repository"
	"github.com/arvinpaundra/sesen-api/domain/user/response"
)

type FindUserByUsernameCommand struct {
	Username string
}

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

func (s *FindUserByUsername) Execute(ctx context.Context, command FindUserByUsernameCommand) (response.User, error) {
	user, err := s.userReader.FindUserByUsername(ctx, command.Username)
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
