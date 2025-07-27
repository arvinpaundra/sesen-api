package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/token"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
	"github.com/arvinpaundra/sesen-api/domain/auth/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/auth/dto/response"
	"github.com/arvinpaundra/sesen-api/domain/auth/entity"
	"github.com/arvinpaundra/sesen-api/domain/auth/repository"
)

type UserLogin struct {
	userReader repository.UserReader
	userWriter repository.UserWriter
	tokenable  token.Tokenable
	uow        repository.UnitOfWork
}

func NewUserLogin(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
	tokenable token.Tokenable,
	uow repository.UnitOfWork,
) *UserLogin {
	return &UserLogin{
		userReader: userReader,
		userWriter: userWriter,
		tokenable:  tokenable,
		uow:        uow,
	}
}

func (s *UserLogin) Execute(ctx context.Context, payload request.UserLogin) (*response.UserLogin, error) {
	user, err := s.userReader.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	err = util.CompareHashAndString(user.Password, payload.Password)
	if err != nil {
		return nil, constant.ErrUserNotFound
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

	tx, err := s.uow.Begin()
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

	result := response.UserLogin{
		UserId:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &result, nil
}
