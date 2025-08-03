package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/overlay/entity"
)

type OverlayReader interface {
	FindOverlayById(ctx context.Context, id string) (*entity.Overlay, error)
	FindOverlayByUserId(ctx context.Context, userId string) (*entity.Overlay, error)
	UserHasOverlay(ctx context.Context, userId string) (bool, error)
}
