package model

import (
	"time"

	"github.com/google/uuid"
)

type OverlayMessage struct {
	ID              uuid.UUID `gorm:"primaryKey;column:id"`
	OverlayId       uuid.UUID `gorm:"column:overlay_id"`
	TextColor       string    `gorm:"column:text_color"`
	BackgroundColor string    `gorm:"column:background_color"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
