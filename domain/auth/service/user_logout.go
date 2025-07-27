package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/token"
	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
	"github.com/arvinpaundra/sesen-api/domain/auth/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/auth/repository"
	"github.com/arvinpaundra/sesen-api/domain/shared/interfaces"
)

type UserLogout struct {
	userReader repository.UserReader
	userWriter repository.UserWriter
	token      token.Tokenable
	auth       interfaces.AuthenticatedUser
	uow        repository.UnitOfWork
}

func NewUserLogout(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
	tokenable token.Tokenable,
	auth interfaces.AuthenticatedUser,
	uow repository.UnitOfWork,
) *UserLogout {
	return &UserLogout{
		userReader: userReader,
		userWriter: userWriter,
		token:      tokenable,
		auth:       auth,
		uow:        uow,
	}
}

func (s *UserLogout) Execute(ctx context.Context, payload request.UserLogout) error {
	claims, err := s.token.Decode(payload.AccessToken)
	if err != nil {
		return err
	}

	user, err := s.userReader.FindUserWithActiveSessionsById(ctx, claims.Identifier)
	if err != nil {
		return err
	}

	if !user.IsIdentifierEqual(s.auth.GetUserId()) {
		return constant.ErrInvalidAccessToken
	}

	err = user.RevokeSessionByAccessToken(payload.AccessToken)
	if err != nil {
		return err
	}

	user.MarkUpdate()

	tx, err := s.uow.Begin()
	if err != nil {
		return err
	}

	err = tx.UserWriter().Save(ctx, user)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	uowErr := tx.Commit()
	if uowErr != nil {
		return uowErr
	}

	return nil
}
