package main

import (
	"log"
	"os"
)

func main() {
	oldName := "test.txt"
	newName := "testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}
