package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
)

type User struct {
	trait.Createable
	trait.Updateable

	ID       string
	Email    string
	Username string
	Password string
	Fullname string
	Role     constant.UserRole
	Status   constant.UserStatus

	ActiveSessions []*Session
}

func NewUser(email, username, password, fullname string) *User {
	user := &User{
		ID:       util.GenerateUUID(),
		Email:    email,
		Username: username,
		Password: password,
		Fullname: fullname,
		Role:     constant.RoleStreamer,
		Status:   constant.StatusActive,
	}

	user.MarkCreate()

	return user
}

func (u *User) IsEmpty() bool {
	return u == nil
}

func (u *User) IsIdentifierEqual(id string) bool {
	return u.ID == id
}

func (u *User) AddActiveSession(session *Session) {
	if session == nil {
		return
	}

	u.ActiveSessions = append(u.ActiveSessions, session)
}

func (u *User) RevokeSessionByAccessToken(accessToken string) error {
	for _, session := range u.ActiveSessions {
		if session.IsAccessTokenEqual(accessToken) {
			session.Revoke()
			return nil
		}
	}

	return constant.ErrInvalidAccessToken
}

func (u *User) RevokeSessionByRefreshToken(refreshToken string) error {
	for _, session := range u.ActiveSessions {
		if session.IsRefreshTokenEqual(refreshToken) {
			session.Revoke()
			return nil
		}
	}

	return constant.ErrInvalidRefreshToken
}

func (u *User) IsBanned() bool {
	return u.Status == constant.StatusBanned
}
