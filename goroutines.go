package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ResponseSize(url string) {

	fmt.Println("Step 1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step 2: ", url)

	defer response.Body.Close()

	fmt.Println("Step3 : ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step 4: ", len(body))
}
func main() {
	go ResponseSize("https://www.golangprograms.com")
	go ResponseSize("https://coderwall.com")
	go ResponseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
}
