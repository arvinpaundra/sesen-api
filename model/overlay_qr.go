package model

import (
	"time"

	"github.com/google/uuid"
)

type OverlayQR struct {
	ID              uuid.UUID `gorm:"primaryKey;column:id"`
	OverlayId       uuid.UUID `gorm:"column:overlay_id"`
	Code            string    `gorm:"column:code"`
	QrColor         string    `gorm:"column:qr_color"`
	BackgroundColor string    `gorm:"column:background_color"`
	CreatedAt       time.Time `gorm:"column:created_at"`
}

func (OverlayQR) TableName() string {
	return "overlay_qr"
}
