package main

import "fmt"

// Go has a special statement called defer that schedules a function call to be run after the function completes
// defer statement is often used with paired operations like open and close, connect and disconnect, or
// lock and unlock to ensure that resources are released in all cases, no matter how complex the control flow.
// The right place for a defer statement that releases a resource is immediately after the resource has been successfully acquired.

// Below is the example to open a file and perform read/write action on it.
// In this example there are often spots where you want to return early.
func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func main() {
	defer second()
	first()
}
