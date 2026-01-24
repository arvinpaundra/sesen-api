package request

type PaymentRequestPayload struct {
	ReferenceID        string  `json:"reference_id"`
	Amount             int64   `json:"amount"`
	Method             string  `json:"method"`
	Description        *string `json:"description,omitempty"`
	CustomerName       string  `json:"customer_name"`
	SuccessRedirectURL string  `json:"success_redirect_url"`
	FailureRedirectURL string  `json:"failure_redirect_url"`
	CallbackURL        string  `json:"callback_url"`
}
