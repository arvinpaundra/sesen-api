package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
)

type DonationWriter interface {
	Save(ctx context.Context, donation *entity.Donation) error
}
