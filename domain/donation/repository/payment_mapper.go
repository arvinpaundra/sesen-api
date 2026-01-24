package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/dto/request"
)

type PaymentMapper interface {
	Pay(ctx context.Context, payload request.CreatePayment) (string, error)
}
