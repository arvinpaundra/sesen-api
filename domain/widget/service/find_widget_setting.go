package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/widget/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/widget/dto/response"
	"github.com/arvinpaundra/sesen-api/domain/widget/repository"
)

type FindWidgetSetting struct {
	widgetReader repository.WidgetReader
}

func NewFindWidgetSetting(widgetReader repository.WidgetReader) *FindWidgetSetting {
	return &FindWidgetSetting{
		widgetReader: widgetReader,
	}
}

func (s *FindWidgetSetting) Execute(ctx context.Context, payload request.FindWidgetSettingPayload) (response.WidgetSetting, error) {
	settings, err := s.widgetReader.FindWidgetSettingsByUserID(ctx, payload.UserID)
	if err != nil {
		return response.WidgetSetting{}, err
	}

	return response.WidgetSetting{
		ID:              settings.ID,
		UserID:          settings.UserID,
		TTSEnabled:      settings.TTSEnabled,
		NSFWFilter:      settings.NSFWFilter,
		MessageDuration: settings.MessageDuration,
		MinAmount:       settings.MinAmount,
	}, nil
}
