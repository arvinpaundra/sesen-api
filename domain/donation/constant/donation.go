package constant

type DonationStatus string

const (
	DonationStatusPending   DonationStatus = "pending"
	DonationStatusCompleted DonationStatus = "completed"
	DonationStatusFailed    DonationStatus = "failed"
	DonationStatusCancelled DonationStatus = "cancelled"
	DonationStatusExpired   DonationStatus = "expired"
)

type PaymentMethod string

const (
	PaymentMethodGopay     PaymentMethod = "gopay"
	PaymentMethodShopeepay PaymentMethod = "shopeepay"
	PaymentMethodDana      PaymentMethod = "dana"
	PaymentMethodQris      PaymentMethod = "qris"
	PaymentMethodLinkAja   PaymentMethod = "link_aja"
)

const (
	DefaultCurrency  = "IDR"
	CategoryDonation = "donation"
)
