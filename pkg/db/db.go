package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var (
	db  *sql.DB
	err error
)

const (
	dbFile               = ".midori.db"
	dbEngine             = "sqlite3"
	journalTableCreation = `CREATE TABLE IF NOT EXISTS journal (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		command TEXT,
		context TEXT,
		entryType TEXT
	)`

	getEntries = "SELECT id, timestamp, command, context, entryType FROM journal ORDER BY timestamp DESC"
	addEntry   = "INSERT INTO journal (command, context, entryType) VALUES (?, ?, ?)"
)

// JournalEntry represents the structure of a journal entry
type JournalEntry struct {
	ID        int
	Timestamp time.Time
	Command   string
	Context   string
	Type      string
}

func init() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open(dbEngine, fmt.Sprintf("%s/%s", homeDir, dbFile))
	if err != nil {
		log.Fatal(err)
	}

	setupTables()
}

func Close() {
	db.Close()
}

func setupTables() {
	// Create journal_entries table if not exists
	_, err = db.Exec(journalTableCreation)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func GetEntries() ([]JournalEntry, error) {
	var err error
	rows, err := db.Query(getEntries)
	if err != nil {
		return []JournalEntry{}, err
	}
	defer rows.Close()

	var entries []JournalEntry

	for rows.Next() {
		var entry JournalEntry
		err := rows.Scan(&entry.ID, &entry.Timestamp, &entry.Command, &entry.Context, &entry.Type)
		if err != nil {
			return []JournalEntry{}, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func AddEntry(command, context, typeStr string) error {
	// Insert new entry into the database
	_, err := db.Exec(addEntry, command, context, typeStr)
	if err != nil {
		return err
	}

	return nil
}
