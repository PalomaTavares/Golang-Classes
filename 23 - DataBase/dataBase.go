package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnect := "golang:golang@/aqua?charset=utf8&parseTime=True&loc=UTC"

	db, error := sql.Open("mysql", stringConnect)

	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("Connection open")

	lines, error := db.Query("select * from users")
	if error != nil {
		log.Fatal(error)
	}
	defer lines.Close()

	fmt.Println(lines)
}
