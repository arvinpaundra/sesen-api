package constant

import "errors"

var (
	ErrDonationAmountBelowMinimum = errors.New("donation amount is below minimum")
)
