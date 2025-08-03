package vo

import (
	"regexp"

	"github.com/arvinpaundra/sesen-api/domain/overlay/constant"
)

type HexColor string

func NewHexColor(color string) (HexColor, error) {
	rg, err := regexp.Compile(`^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$`)
	if err != nil {
		return "", err
	}

	if !rg.MatchString(color) {
		return "", constant.ErrInvalidHexColor
	}

	return HexColor(color), nil
}

func (h HexColor) String() string {
	return string(h)
}
