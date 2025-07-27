package service

import (
	"context"
	"errors"

	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
	"github.com/arvinpaundra/sesen-api/domain/auth/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/auth/entity"
	"github.com/arvinpaundra/sesen-api/domain/auth/repository"
)

type UserRegister struct {
	userReader repository.UserReader
	userWriter repository.UserWriter
}

func NewUserRegister(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
) *UserRegister {
	return &UserRegister{
		userReader: userReader,
		userWriter: userWriter,
	}
}

func (s *UserRegister) Execute(ctx context.Context, payload request.UserRegister) error {
	existingUser, err := s.userReader.FindUserByEmail(ctx, payload.Email)
	if err != nil && !errors.Is(err, constant.ErrUserNotFound) {
		return err
	}

	if !existingUser.IsEmpty() {
		return constant.ErrEmailAlreadyExists
	}

	existingUser, err = s.userReader.FindUserByUsername(ctx, payload.Username)
	if err != nil && !errors.Is(err, constant.ErrUserNotFound) {
		return err
	}

	if !existingUser.IsEmpty() {
		return constant.ErrUsernameAlreadyExists
	}

	password, err := util.HashString(payload.Password)
	if err != nil {
		return err
	}

	user := entity.NewUser(
		payload.Email,
		password,
		payload.Username,
		payload.Fullname,
	)

	if err := s.userWriter.Save(ctx, user); err != nil {
		return err
	}

	return nil
}
