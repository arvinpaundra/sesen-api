package model

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v6"
	"gorm.io/datatypes"
)

type WebhookEvent struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Provider  null.String    `gorm:"column:provider;nullable"`
	EventID   null.String    `gorm:"column:event_id;nullable"`
	Payload   datatypes.JSON `gorm:"column:payload;type:jsonb;nullable"`
	CreatedAt null.Time      `gorm:"column:created_at;nullable"`
}
