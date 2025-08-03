package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	vo "github.com/arvinpaundra/sesen-api/domain/overlay/value_object"
)

type Overlay struct {
	trait.Createable
	trait.Updateable

	ID            string
	UserId        string
	RingtoneUrl   string
	IsTTSEnabled  bool
	IsNFSWEnabled bool
	CreatedAt     string

	QRCode  *OverlayQR
	Message *OverlayMessage
}

func NewOverlay(
	userId, ringtoneUrl string,
	isTTSEnabled, isNFSWEnabled bool,
) *Overlay {
	overlay := &Overlay{
		ID:            util.GenerateUUID(),
		UserId:        userId,
		RingtoneUrl:   ringtoneUrl,
		IsTTSEnabled:  isTTSEnabled,
		IsNFSWEnabled: isNFSWEnabled,
	}

	overlay.MarkCreate()

	return overlay
}

func (o *Overlay) SetQRCode(qr *OverlayQR) error {
	if qr == nil {
		return nil
	}

	o.QRCode = qr

	return nil
}

func (o *Overlay) HasQRCode() bool {
	return o.QRCode != nil
}

func (o *Overlay) SetMessage(message *OverlayMessage) error {
	if message == nil {
		return nil
	}

	o.Message = message

	return nil
}

func (o *Overlay) HasMessage() bool {
	return o.Message != nil
}

type OverlayQR struct {
	trait.Createable
	trait.Updateable

	ID              string
	OverlayId       string
	Code            string
	QrColor         vo.HexColor
	BackgroundColor vo.HexColor
	CreatedAt       string
}

func NewOverlayQR(overlayId, code, qrColor, backgroundColor string) (*OverlayQR, error) {
	qrColorHex, err := vo.NewHexColor(qrColor)
	if err != nil {
		return nil, err
	}

	backgroundColorHex, err := vo.NewHexColor(backgroundColor)
	if err != nil {
		return nil, err
	}

	qr := &OverlayQR{
		ID:              util.GenerateUUID(),
		OverlayId:       overlayId,
		Code:            code,
		QrColor:         qrColorHex,
		BackgroundColor: backgroundColorHex,
	}

	qr.MarkCreate()

	return qr, nil
}

type OverlayMessage struct {
	trait.Createable
	trait.Updateable

	ID              string
	OverlayId       string
	TextColor       vo.HexColor
	BackgroundColor vo.HexColor
	CreatedAt       string
	UpdatedAt       string
}

func NewOverlayMessage(overlayId, textColor, backgroundColor string) (*OverlayMessage, error) {
	textColorHex, err := vo.NewHexColor(textColor)
	if err != nil {
		return nil, err
	}

	backgroundColorHex, err := vo.NewHexColor(backgroundColor)
	if err != nil {
		return nil, err
	}

	message := &OverlayMessage{
		ID:              util.GenerateUUID(),
		OverlayId:       overlayId,
		TextColor:       textColorHex,
		BackgroundColor: backgroundColorHex,
	}

	message.MarkCreate()

	return message, nil
}
