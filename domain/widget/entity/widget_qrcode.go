package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
	"github.com/arvinpaundra/sesen-api/domain/widget/valueobject"
)

type WidgetQrcode struct {
	trait.Createable
	trait.Updateable
	trait.Removeable

	ID          string
	SettingID   string
	QrCodeData  string
	Description string
	Styles      valueobject.WidgetStyles
}

func NewWidgetQrcode(settingID, qrCodeData string) *WidgetQrcode {
	qrcode := &WidgetQrcode{
		ID:          util.GenerateUUID(),
		SettingID:   settingID,
		QrCodeData:  qrCodeData,
		Description: constant.DefaultQrCodeDescription,
		Styles:      valueobject.NewWidgetStyles(constant.DefaultQrCodeStyles),
	}

	qrcode.MarkCreate()

	return qrcode
}
