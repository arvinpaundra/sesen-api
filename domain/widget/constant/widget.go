package constant

// Default widget constants
const (
	DefaultMessageDuration       = 10   // in seconds
	DefaultMinimumDonationAmount = 5000 // in rupiah
)

// Default widget colors
const (
	PrimaryColor = "#8B5CF6"
	WhiteColor   = "#FFFFFF"
	BgColor      = "#00000040"
)

type StyleKey string

func (sk StyleKey) String() string {
	return string(sk)
}

// Common Styles
const (
	BackgroundColorKey            StyleKey = "bg_color"
	MessageColorKey               StyleKey = "message_color"
	DescriptionColorKey           StyleKey = "description_color"
	DescriptionBackgroundColorKey StyleKey = "description_bg_color"
)

// QR Code Styles
const (
	QrCodeColorKey           StyleKey = "qr_color"
	QrCodeBackgroundColorKey StyleKey = "qr_bg_color"
)

// Alert Styles
const (
	AlertHighlightColorKey StyleKey = "alert_highlight_color"
)

// Default QR Code Styles
var DefaultQrCodeStyles = map[StyleKey]string{
	QrCodeColorKey:                PrimaryColor,
	QrCodeBackgroundColorKey:      WhiteColor,
	DescriptionColorKey:           WhiteColor,
	DescriptionBackgroundColorKey: BgColor,
}

// Default Alert Styles
var DefaultAlertStyles = map[StyleKey]string{
	AlertHighlightColorKey: PrimaryColor,
	MessageColorKey:        WhiteColor,
	BackgroundColorKey:     BgColor,
}

// Default qrcode values
const (
	DefaultQrCodeDescription = "Scan to support the stream!"
)

// Default alert values
const (
	DefaultAlertURL  = "https://exampple.com/alert.odd"
	DefaultAlertText = "{name} just donated {amount}!"
)
