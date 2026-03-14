package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	dbPath := os.Getenv("DATABASE_URL")
	if dbPath == ""{
		dbPath = "./db/database.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	query := `CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY,
		long_link TEXT NOT NULL,
		short_link TEXT NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil{
		log.Fatal("error creating table:", err)
	}

	db.SetMaxOpenConns(1)
	return db
}

// PRAGMA journal_mode=WAL;
// db.Exec("PRAGMA journal_mode=WAL;")
// ->
// concurrent readers
// better performance
