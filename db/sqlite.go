package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	dbPath := os.Getenv("DATABASE_URL")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	db.SetMaxOpenConns(1)
	return db
}

// PRAGMA journal_mode=WAL;
// db.Exec("PRAGMA journal_mode=WAL;")
// ->
// concurrent readers
// better performance
