package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/widget/constant"
	"github.com/arvinpaundra/sesen-api/domain/widget/valueobject"
)

type WidgetAlert struct {
	trait.Createable
	trait.Updateable
	trait.Removeable

	ID            string
	SettingID     string
	AlertText     string
	AlertURL      string
	AttachmentURL *string
	Styles        valueobject.WidgetStyles
}

func NewWidgetAlert(settingID string) *WidgetAlert {
	alert := &WidgetAlert{
		ID:        util.GenerateUUID(),
		SettingID: settingID,
		AlertText: constant.DefaultAlertText,
		AlertURL:  constant.DefaultAlertURL,
		Styles:    valueobject.NewWidgetStyles(constant.DefaultAlertStyles),
	}

	alert.MarkCreate()

	return alert
}

func (w *WidgetAlert) SetAttachmentURL(url string) {
	w.AttachmentURL = &url
	w.MarkUpdate()
}
