package request

import "github.com/arvinpaundra/sesen-api/domain/donation/constant"

type CreatePayment struct {
	UserID      string
	ReferenceID string
	Amount      float64
	Method      constant.PaymentMethod
}
