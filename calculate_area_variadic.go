package main

import "fmt"

func calculate_area(str string, y ...int) int {
	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}
func main() {

	fmt.Println("Area of Rectangle: ", calculate_area("Rectangle", 20, 30))
	fmt.Println("Area of Square: ", calculate_area("Square", 20))
}
