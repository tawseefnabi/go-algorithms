package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// using ioutil
	fmt.Println("*********** opening file using ioutil ***********")
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal("failed opening file: %s", err)
	}
	fmt.Println(string(content))

	fmt.Println("*********** opening file using scanner ***********")
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string

	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}
	file.Close()
	for _, eacLine := range textLines {
		fmt.Println(eacLine)
	}
}
