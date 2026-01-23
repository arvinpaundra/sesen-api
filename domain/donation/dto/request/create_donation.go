package request

type CreateDonationPayload struct {
	Username      string  `json:"username" validate:"required"`
	Amount        int64   `json:"amount" validate:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" validate:"required,oneof=gopay shopeepay dana qris link_aja"`
	Name          *string `json:"name"`
	Message       *string `json:"message"`
}

func (p *CreateDonationPayload) IsAmountValid(minAmount int64) bool {
	return p.Amount >= minAmount
}
