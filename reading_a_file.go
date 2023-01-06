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
	scanner.Split(bufio.ScanWords)
	// var textLines []string

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	file.Close()
}
