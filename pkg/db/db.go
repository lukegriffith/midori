package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	db  *sql.DB
	err error
)

const (
	dbFile                 = "midori.db"
	dbEngine               = "sqlite3"
	journal_table_creation = `CREATE TABLE IF NOT EXISTS journal_entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		command TEXT,
		context TEXT
	)`

	command_table_creation = `CREATE TABLE IF NOT EXISTS command_log (
		id INTEGER PRIMARY KEY,
		command TEXT,
		executed_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
)

func init() {
	db, err = sql.Open(dbEngine, dbFile)
	if err != nil {
		log.Fatal(err)
	}

	setupTables()
}

func GetDBCon() *sql.DB {
	return db
}

func setupTables() {

	// Create journal_entries table if not exists
	_, err = db.Exec(journal_table_creation)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// Create command log table if it doesn't exist
	_, err = db.Exec(command_table_creation)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

}
