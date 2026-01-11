package valueobject

import "github.com/arvinpaundra/sesen-api/domain/user/constant"

type Money struct {
	amount int64
}

func NewMoney(amount int64) (Money, error) {
	if amount < 0 {
		return Money{}, constant.ErrBalanceLessThanZero
	}

	return Money{amount: amount}, nil
}

func (b *Money) Number() int64 {
	return b.amount
}

func (b *Money) Add(amount int64) int64 {
	total := b.amount + amount

	return total
}

func (b *Money) Deduct(amount int64) (int64, error) {
	total := b.amount - amount

	if total < 0 {
		return 0, constant.ErrBalanceLessThanZero
	}

	return total, nil
}
