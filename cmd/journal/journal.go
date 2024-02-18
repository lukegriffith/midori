package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lukegriffith/midori/pkg/db"
	"github.com/urfave/cli/v2"
)

func main() {
	// Open SQLite database
	dbCon := db.GetDBCon()
	defer dbCon.Close()
	var err error

	// Initialize CLI app
	app := &cli.App{
		Name:  "journal",
		Usage: "Store journal entries to SQLite database",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new journal entry",
				Action:  Add,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all journal entries",
				Action:  List,
			},
		},
	}

	// Run the CLI app
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Add(c *cli.Context) error {
	var err error
	command := c.Args().First()
	context := c.Args().Get(1)
	typeStr := c.Args().Get(2)

	if command == "" {
		return fmt.Errorf("please provide a command")
	}

	// Insert new entry into the database
	db.AddEntry(command, context, typeStr)
	if err != nil {
		return err
	}

	fmt.Println("Journal entry added successfully.")
	return nil
}

func List(c *cli.Context) error {

	entries, err := db.GetEntries()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID\tTimestamp\t\tCommand\t\tContext\t\tType")
	fmt.Println("----------------------------------------------")
	for _, entry := range entries {
		fmt.Printf("%d\t%s\t%s\t%s\t%s\n", entry.ID, entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Command, entry.Context, entry.Type)
	}
	return nil
}
