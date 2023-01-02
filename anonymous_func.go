package main

import (
	"fmt"
)

// An anonymous function is a function that was declared without any named identifier to refer to it.
// Anonymous functions can accept inputs and return outputs, just as standard functions do.

// syntax

/*
	func(parameter_list)(return_type){
		// code..

		// Use return statement if return_type are given
		// if return_type is not given, then do not
		// use return statement

return
}()
*/
func main() {

	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	var value = func() {
		fmt.Println("Welcome to anonymous")
	}
	value()

}
