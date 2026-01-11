package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/widget/repository"
	"github.com/arvinpaundra/sesen-api/domain/widget/response"
)

type FindWidgetSettingCommand struct {
	UserID string
}

type FindWidgetSetting struct {
	widgetReader repository.WidgetReader
}

func NewFindWidgetSetting(widgetReader repository.WidgetReader) *FindWidgetSetting {
	return &FindWidgetSetting{
		widgetReader: widgetReader,
	}
}

func (s *FindWidgetSetting) Execute(ctx context.Context, command FindWidgetSettingCommand) (response.WidgetSetting, error) {
	settings, err := s.widgetReader.FindWidgetSettingsByUserID(ctx, command.UserID)
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
