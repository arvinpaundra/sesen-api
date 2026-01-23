package adapter

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	"github.com/arvinpaundra/sesen-api/domain/widget/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/widget/service"
	"github.com/arvinpaundra/sesen-api/infrastructure/widget"
	"gorm.io/gorm"
)

var _ repository.WidgetMapper = (*WidgetAdapter)(nil)

type WidgetAdapter struct {
	db *gorm.DB
}

func NewWidgetAdapter(db *gorm.DB) *WidgetAdapter {
	return &WidgetAdapter{db: db}
}

func (w *WidgetAdapter) FindWidgetSettingsByUserID(ctx context.Context, userID string) (entity.WidgetSetting, error) {
	payload := request.FindWidgetSettingPayload{
		UserID: userID,
	}

	svc := service.NewFindWidgetSetting(widget.NewWidgetReaderRepository(w.db))

	settings, err := svc.Execute(ctx, payload)
	if err != nil {
		return entity.WidgetSetting{}, err
	}

	return entity.WidgetSetting{
		ID:        settings.ID,
		UserID:    settings.UserID,
		MinAmount: settings.MinAmount,
	}, nil
}
