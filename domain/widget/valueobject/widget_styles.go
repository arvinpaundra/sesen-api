package valueobject

import (
	"maps"

	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
)

type WidgetStyles struct {
	styles map[constant.StyleKey]string
}

func NewWidgetStyles(defaultStyles map[constant.StyleKey]string) WidgetStyles {
	// copy for immutability
	copied := make(map[constant.StyleKey]string)

	maps.Copy(copied, defaultStyles)

	return WidgetStyles{styles: copied}
}

func (ws WidgetStyles) Get(key constant.StyleKey) string {
	return ws.styles[key]
}

func (ws WidgetStyles) Set(key constant.StyleKey, value string) WidgetStyles {
	// copy for immutability
	copied := make(map[constant.StyleKey]string)

	maps.Copy(copied, ws.styles)

	copied[key] = value

	return WidgetStyles{styles: copied}
}

func (ws WidgetStyles) Map() map[constant.StyleKey]string {
	// copy for immutability
	copied := make(map[constant.StyleKey]string)

	maps.Copy(copied, ws.styles)

	return copied
}
