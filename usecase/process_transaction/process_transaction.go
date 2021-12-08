package process_transaction

import (
	"github.com/lassulfi/imersao5-desafiogo/domain/entity"
	"github.com/lassulfi/imersao5-desafiogo/domain/repository"
)

type ProcessTransacion struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransacion {
	return &ProcessTransacion{Repository: repository}
}

func (p *ProcessTransacion) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectTransaction(*transaction, invalidTransaction)
	}

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount)
	if err != nil {
		return p.rejectTransaction(*transaction, err)
	}

	output := TransactionDtoOutput{
		ID:           input.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	return output, nil
}

func (p *ProcessTransacion) rejectTransaction(transaction entity.Transaction, err error) (TransactionDtoOutput, error) {
	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: err.Error(),
	}

	return output, err
}
