package valueobject

import (
	"strings"

	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
)

type Color struct {
	hex string
}

func NewColor(hex string) (Color, error) {
	color := Color{hex}

	if !color.isValid() {
		return Color{}, constant.ErrInvalidColorFormat
	}

	return color, nil
}

func (c Color) String() string {
	return strings.ToUpper(c.hex)
}

func (c Color) isValid() bool {
	if len(c.hex) < 4 || len(c.hex) > 9 {
		return false
	}

	if !strings.HasPrefix(c.hex, "#") {
		return false
	}

	return true
}
