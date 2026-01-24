package service

import (
	"context"
	"fmt"

	"github.com/arvinpaundra/sesen-api/domain/donation/constant"
	"github.com/arvinpaundra/sesen-api/domain/donation/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	"github.com/arvinpaundra/sesen-api/domain/shared/valueobject"
)

type CreateDonation struct {
	userMapper    repository.UserMapper
	widgetMapper  repository.WidgetMapper
	paymentMapper repository.PaymentMapper
	uow           repository.UnitOfWork
}

func NewCreateDonation(
	userMapper repository.UserMapper,
	widgetMapper repository.WidgetMapper,
	paymentMapper repository.PaymentMapper,
	uow repository.UnitOfWork,
) *CreateDonation {
	return &CreateDonation{
		userMapper:    userMapper,
		widgetMapper:  widgetMapper,
		paymentMapper: paymentMapper,
		uow:           uow,
	}
}

func (s *CreateDonation) Execute(ctx context.Context, payload request.CreateDonationPayload) error {
	// find donated user by username
	user, err := s.userMapper.FindUserByUsername(ctx, payload.Username)
	if err != nil {
		return err
	}

	// validate amount against user's minimum donation amount
	settings, err := s.widgetMapper.FindWidgetSettingsByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	if !payload.IsAmountValid(settings.MinAmount) {
		return fmt.Errorf("%s: %d", constant.ErrDonationAmountBelowMinimum, settings.MinAmount)
	}

	// create donation entity
	amount, err := valueobject.NewMoney(payload.Amount)
	if err != nil {
		return err
	}

	donation, err := entity.NewDonation(
		user.ID,
		amount,
		constant.PaymentMethod(payload.PaymentMethod),
	)
	if err != nil {
		return err
	}

	donation.SetDonorName(payload.Name)
	donation.SetMessage(payload.Message)

	// save donation to database
	tx, err := s.uow.Begin(ctx)
	if err != nil {
		return err
	}

	err = tx.DonationWriter().Save(ctx, donation)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}
		return err
	}

	// process payment via payment gateway
	paymentPayload := request.CreatePayment{
		UserID:      user.ID,
		ReferenceID: donation.ID,
		Amount:      float64(donation.Amount.Number()),
		Method:      donation.PaymentMethod,
	}

	paymentURL, err := s.paymentMapper.Pay(ctx, paymentPayload)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}
		return fmt.Errorf("failed to create payment: %w", err)
	}

	// store payment gateway reference
	donation.SetPaymentGatewayRef(&paymentURL)
	err = tx.DonationWriter().Save(ctx, donation)
	if err != nil {
		if uowErr := tx.Rollback(); uowErr != nil {
			return uowErr
		}
		return err
	}

	// commit transaction
	uowErr := tx.Commit()
	if uowErr != nil {
		return uowErr
	}

	return nil
}
