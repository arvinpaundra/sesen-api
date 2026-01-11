package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleStreamer UserRole = "streamer"
)

type UserStatus string

const (
	StatusActive UserStatus = "active"
	StatusBanned UserStatus = "banned"
)

type User struct {
	ID        uuid.UUID  `gorm:"primaryKey;column:id"`
	Email     string     `gorm:"column:email"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Fullname  string     `gorm:"column:fullname"`
	Role      UserRole   `gorm:"type:user_role;column:role"`
	Status    UserStatus `gorm:"type:user_status;column:status"`
	Token     string     `gorm:"column:token;unique"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`

	Sessions []*Session `gorm:"foreignKey:UserId;references:ID"`
}
