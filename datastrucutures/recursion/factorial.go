package main

import "fmt"

func main() {
	fmt.Println("=", fact(4))
}

func fact(n int) int {
	if n == 0 || n == 1 { // base case
		return 1
	} else {
		fmt.Print(n, " ")
		return n * fact(n-1) // recursive case
	}

}
