package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	empData := [][]string{
		{"Name", "City", "Skills"},
		{"Smith", "Newyork", "Java"},
		{"William", "Paris", "Golang"},
		{"Rose", "London", "PHP"},
	}
	csvfile, err := os.Create("test.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range empData {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()

	csvfile.Close()
}
