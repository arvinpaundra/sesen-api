package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/token"
	"github.com/arvinpaundra/sesen-api/domain/user/constant"
	"github.com/arvinpaundra/sesen-api/domain/user/entity"
	"github.com/arvinpaundra/sesen-api/domain/user/repository"
	"github.com/arvinpaundra/sesen-api/domain/user/response"
)

type RefreshTokenCommand struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshToken struct {
	userReader repository.UserReader
	userWriter repository.UserWriter
	tokenable  token.Tokenable
	uow        repository.UnitOfWork
}

func NewRefreshToken(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
	tokenable token.Tokenable,
	uow repository.UnitOfWork,
) *RefreshToken {
	return &RefreshToken{
		userReader: userReader,
		userWriter: userWriter,
		tokenable:  tokenable,
		uow:        uow,
	}
}

func (s *RefreshToken) Execute(ctx context.Context, command RefreshTokenCommand) (*response.RefreshToken, error) {
	claims, err := s.tokenable.Decode(command.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userReader.FindUserWithActiveSessionsById(ctx, claims.Identifier)
	if err != nil {
		return nil, err
	}

	err = user.RevokeSessionByRefreshToken(command.RefreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenable.Encode(user.ID, constant.TokenValidThreeHours, constant.TokenValidImmediately)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenable.Encode(user.ID, constant.TokenValidSevenDays, constant.TokenValidAfterThreeHours)
	if err != nil {
		return nil, err
	}

	session := entity.NewSession(user.ID, accessToken, refreshToken)

	user.AddActiveSession(session)

	user.MarkUpdate()

	tx, err := s.uow.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.UserWriter().Save(ctx, user)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return nil, uowErr
		}

		return nil, err
	}

	uowErr := tx.Commit()
	if uowErr != nil {
		return nil, uowErr
	}

	result := &response.RefreshToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       user.ID,
	}

	return result, nil
}
