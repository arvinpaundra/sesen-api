package donation

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	"github.com/arvinpaundra/sesen-api/model"
	"github.com/guregu/null/v6"
	"gorm.io/gorm"
)

var _ repository.DonationWriter = (*DonationWriter)(nil)

type DonationWriter struct {
	db *gorm.DB
}

func NewDonationWriter(db *gorm.DB) *DonationWriter {
	return &DonationWriter{db: db}
}

func (w *DonationWriter) Save(ctx context.Context, donation *entity.Donation) error {
	if donation.IsUpdated() {
		return w.update(ctx, donation)
	}

	return w.insert(ctx, donation)
}

func (w *DonationWriter) insert(ctx context.Context, donation *entity.Donation) error {
	donationModel := model.Donation{
		ID:            util.ParseUUID(donation.ID),
		ToUserId:      util.ParseUUID(donation.ToUserID),
		Amount:        donation.Amount.Number(),
		Currency:      donation.Currency,
		PaymentMethod: model.PaymentMethod(donation.PaymentMethod),
		DonorName:     null.StringFromPtr(donation.DonorName),
		Message:       null.StringFromPtr(donation.Message),
		Status:        model.DonationStatus(donation.Status),
	}

	if err := w.db.WithContext(ctx).Create(&donationModel).Error; err != nil {
		return err
	}

	return nil
}

func (w *DonationWriter) update(_ context.Context, _ *entity.Donation) error {
	return nil
}
