package widget

import (
	"context"
	"errors"

	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
	"github.com/arvinpaundra/sesen-api/domain/widget/entity"
	"github.com/arvinpaundra/sesen-api/domain/widget/repository"
	"github.com/arvinpaundra/sesen-api/model"
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
	var settingsModel model.WidgetSetting

	err := r.db.WithContext(ctx).
		Model(&model.WidgetSetting{}).
		Where("user_id = ?", userID).
		First(&settingsModel).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrWidgetSettingsNotFound
		}
		return nil, err
	}

	return &entity.WidgetSetting{
		ID:              settingsModel.ID.String(),
		UserID:          settingsModel.UserID.String(),
		TTSEnabled:      settingsModel.TTSEnabled,
		NSFWFilter:      settingsModel.NSFWFilter,
		MessageDuration: settingsModel.MessageDuration,
		MinAmount:       settingsModel.MinAmount,
	}, nil
}
