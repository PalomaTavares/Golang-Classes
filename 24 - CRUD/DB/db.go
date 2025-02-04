package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //connection driver
)

// opne conncetion with database
func Connect() (*sql.DB, error) {
	stringConnect := "golang:golang@/aqua?charset=utf8&parseTime=True&loc=UTC"

	db, error := sql.Open("mysql", stringConnect)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil

}
