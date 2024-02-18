package journal

import (
	"bytes"
	"fmt"
	"github.com/lukegriffith/midori/pkg/db"
	"log"
)

var (
	commandTypeStr = "command"
	journalTypeStr = "journal"
)

func ListJournal() (string, error) {
	var buffer bytes.Buffer

	entries, err := db.GetEntries()
	if err != nil {
		log.Fatal(err)
	}

	buffer.WriteString("ID\tTimestamp\t\tCommand\t\tContext\t\tType\n")
	buffer.WriteString("----------------------------------------------\n")
	for _, entry := range entries {
		line := fmt.Sprintf("%d\t%s\t%s\t%s\t%s\n", entry.ID, entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Command, entry.Context, entry.Type)
		buffer.WriteString(line)
	}

	groupedCommands, err := db.GroupEntries()
	if err != nil {
		log.Fatal(err)
	}

	buffer.WriteString("Count\tCommand\tWorking Dir\n")
	buffer.WriteString("----------------------------------------------\n")
	for _, group := range groupedCommands {
		line := fmt.Sprintf("%d\t\t%s\t%s\n", group.Count, group.Command, group.Context)
		buffer.WriteString(line)
	}

	return buffer.String(), nil
}

func AddCommand(command, pwd string) error {
	return db.AddEntry(command, pwd, commandTypeStr)
}

func AddJournal(entry, pwd string) error {
	return db.AddEntry(entry, pwd, journalTypeStr)
}

func Close() {
	db.Close()
}
