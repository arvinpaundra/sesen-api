package constant

import "errors"

var (
	ErrWidgetSettingsNotFound = errors.New("widget settings not found")
	ErrWidgetQrCodeNotFound   = errors.New("widget QR code not found")
	ErrWidgetAlertNotFound    = errors.New("widget alert not found")

	ErrUserAlreadyHasWidgetSettings = errors.New("user already has widget settings")
	ErrInvalidColorFormat           = errors.New("invalid color format")

	ErrWidgetTransactionFailed = errors.New("widget: transaction operation failed in nested context")
)
