package service

import (
	"context"
	"fmt"

	"github.com/arvinpaundra/sesen-api/domain/donation/constant"
	"github.com/arvinpaundra/sesen-api/domain/donation/entity"
	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	"github.com/arvinpaundra/sesen-api/domain/shared/valueobject"
)

type CreateDonationCommand struct {
	Username      string  `json:"username" validate:"required"`
	Amount        int64   `json:"amount" validate:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" validate:"required,oneof=gopay shopeepay dana qris link_aja"`
	Name          *string `json:"name"`
	Message       *string `json:"message"`
}

type CreateDonation struct {
	userMapper   repository.UserMapper
	widgetMapper repository.WidgetMapper
	uow          repository.UnitOfWork
}

func NewCreateDonation(
	userMapper repository.UserMapper,
	widgetMapper repository.WidgetMapper,
	uow repository.UnitOfWork,
) *CreateDonation {
	return &CreateDonation{
		userMapper:   userMapper,
		widgetMapper: widgetMapper,
		uow:          uow,
	}
}

func (s *CreateDonation) Execute(ctx context.Context, command CreateDonationCommand) error {
	// find donated user by username
	user, err := s.userMapper.FindUserByUsername(ctx, command.Username)
	if err != nil {
		return err
	}

	// validate amount against user's minimum donation amount
	settings, err := s.widgetMapper.FindWidgetSettingsByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	if command.Amount < settings.MinAmount {
		return fmt.Errorf("%s: %d", constant.ErrDonationAmountBelowMinimum, settings.MinAmount)
	}

	// create donation entity
	amount, err := valueobject.NewMoney(command.Amount)
	if err != nil {
		return err
	}

	donation, err := entity.NewDonation(
		user.ID,
		amount,
		constant.PaymentMethod(command.PaymentMethod),
	)
	if err != nil {
		return err
	}

	donation.SetDonorName(command.Name)
	donation.SetMessage(command.Message)

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

	// commit transaction
	uowErr := tx.Commit()
	if uowErr != nil {
		return uowErr
	}

	return nil
}
