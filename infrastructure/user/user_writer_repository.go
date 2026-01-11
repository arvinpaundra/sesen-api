package user

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/user/entity"
	"github.com/arvinpaundra/sesen-api/domain/user/repository"
	"github.com/arvinpaundra/sesen-api/model"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

var _ repository.UserWriter = (*UserWriterRepository)(nil)

type UserWriterRepository struct {
	db *gorm.DB
}

func NewUserWriterRepository(db *gorm.DB) *UserWriterRepository {
	return &UserWriterRepository{db: db}
}

func (r *UserWriterRepository) Save(ctx context.Context, user *entity.User) error {
	if user.IsUpdated() {
		return r.update(ctx, user)
	}

	return r.insert(ctx, user)
}

func (r *UserWriterRepository) insert(ctx context.Context, user *entity.User) error {
	userModel := model.User{
		ID:       util.ParseUUID(user.ID),
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		Fullname: user.Fullname,
		Role:     model.UserRole(user.Role),
		Status:   model.UserStatus(user.Status),
		Token:    user.Token,
	}

	if err := r.db.WithContext(ctx).Create(&userModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserWriterRepository) update(ctx context.Context, user *entity.User) error {
	userModel := model.User{
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Status:   model.UserStatus(user.Status),
	}

	if err := r.db.WithContext(ctx).Where("id = ?", user.ID).Updates(&userModel).Error; err != nil {
		return err
	}

	for _, session := range user.ActiveSessions {
		sessionModel := model.Session{
			ID:           util.ParseUUID(session.ID),
			UserId:       util.ParseUUID(user.ID),
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			RevokedAt:    null.TimeFromPtr(session.RevokedAt),
		}

		if session.IsCreated() {
			if err := r.db.WithContext(ctx).Create(&sessionModel).Error; err != nil {
				return err
			}
		} else if session.IsUpdated() {
			if err := r.db.WithContext(ctx).Where("id = ?", session.ID).Updates(&sessionModel).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
