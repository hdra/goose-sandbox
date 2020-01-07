package main

import (
	"log"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")

	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatalf("migration error: %v", err)
	}
}
