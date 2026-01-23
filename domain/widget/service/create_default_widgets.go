package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
	"github.com/arvinpaundra/sesen-api/domain/widget/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/widget/entity"
	"github.com/arvinpaundra/sesen-api/domain/widget/repository"
)

type CreateDefaultWidgets struct {
	widgetReader repository.WidgetReader
	widgetWriter repository.WidgetWriter
	uow          repository.UnitOfWork
}

func NewCreateDefaultWidgets(
	widgetReader repository.WidgetReader,
	widgetWriter repository.WidgetWriter,
	uow repository.UnitOfWork,
) CreateDefaultWidgets {
	return CreateDefaultWidgets{
		widgetReader: widgetReader,
		widgetWriter: widgetWriter,
		uow:          uow,
	}
}

func (s *CreateDefaultWidgets) Execute(ctx context.Context, payload request.CreateDefaultWidgetsPayload) error {
	// check if user already has widget settings
	hasSettings, err := s.widgetReader.HasWidgetSettings(ctx, payload.UserID)
	if err != nil {
		return err
	}

	if hasSettings {
		return constant.ErrUserAlreadyHasWidgetSettings
	}

	// create default widget settings
	settings := entity.NewWidgetSetting(payload.UserID)

	// create default widget qr code
	qrcode := entity.NewWidgetQrcode(settings.ID, payload.Username)

	settings.AddQrCode(qrcode)

	// create default widget alert
	alert := entity.NewWidgetAlert(settings.ID)

	settings.AddAlert(alert)

	uow, err := s.uow.Begin(ctx)
	if err != nil {
		return err
	}

	// save widget settings
	err = uow.WidgetWriter().Save(ctx, settings)
	if err != nil {
		if uowErr := uow.Rollback(); uowErr != nil {
			return uowErr
		}
		return err
	}

	err = uow.Commit()
	if err != nil {
		return err
	}

	return nil
}
