package overlay

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/overlay/entity"
	"github.com/arvinpaundra/sesen-api/model"
	"gorm.io/gorm"
)

type OverlayWriterRepository struct {
	db *gorm.DB
}

func NewOverlayWriterRepository(db *gorm.DB) *OverlayWriterRepository {
	return &OverlayWriterRepository{db: db}
}

func (r *OverlayWriterRepository) Save(ctx context.Context, overlay *entity.Overlay) error {
	if overlay.IsCreated() {
		return r.insert(ctx, overlay)
	}

	return nil
}

func (r *OverlayWriterRepository) insert(ctx context.Context, overlay *entity.Overlay) error {
	overlayModel := model.Overlay{
		ID:            util.ParseUUID(overlay.ID),
		UserId:        util.ParseUUID(overlay.UserId),
		RingtoneUrl:   overlay.RingtoneUrl,
		IsTTSEnabled:  overlay.IsTTSEnabled,
		IsNFSWEnabled: overlay.IsNFSWEnabled,
	}

	if err := r.db.WithContext(ctx).Create(&overlayModel).Error; err != nil {
		return err
	}

	if overlay.HasQRCode() {
		qr := overlay.QRCode

		qrModel := model.OverlayQR{
			ID:              util.ParseUUID(qr.ID),
			OverlayId:       util.ParseUUID(overlay.ID),
			Code:            qr.Code,
			QrColor:         qr.QrColor.String(),
			BackgroundColor: qr.BackgroundColor.String(),
		}

		if err := r.db.WithContext(ctx).Create(&qrModel).Error; err != nil {
			return err
		}
	}

	if overlay.HasMessage() {
		message := overlay.Message

		messageModel := model.OverlayMessage{
			ID:              util.ParseUUID(message.ID),
			OverlayId:       util.ParseUUID(overlay.ID),
			TextColor:       message.TextColor.String(),
			BackgroundColor: message.BackgroundColor.String(),
		}

		if err := r.db.WithContext(ctx).Create(&messageModel).Error; err != nil {
			return err
		}
	}

	return nil
}
