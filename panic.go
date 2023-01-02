package main

import (
	"errors"
	"fmt"
)

// Panic, indicates that something has gone wrong in such a way that the system can't continue to function and halt the execution of the program.
// Panics can be initiated by the developer and can be caused during the execution of a program by runtime errors.
// panic(interface{})
var result = 1

func chain(n int) {
	// Go provides us with the ability to regain control after panic has occurred.
	// The built-in recover() function is called from within a defer to check if a panic happened.
	// func recover() interface{}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if n == 0 {
		panic(errors.New("Cannot multiply a number by zero"))
	} else {
		result *= n
		fmt.Println("Output: ", result)
	}
}

func main() {
	chain(2)
	chain(3)
	chain(0)
	chain(4)
	chain(8)
}
