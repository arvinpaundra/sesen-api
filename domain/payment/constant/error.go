package constant

import "errors"

var (
	ErrPaymentNotFound = errors.New("payment not found")

	ErrInvalidPaymentStatus = errors.New("invalid payment status")
	ErrInvalidPaymentAmount = errors.New("invalid payment amount")
)
