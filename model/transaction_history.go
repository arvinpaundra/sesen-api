package model

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "credit"
	TransactionTypeDebit  TransactionType = "debit"
)

type TransactionCategory string

const (
	TransactionCategoryDonation TransactionCategory = "donation"
	TransactionCategoryPayout   TransactionCategory = "payout"
	TransactionCategoryOther    TransactionCategory = "other"
)

type TransactionHistory struct {
	ID            uuid.UUID           `gorm:"type:uuid;primaryKey"`
	UserId        uuid.UUID           `gorm:"type:uuid;column:user_id"`
	Type          TransactionType     `gorm:"column:type"`
	Category      TransactionCategory `gorm:"column:category"`
	Amount        int64               `gorm:"column:amount"`
	BalanceBefore int64               `gorm:"column:balance_before"`
	BalanceAfter  int64               `gorm:"column:balance_after"`
	ReferenceId   null.String         `gorm:"column:reference_id;unique"`
	Description   null.String         `gorm:"column:description"`
	CreatedAt     int64               `gorm:"column:created_at"`
}
