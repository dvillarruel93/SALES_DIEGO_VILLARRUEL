package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Database interface {
	SelectOne(query string, destinationResult ...interface{}) error
	SelectMultiple(query string, arguments interface{}) (*sql.Rows, error)
	ExecuteQuery(query string) (sql.Result, error)
}

type DBHandler struct {
	Database
}

func NewDatabase() Database {
	openConnection()
	return DBHandler{}
}

func openConnection() {
	var err error
	connectionString := "docker:docker@tcp(db:3306)/ticket_test"
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		errorMessage := fmt.Sprintf("Could not connect to database in '%s'", connectionString)
		log.Print(errorMessage)
	}

	dbError := db.Ping()

	if dbError != nil {
		log.Print("Could not ping to database")
	}
}

func (dbHandler DBHandler) SelectMultiple(query string, arguments interface{}) (*sql.Rows, error) {
	stmtOut, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if arguments != nil {
		rows, err = stmtOut.Query(arguments)
	} else {
		rows, err = stmtOut.Query()
	}

	stmtOut.Close()

	if err != nil {
		return nil, err
	}

	return rows, err
}

func (dbHandler DBHandler) SelectOne(query string, destinationResult ...interface{}) error {
	stmtOut, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmtOut.Close()
	err = stmtOut.QueryRow().Scan(destinationResult...)
	return err
}

func (dbHandler DBHandler) ExecuteQuery(query string) (sql.Result, error) {
	stmtIns, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}

	res, err := stmtIns.Exec()
	stmtIns.Close()
	return res, err
}
