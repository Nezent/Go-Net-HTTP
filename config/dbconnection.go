package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	connStr := "postgres://postgres:anon@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
