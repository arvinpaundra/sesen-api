package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WidgetQrcode struct {
	ID          uuid.UUID      `gorm:"primaryKey;column:id"`
	SettingID   uuid.UUID      `gorm:"column:setting_id;unique"`
	QrCodeData  string         `gorm:"column:qr_code_data"`
	Description string         `gorm:"column:description"`
	Styles      datatypes.JSON `gorm:"type:jsonb;column:styles"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
}

func (wq WidgetQrcode) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{wq},
	}
}
