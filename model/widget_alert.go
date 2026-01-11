package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
	"gorm.io/datatypes"
)

type WidgetAlert struct {
	ID            uuid.UUID      `gorm:"primaryKey;column:id"`
	SettingID     uuid.UUID      `gorm:"column:setting_id;unique"`
	AlertText     string         `gorm:"column:alert_text"`
	AlertURL      string         `gorm:"column:alert_url"`
	AttachmentURL null.String    `gorm:"column:attachment_url;nullable"`
	Styles        datatypes.JSON `gorm:"type:jsonb;column:styles"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
}
