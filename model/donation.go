package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type DonationStatus string

const (
	DonationStatusPending   DonationStatus = "pending"
	DonationStatusCompleted DonationStatus = "completed"
	DonationStatusFailed    DonationStatus = "failed"
	DonationStatusCancelled DonationStatus = "cancelled"
	DonationStatusExpired   DonationStatus = "expired"
)

type PaymentMethod string

const (
	PaymentMethodGopay     PaymentMethod = "gopay"
	PaymentMethodShopeepay PaymentMethod = "shopeepay"
	PaymentMethodDana      PaymentMethod = "dana"
	PaymentMethodQris      PaymentMethod = "qris"
	PaymentMethodLinkAja   PaymentMethod = "link_aja"
)

type Donation struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ToUserId          uuid.UUID      `gorm:"type:uuid;column:to_user_id"`
	Amount            int64          `gorm:"column:amount"`
	Currency          string         `gorm:"column:currency"`
	Status            DonationStatus `gorm:"column:status"`
	PaymentMethod     PaymentMethod  `gorm:"column:payment_method"`
	PaymentGatewayRef null.String    `gorm:"column:payment_gateway_ref;nullable"`
	DonorName         null.String    `gorm:"column:donor_name;nullable"`
	Message           null.String    `gorm:"column:message;nullable"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at"`
}
