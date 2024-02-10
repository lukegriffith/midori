package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

// JournalEntry represents the structure of a journal entry
type JournalEntry struct {
	ID        int
	Timestamp time.Time
	Command   string
	Context   string
}

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite3", "journal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create journal_entries table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS journal_entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		command TEXT,
		context TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize CLI app
	app := &cli.App{
		Name:  "journal",
		Usage: "Store journal entries to SQLite database",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new journal entry",
				Action: func(c *cli.Context) error {
					command := c.Args().First()
					context := c.Args().Get(1)

					if command == "" {
						return fmt.Errorf("please provide a command")
					}

					// Insert new entry into the database
					_, err := db.Exec("INSERT INTO journal_entries (command, context) VALUES (?, ?)", command, context)
					if err != nil {
						return err
					}

					fmt.Println("Journal entry added successfully.")
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all journal entries",
				Action: func(c *cli.Context) error {
					rows, err := db.Query("SELECT id, timestamp, command, context FROM journal_entries ORDER BY timestamp DESC")
					if err != nil {
						return err
					}
					defer rows.Close()

					fmt.Println("ID\tTimestamp\t\tCommand\t\tContext")
					fmt.Println("----------------------------------------------")
					for rows.Next() {
						var entry JournalEntry
						err := rows.Scan(&entry.ID, &entry.Timestamp, &entry.Command, &entry.Context)
						if err != nil {
							return err
						}
						fmt.Printf("%d\t%s\t%s\t%s\n", entry.ID, entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Command, entry.Context)
					}

					return nil
				},
			},
		},
	}

	// Run the CLI app
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
