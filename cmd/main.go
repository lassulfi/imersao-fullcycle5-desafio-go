package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lassulfi/imersao5-desafiogo/adapter/factory"
	"github.com/lassulfi/imersao5-desafiogo/adapter/repository/fixture"
	"github.com/lassulfi/imersao5-desafiogo/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Proccessing account")
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	db = fixture.CreateDb(db)

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    100,
	}
	fmt.Printf("Account to process %v\n", input)

	usecase := process_transaction.NewProcessTransaction(repository)
	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Printf("Error while proccess account %v\n", output)
	} else {
		fmt.Printf("Success while proccess account %v\n", output)
	}

}
