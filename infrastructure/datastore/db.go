package datastore

import (
	"encoding/csv"
	"log"
	"os"
)

var csvfile *os.File

// NewDB returns a new CSV Reader with the DB content
func NewDB() *csv.Reader {
	csvfile, err := os.Open("tasks.csv")
	if err != nil {
		log.Fatalln("Could not open the CSV file", err)
	}

	return csv.NewReader(csvfile)
}

// CloseDB closes the file used as DB
func CloseDB() {
	csvfile.Close()
}
