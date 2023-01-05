package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("test")
	if os.IsNotExist(err) {
		errDir := os.Mkdir("test", 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}
}
