package service

import (
	"context"
	"errors"

	"github.com/arvinpaundra/sesen-api/core/event"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
	"github.com/arvinpaundra/sesen-api/domain/auth/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/auth/entity"
	"github.com/arvinpaundra/sesen-api/domain/auth/repository"
)

type UserRegister struct {
	userReader repository.UserReader
	userWriter repository.UserWriter
	uow        repository.UnitOfWork
	publisher  event.DomainEventPublisher
}

func NewUserRegister(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
	uow repository.UnitOfWork,
	publisher event.DomainEventPublisher,
) *UserRegister {
	return &UserRegister{
		userReader: userReader,
		userWriter: userWriter,
		uow:        uow,
		publisher:  publisher,
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

	tx, err := s.uow.Begin()
	if err != nil {
		return err
	}

	// Save user and publish domain events to Asynq queue
	if err := tx.UserWriter().Save(ctx, user); err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	if user.HasDomainEvents() {
		events := user.GetDomainEvents()

		err := s.publisher.Publish(ctx, events...)
		if err != nil {
			if uowErr := tx.Rollback(); uowErr != nil {
				return uowErr
			}

			return err
		}
	}

	uowErr := tx.Commit()
	if uowErr != nil {
		return uowErr
	}

	return nil
}
