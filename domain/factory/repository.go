package factory

import "github.com/lassulfi/imersao5-desafiogo/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
