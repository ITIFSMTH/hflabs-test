package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hflabstest/sheets"
)

func main() {
	// Get sheetId
	var sheetId string
	var pathToCreds string
	flag.StringVar(&sheetId, "sheet", "", "Sheet ID")
	flag.StringVar(&pathToCreds, "path", "", "Path to ")
	flag.Parse()

	if sheetId == "" || pathToCreds == "" {
		fmt.Print("Write --help for check usage")
		return
	}

	// Read creds JSON
	creds, err := os.ReadFile(pathToCreds)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Update google sheets
	if err := sheets.UpdateGoogleSheets(creds, sheetId); err != nil {
		log.Fatalf("Error parsing table: %v", err)
	}

	fmt.Print("* Completed *")
}
