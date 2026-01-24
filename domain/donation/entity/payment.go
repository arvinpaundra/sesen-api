package entity

type Payment struct {
	UserID      string
	ReferenceID string
	Amount      float64
	Method      string
}
