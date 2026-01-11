package widget

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/widget/entity"
	"github.com/arvinpaundra/sesen-api/domain/widget/repository"
	"gorm.io/gorm"
)

var _ repository.WidgetReader = (*WidgetReaderRepository)(nil)

type WidgetReaderRepository struct {
	db *gorm.DB
}

func NewWidgetReaderRepository(db *gorm.DB) *WidgetReaderRepository {
	return &WidgetReaderRepository{db: db}
}

func (r *WidgetReaderRepository) HasWidgetSettings(ctx context.Context, userID string) (bool, error) {
	var isExists bool

	err := r.db.WithContext(ctx).
		Raw("SELECT EXISTS(SELECT 1 FROM widget_settings WHERE user_id = ?)", userID).
		Scan(&isExists).Error

	if err != nil {
		return false, err
	}

	return isExists, nil
}

func (r *WidgetReaderRepository) FindWidgetSettingsByUserID(ctx context.Context, userID string) (*entity.WidgetSetting, error) {
	panic("not implemented")
}
