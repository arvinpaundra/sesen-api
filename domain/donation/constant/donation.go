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

func (pm PaymentMethod) String() string {
	return string(pm)
}

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
