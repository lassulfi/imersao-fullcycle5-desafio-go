package repository

import (
	"os"
	"testing"

	"github.com/lassulfi/imersao5-desafiogo/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
)

func TestTransacionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 500)

	assert.Nil(t, err)
}
