package repository

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/payment/dto/request"
)

type PaymentGatewayMapper interface {
	Pay(ctx context.Context, payload request.PaymentRequestPayload) (string, error)
}
