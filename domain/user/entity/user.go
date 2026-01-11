package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/user/constant"
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
	Token    string

	ActiveSessions []*Session
}

func NewUser(email, username, password, fullname string) (*User, error) {
	user := &User{
		ID:       util.GenerateUUID(),
		Email:    email,
		Username: username,
		Fullname: fullname,
		Role:     constant.RoleStreamer,
		Status:   constant.StatusActive,
	}

	err := user.GenToken()
	if err != nil {
		return nil, err
	}

	user.Password, err = util.HashString(password)
	if err != nil {
		return nil, err
	}

	user.MarkCreate()

	return user, nil
}

func (u *User) IsEmpty() bool {
	return u == nil
}

func (u *User) GenToken() error {
	apiKey, err := util.RandomAlphanumeric(32)
	if err != nil {
		return err
	}

	u.Token = apiKey

	return nil
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
