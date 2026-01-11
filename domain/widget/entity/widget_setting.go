package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
)

type WidgetSetting struct {
	trait.Createable
	trait.Updateable
	trait.Removeable

	ID              string
	UserID          string
	TTSEnabled      bool
	NSFWFilter      bool
	MessageDuration int
	MinAmount       int64

	QrCode *WidgetQrcode
	Alert  *WidgetAlert
}

func NewWidgetSetting(userID string) *WidgetSetting {
	setting := &WidgetSetting{
		ID:              util.GenerateUUID(),
		UserID:          userID,
		NSFWFilter:      true,
		MessageDuration: int(constant.DefaultMessageDuration),
		MinAmount:       constant.DefaultMinimumDonationAmount,
	}

	setting.MarkCreate()

	return setting
}

func (w *WidgetSetting) AddQrCode(qrcode *WidgetQrcode) {
	w.QrCode = qrcode
}

func (w *WidgetSetting) HasQrCode() bool {
	return w.QrCode != nil
}

func (w *WidgetSetting) AddAlert(alert *WidgetAlert) {
	w.Alert = alert
}

func (w *WidgetSetting) HasAlert() bool {
	return w.Alert != nil
}
