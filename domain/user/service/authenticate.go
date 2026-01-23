package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/token"
	"github.com/arvinpaundra/sesen-api/domain/user/constant"
	"github.com/arvinpaundra/sesen-api/domain/user/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/user/dto/response"
	"github.com/arvinpaundra/sesen-api/domain/user/repository"
)

type CheckSession struct {
	userReader repository.UserReader
	tokenable  token.Tokenable
}

func NewCheckSession(
	userReader repository.UserReader,
	tokenable token.Tokenable,
) CheckSession {
	return CheckSession{
		userReader: userReader,
		tokenable:  tokenable,
	}
}

func (s *CheckSession) Execute(ctx context.Context, payload request.CheckSessionPayload) (*response.AuthenticatedUser, error) {
	claims, err := s.tokenable.Decode(payload.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userReader.FindUserById(ctx, claims.Identifier)
	if err != nil {
		return nil, err
	}

	if user.IsBanned() {
		return nil, constant.ErrUserHasBeenBanned
	}

	result := &response.AuthenticatedUser{
		UserId:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.Fullname,
	}

	return result, nil
}
