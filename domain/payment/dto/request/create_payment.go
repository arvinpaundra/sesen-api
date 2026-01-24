package request

type CreatePaymentPayload struct {
	UserID      string
	ReferenceID string
	Amount      float64
	Method      string
	Category    string
}
