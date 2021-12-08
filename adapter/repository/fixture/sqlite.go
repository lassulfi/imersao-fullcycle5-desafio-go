package fixture

import (
	"context"
	"database/sql"
	"os"

	"github.com/maragudk/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDb(db *sql.DB) *sql.DB {
	migrationsDir := os.DirFS("adapter/repository/fixture/sql")
	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	return db
}
