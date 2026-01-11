package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
)

type WidgetMapper interface {
	FindWidgetSettingsByUserID(ctx context.Context, userID string) (entity.WidgetSetting, error)
}
