package main

import (
	"fmt"
	"github.com/lukegriffith/midori/pkg/journal"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var isCmd bool

func main() {
	// Open SQLite database
	defer journal.Close()
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
				Action:  add,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "cmd",
						Value:       false,
						Destination: &isCmd,
					},
				},
			},
			{
				Name:    "summarise",
				Aliases: []string{"s", "llm"},
				Usage:   "Summarise current journal",
				Action:  summarise,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all journal entries",
				Action:  list,
			},
		},
	}

	// Run the CLI app
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func add(c *cli.Context) error {
	var err error
	entry := c.Args().First()

	if entry == "" {
		return fmt.Errorf("please provide a command")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to determine PWD")
	}

	if isCmd {
		err = journal.AddCommand(entry, pwd)
		if err != nil {
			return err
		}

	} else {
		err = journal.AddJournal(entry, pwd)
		if err != nil {
			return err
		}
		fmt.Println("Journal entry added successfully.")
	}

	return nil
}

func list(c *cli.Context) error {
	output, err := journal.ListJournal()
	fmt.Println(output)
	return err
}

func summarise(c *cli.Context) error {
	return nil
}
