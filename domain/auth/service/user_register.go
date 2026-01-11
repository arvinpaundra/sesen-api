package service

import (
	"context"
	"errors"

	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
	"github.com/arvinpaundra/sesen-api/domain/auth/entity"
	"github.com/arvinpaundra/sesen-api/domain/auth/repository"
)

type UserRegisterCommand struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8"`
	Fullname string `json:"fullname" validate:"required,min=3,max=100"`
}

type UserRegister struct {
	userReader   repository.UserReader
	userWriter   repository.UserWriter
	uow          repository.UnitOfWork
	widgetMapper repository.WidgetMapper
}

func NewUserRegister(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
	widgetMapper repository.WidgetMapper,
	uow repository.UnitOfWork,
) *UserRegister {
	return &UserRegister{
		userReader:   userReader,
		userWriter:   userWriter,
		widgetMapper: widgetMapper,
		uow:          uow,
	}
}

func (s *UserRegister) Execute(ctx context.Context, command UserRegisterCommand) error {
	existingUser, err := s.userReader.FindUserByEmail(ctx, command.Email)
	if err != nil && !errors.Is(err, constant.ErrUserNotFound) {
		return err
	}

	if !existingUser.IsEmpty() {
		return constant.ErrEmailAlreadyExists
	}

	existingUser, err = s.userReader.FindUserByUsername(ctx, command.Username)
	if err != nil && !errors.Is(err, constant.ErrUserNotFound) {
		return err
	}

	if !existingUser.IsEmpty() {
		return constant.ErrUsernameAlreadyExists
	}

	user, err := entity.NewUser(
		command.Email,
		command.Username,
		command.Password,
		command.Fullname,
	)
	if err != nil {
		return err
	}

	uow, err := s.uow.Begin(ctx)
	if err != nil {
		return err
	}

	err = uow.UserWriter().Save(ctx, user)
	if err != nil {
		if uowErr := uow.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	// Pass transaction context to widget domain
	err = s.widgetMapper.CreateDefaultWidgets(uow.Context(), user.ID, user.Username)
	if err != nil {
		if uowErr := uow.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	uowErr := uow.Commit()
	if uowErr != nil {
		return uowErr
	}

	return nil
}
