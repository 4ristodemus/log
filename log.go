package main

import (
	"fmt"
	"github.com/4ristodemus/log/cli"
	"github.com/4ristodemus/log/db"
	"os"
)

func main() {
	args := os.Args[1:]
	if !(len(args) > 0) {
		fmt.Println("An argument is required")
		os.Exit(3)
	}

	// Connect to DB & run migrations
	db.InitDatabase()

	cmd := args[0]
	switch cmd {
	case "serve":
		fmt.Printf("\nStarting server\n")
		router := CreateRouter()
		router.Run()
	default:
		cli.HandleArgs(args)
	}
}
