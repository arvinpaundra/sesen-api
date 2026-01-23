package service

import (
	"context"
	"errors"

	"github.com/arvinpaundra/sesen-api/domain/user/constant"
	"github.com/arvinpaundra/sesen-api/domain/user/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/user/entity"
	"github.com/arvinpaundra/sesen-api/domain/user/repository"
)

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

func (s *UserRegister) Execute(ctx context.Context, payload request.UserRegisterPayload) error {
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

	user, err := entity.NewUser(
		payload.Email,
		payload.Username,
		payload.Password,
		payload.Fullname,
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
