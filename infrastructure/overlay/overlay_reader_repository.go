package overlay

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/overlay/entity"
	"gorm.io/gorm"
)

type OverlayReaderRepository struct {
	db *gorm.DB
}

func NewOverlayReaderRepository(db *gorm.DB) *OverlayReaderRepository {
	return &OverlayReaderRepository{db: db}
}

func (r *OverlayReaderRepository) FindOverlayById(ctx context.Context, id string) (*entity.Overlay, error) {
	panic("not implemented")
}

func (r *OverlayReaderRepository) FindOverlayByUserId(ctx context.Context, userId string) (*entity.Overlay, error) {
	panic("not implemented")
}

func (r *OverlayReaderRepository) UserHasOverlay(ctx context.Context, userId string) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&entity.Overlay{}).
		Select("id").
		Where("user_id = ?", userId).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
