package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/widget/entity"
)

type WidgetWriter interface {
	Save(ctx context.Context, settings *entity.WidgetSetting) error
}
