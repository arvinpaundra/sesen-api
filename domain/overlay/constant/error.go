package constant

import "errors"

var (
	ErrOverlayNotFound       = errors.New("overlay not found")
	ErrUserAlreadyHasOverlay = errors.New("user already has an overlay")
	ErrInvalidHexColor       = errors.New("invalid hex color")
)
