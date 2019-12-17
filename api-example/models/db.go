package models

import (
	"database/sql"
	"fmt"
	"log"
)

var conn *sql.DB

// Connect -> Connect to db
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin12345@tcp(localhost:3306)/dbnotes")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Connected")
	conn = db
	return db
}
