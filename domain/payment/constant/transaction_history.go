package constant

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
