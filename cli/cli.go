package cli

import (
	"fmt"
	"os/exec"

	"github.com/4ristodemus/log/db"
	"github.com/4ristodemus/log/models"
)

type CLIAction string

const (
	New CLIAction = "new"
)

/*
 * Resolves command line args
 */
func HandleArgs(args []string) {
	action := CLIAction(args[0])

	switch action {
	case New:
		newEntry(args[1:])
	default:
		fmt.Println("Not a valid action")
	}
}

/*
 * Creates a new markdown entry
 */
func newEntry(args []string) {
	argMax := 2

	if db.Database == nil {
		panic("failed to connect database")
	}

	if argMax > len(args) {
		panic("An action and title are required to create an entry")
	}

	entryType := args[0]
	entryTitle := args[1]

	script := &exec.Cmd{
		Path: "./new_entry.sh",
		Args: []string{"./new_entry.sh", entryTitle},
	}

	if output, err := script.Output(); err != nil {
		fmt.Println(err)
		panic("Could not create markdown entry")
	} else {
		createRecord(
			entryTitle,
			entryType,
			"entries/"+string(output),
		)
	}
}

/*
 * Creates SQLite record for an entry
 */
func createRecord(
	entryTitle string,
	entryType string,
	entryPath string,
) {
	entry := models.Entry{
		Title: entryTitle,
		Type:  models.EntryType(entryType),
		Path:  entryPath,
	}

	result := db.Database.Create(&entry)

	if result.Error != nil {
		panic(result.Error)
	} else {
		fmt.Println("Successfully created entry.")
	}
}
