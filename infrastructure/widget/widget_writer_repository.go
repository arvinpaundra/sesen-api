package widget

import (
	"context"
	"encoding/json"

	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/widget/entity"
	"github.com/arvinpaundra/sesen-api/domain/widget/repository"
	"github.com/arvinpaundra/sesen-api/model"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

var _ repository.WidgetWriter = (*WidgetWriterRepository)(nil)

type WidgetWriterRepository struct {
	db *gorm.DB
}

func NewWidgetWriterRepository(db *gorm.DB) *WidgetWriterRepository {
	return &WidgetWriterRepository{db: db}
}

func (r *WidgetWriterRepository) Save(ctx context.Context, settings *entity.WidgetSetting) error {
	if settings.IsUpdated() {
		return r.update(ctx, settings)
	}

	return r.insert(ctx, settings)
}

func (r *WidgetWriterRepository) insert(ctx context.Context, settings *entity.WidgetSetting) error {
	settingModel := model.WidgetSetting{
		ID:                    util.ParseUUID(settings.ID),
		UserID:                util.ParseUUID(settings.UserID),
		TTSEnabled:            settings.TTSEnabled,
		NSFWFilter:            settings.NSFWFilter,
		MessageDuration:       settings.MessageDuration,
		MinimumDonationAmount: settings.MinimumDonationAmount,
	}

	if err := r.db.WithContext(ctx).Create(&settingModel).Error; err != nil {
		return err
	}

	if settings.HasQrCode() {
		if err := r.insertQrCode(ctx, settings.QrCode); err != nil {
			return err
		}
	}

	if settings.HasAlert() {
		if err := r.insertAlert(ctx, settings.Alert); err != nil {
			return err
		}
	}

	return nil
}

func (r *WidgetWriterRepository) insertQrCode(ctx context.Context, qrcode *entity.WidgetQrcode) error {
	styles, err := json.Marshal(qrcode.Styles.Map())
	if err != nil {
		return err
	}

	qrcodeModel := model.WidgetQrcode{
		ID:          util.ParseUUID(qrcode.ID),
		SettingID:   util.ParseUUID(qrcode.SettingID),
		QrCodeData:  qrcode.QrCodeData,
		Description: qrcode.Description,
		Styles:      styles,
	}

	if err := r.db.WithContext(ctx).Create(&qrcodeModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *WidgetWriterRepository) insertAlert(ctx context.Context, alert *entity.WidgetAlert) error {
	styles, err := json.Marshal(alert.Styles.Map())
	if err != nil {
		return err
	}

	alertModel := model.WidgetAlert{
		ID:            util.ParseUUID(alert.ID),
		SettingID:     util.ParseUUID(alert.SettingID),
		AlertText:     alert.AlertText,
		AlertURL:      alert.AlertURL,
		AttachmentURL: null.StringFromPtr(alert.AttachmentURL),
		Styles:        styles,
	}

	if err := r.db.WithContext(ctx).Create(&alertModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *WidgetWriterRepository) update(_ context.Context, _ *entity.WidgetSetting) error {
	return nil
}
