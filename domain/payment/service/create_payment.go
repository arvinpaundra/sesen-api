package service

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/payment/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/payment/repository"
)

type CreatePayment struct {
	paymentGatewayMapper repository.PaymentGatewayMapper
	uow                  repository.UnitOfWork
}

func NewCreatePayment(
	paymentGatewayMapper repository.PaymentGatewayMapper,
	uow repository.UnitOfWork,
) CreatePayment {
	return CreatePayment{
		paymentGatewayMapper: paymentGatewayMapper,
		uow:                  uow,
	}
}

func (s *CreatePayment) Execute(ctx context.Context, payload request.CreatePaymentPayload) (string, error) {
	// Convert CreatePaymentPayload to PaymentRequestPayload
	paymentRequest := request.PaymentRequestPayload{
		ReferenceID:        payload.ReferenceID,
		Amount:             int64(payload.Amount),
		Method:             payload.Method,
		CustomerName:       "", // TODO: Get from payload if needed
		SuccessRedirectURL: "", // TODO: Get from config or payload
		FailureRedirectURL: "", // TODO: Get from config or payload
		CallbackURL:        "", // TODO: Get from config or payload
	}

	token, err := s.paymentGatewayMapper.Pay(ctx, paymentRequest)
	if err != nil {
		return "", err
	}

	return token, nil
}
