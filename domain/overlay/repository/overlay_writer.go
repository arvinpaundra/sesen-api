package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/overlay/entity"
)

type OverlayWriter interface {
	Save(ctx context.Context, overlay *entity.Overlay) error
}
