package entity

import "errors"

const (
	APPROVED = "approved"
	REJECTED = "rejected"
)

type Transaction struct {
	ID        string
	AccountID string
	Amount    float64
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}

	return nil
}
