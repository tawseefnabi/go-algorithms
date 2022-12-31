package main

import (
	"fmt"
	"time"
)

// The fallthrough keyword used to force the execution flow to fall through the successive case block.
// With the help of fallthrough statement, we can use to transfer the program control just after the statement is
// executed in the switch cases even if the expression does not match.
// Normally, control will come out of the statement of switch just after
// the execution of first line after match. Don’t put the fallthrough in the last statement of switch case.
func main() {
	today := time.Now()
	fmt.Println("time", today)

	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}
