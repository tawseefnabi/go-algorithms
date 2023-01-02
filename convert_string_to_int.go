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

	// ParseInt() Function
	// ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and
	// returns the corresponding value i. This function accepts a string parameter, convert it into a
	//  corresponding int type based on a base parameter. By default, it returns Int64 value.

	intVar8, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar8, err, reflect.TypeOf(intVar))
	intVar16, err := strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar16, err, reflect.TypeOf(intVar))

}
