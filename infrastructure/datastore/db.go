package datastore

import (
	"encoding/csv"
	"log"
	"os"
)

// NewDB returns a new CSV Reader with the DB content
func NewDB() *csv.Reader {
	csvfile, err := os.Open("tasks.csv")
	if err != nil {
		log.Fatalln("Could not open the CSV file", err)
	}

	return csv.NewReader(csvfile)
}
