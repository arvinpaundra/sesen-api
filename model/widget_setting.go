package model

import (
	"time"

	"github.com/google/uuid"
)

type WidgetSetting struct {
	ID              uuid.UUID `gorm:"primaryKey;column:id"`
	UserID          uuid.UUID `gorm:"column:user_id;unique"`
	TTSEnabled      bool      `gorm:"column:tts_enabled"`
	NSFWFilter      bool      `gorm:"column:nsfw_filter"`
	MessageDuration int       `gorm:"column:message_duration"`
	MinAmount       int64     `gorm:"column:min_amount"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`

	QrCode *WidgetQrcode `gorm:"foreignKey:SettingID;references:ID"`
	Alert  *WidgetAlert  `gorm:"foreignKey:SettingID;references:ID"`
}
