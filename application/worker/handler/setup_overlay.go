package handler

import (
	"context"
	"encoding/json"

	"github.com/arvinpaundra/sesen-api/domain/overlay/dto"
	"github.com/arvinpaundra/sesen-api/domain/overlay/service"
	"github.com/arvinpaundra/sesen-api/infrastructure/overlay"
	"gorm.io/gorm"
)

type SetupOverlayHandler struct {
	db *gorm.DB
}

func NewSetupOverlayHandler(db *gorm.DB) *SetupOverlayHandler {
	return &SetupOverlayHandler{
		db: db,
	}
}

func (h *SetupOverlayHandler) Handle(ctx context.Context, eventData []byte) error {
	var payload dto.SetupOverlay

	err := json.Unmarshal(eventData, &payload)
	if err != nil {
		return err
	}

	svc := service.NewSetupOverlay(
		overlay.NewOverlayReaderRepository(h.db),
		overlay.NewOverlayWriterRepository(h.db),
		overlay.NewUnitOfWork(h.db),
	)

	err = svc.Execute(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
