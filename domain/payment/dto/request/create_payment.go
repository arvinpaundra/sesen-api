package request

type CreatePaymentPayload struct {
	UserID      string
	ReferenceID string
	UserName    string
	Amount      float64
	Method      string
	Category    string
}
