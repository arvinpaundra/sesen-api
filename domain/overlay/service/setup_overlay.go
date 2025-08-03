package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/overlay/constant"
	"github.com/arvinpaundra/sesen-api/domain/overlay/dto"
	"github.com/arvinpaundra/sesen-api/domain/overlay/entity"
	"github.com/arvinpaundra/sesen-api/domain/overlay/repository"
)

type SetupOverlay struct {
	overlayReader repository.OverlayReader
	overlayWriter repository.OverlayWriter
	uow           repository.UnitOfWork
}

func NewSetupOverlay(
	overlayReader repository.OverlayReader,
	overlayWriter repository.OverlayWriter,
	uow repository.UnitOfWork,
) *SetupOverlay {
	return &SetupOverlay{
		overlayReader: overlayReader,
		overlayWriter: overlayWriter,
		uow:           uow,
	}
}

func (s *SetupOverlay) Execute(ctx context.Context, payload dto.SetupOverlay) error {
	exist, err := s.overlayReader.UserHasOverlay(ctx, payload.UserId)
	if err != nil {
		return err
	}

	if exist {
		return constant.ErrUserAlreadyHasOverlay
	}

	overlay := entity.NewOverlay(
		payload.UserId,
		constant.DefaultRingtoneUrl,
		constant.TTSDisabled,
		constant.NFSWDisabled,
	)

	qrcode, err := entity.NewOverlayQR(overlay.ID, overlay.UserId, constant.DefaultQRColor, constant.DefaultBackgroundColor)
	if err != nil {
		return err
	}

	overlay.SetQRCode(qrcode)

	message, err := entity.NewOverlayMessage(
		overlay.ID,
		constant.DefaultTextColor,
		constant.DefaultBackgroundColor,
	)

	if err != nil {
		return err
	}

	overlay.SetMessage(message)

	tx, err := s.uow.Begin()
	if err != nil {
		return err
	}

	err = tx.OverlayWriter().Save(ctx, overlay)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}

		return err
	}

	uowErr := tx.Commit()
	if uowErr != nil {
		return uowErr
	}

	return nil
}
