package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/widget/entity"
)

type WidgetReader interface {
	HasWidgetSettings(ctx context.Context, userID string) (bool, error)
	FindWidgetSettingsByUserID(ctx context.Context, userID string) (*entity.WidgetSetting, error)
}
