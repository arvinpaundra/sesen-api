package entity

import (
	"github.com/arvinpaundra/sesen-api/core/trait"
	"github.com/arvinpaundra/sesen-api/core/util"
	"github.com/arvinpaundra/sesen-api/domain/payment/constant"
	"github.com/arvinpaundra/sesen-api/domain/shared/valueobject"
)

type TransactionHistory struct {
	trait.Createable
	trait.Updateable
	trait.Removeable

	ID            string
	UserID        string
	Type          constant.TransactionType
	Category      constant.TransactionCategory
	Amount        valueobject.Money
	BalanceBefore valueobject.Money
	BalanceAfter  valueobject.Money
	ReferenceID   *string
	Description   *string
}

func NewTransactionHistory(
	userID string,
	typ constant.TransactionType,
	category constant.TransactionCategory,
	amount, balanceBefore, balanceAfter valueobject.Money,
) *TransactionHistory {
	history := &TransactionHistory{
		ID:            util.GenerateUUID(),
		UserID:        userID,
		Type:          typ,
		Category:      category,
		Amount:        amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceAfter,
	}

	history.MarkCreate()

	return history
}
