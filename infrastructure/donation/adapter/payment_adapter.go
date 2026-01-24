package adapter

import (
	"context"

	"github.com/arvinpaundra/sesen-api/config"
	donationrequest "github.com/arvinpaundra/sesen-api/domain/donation/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	paymentrequest "github.com/arvinpaundra/sesen-api/domain/payment/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/payment/service"
	"github.com/arvinpaundra/sesen-api/infrastructure/payment"
	"github.com/xendit/xendit-go/v7"
	"gorm.io/gorm"
)

var _ repository.PaymentMapper = (*PaymentAdapter)(nil)

type PaymentAdapter struct {
	db     *gorm.DB
	client *xendit.APIClient
}

func NewPaymentAdapter(db *gorm.DB, client *xendit.APIClient) *PaymentAdapter {
	return &PaymentAdapter{
		db:     db,
		client: client,
	}
}

func (p *PaymentAdapter) Pay(ctx context.Context, payload donationrequest.CreatePayment) (string, error) {
	paymentPayload := paymentrequest.CreatePaymentPayload{
		UserID:      payload.UserID,
		ReferenceID: payload.ReferenceID,
		UserName:    payload.UserName,
		Amount:      payload.Amount,
		Method:      payload.Method.String(),
	}

	svc := service.NewCreatePayment(
		payment.NewXenditPaymentAdapter(config.GetString("XENDIT_API_KEY")),
		payment.NewUnitOfWork(p.db),
	)

	return svc.Execute(ctx, paymentPayload)
}
