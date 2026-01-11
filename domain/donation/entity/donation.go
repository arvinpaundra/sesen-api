package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/donation/constant"
	"github.com/arvinpaundra/sesen-api/domain/shared/valueobject"
)

type Donation struct {
	trait.Createable
	trait.Updateable
	trait.Removeable

	ID                string
	ToUserID          string
	Amount            valueobject.Money
	Currency          string
	Status            constant.DonationStatus
	PaymentMethod     constant.PaymentMethod
	PaymentGatewayRef *string
	DonorName         *string
	Message           *string
}

func NewDonation(toUserID string, amount valueobject.Money, paymentMethod constant.PaymentMethod) (*Donation, error) {
	donation := &Donation{
		ID:            util.GenerateUUID(),
		ToUserID:      toUserID,
		Amount:        amount,
		Currency:      constant.DefaultCurrency,
		Status:        constant.DonationStatusPending,
		PaymentMethod: paymentMethod,
	}

	donation.MarkCreate()

	return donation, nil
}

func (d *Donation) SetDonorName(name *string) {
	if name != nil {
		d.DonorName = name
	}
}

func (d *Donation) SetMessage(message *string) {
	if message != nil {
		d.Message = message
	}
}

func (d *Donation) SetPaymentGatewayRef(ref *string) {
	if ref != nil {
		d.PaymentGatewayRef = ref
	}
}
