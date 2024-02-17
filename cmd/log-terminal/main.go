package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lukegriffith/midori/pkg/db"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./log-terminal <command>")
		os.Exit(1)
	}

	// Extract the command from arguments
	command := strings.Join(os.Args[1:], " ")

	db := db.GetDBCon()
	defer db.Close()

	var err error

	// Insert command into the database
	_, err = db.Exec("INSERT INTO command_log (command, executed_at) VALUES (?, ?)", command, time.Now())
	if err != nil {
		fmt.Println("Error inserting command into database:", err)
		os.Exit(1)
	}

	fmt.Println("Command logged successfully!")
}
