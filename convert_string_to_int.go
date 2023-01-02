package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// You can use the strconv package's Atoi() function to convert the string into an integer value.
// Atoi stands for ASCII to integer. The Atoi() function returns two values:
// the result of the conversion, and the error (if any).

// func Atoi(s string) (int, error)
func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}
