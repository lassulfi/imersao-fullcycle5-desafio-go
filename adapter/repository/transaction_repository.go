package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	DB *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{DB: db}
}

func (t *TransactionRepositoryDb) Insert(id string, accountId string, amount float64) error {
	stmt, err := t.DB.Prepare(`
		insert into transactions (id, account_id, amount, created_at)
		values ($1, $2, $3, $4)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
		accountId,
		amount,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
