package model

import (
	"time"

	"github.com/google/uuid"
)

type Overlay struct {
	ID            uuid.UUID `gorm:"primaryKey;column:id"`
	UserId        uuid.UUID `gorm:"column:user_id"`
	RingtoneUrl   string    `gorm:"column:ringtone_url"`
	IsTTSEnabled  bool      `gorm:"column:is_tts_enabled"`
	IsNFSWEnabled bool      `gorm:"column:is_nsfw_enabled"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}
