package main

import (
	"fmt"
)

type rectangle struct {
	length  int
	breadth int
	color   string
}
type rectangle1 struct {
	length   int
	breadth  int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
}

func main() {
	// A struct (short for "structure") is a collection of data fields with declared data types.
	// Golang has the ability to declare and create own data types by combining one or more types, including both built-in and user-defined types.
	// Each data field in a struct is declared with a known type, which could be a built-in type or another user-defined type
	fmt.Println("-----------------------   -----------------------")
	// #1: Creating a Struct Instance Using a Struct Literal
	fmt.Println(rectangle{10, 2, "yellow"})

	// #2 Struct Instantiation Using new Keyword

	fmt.Println("-----------------------   -----------------------")
	rec1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rec1.length = 10
	rec1.breadth = 20
	rec1.color = "red"
	fmt.Println("rec1", rec1)

	var rect2 = new(rectangle) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)

	fmt.Println("-----------------------   -----------------------")
	// #3 dot notation
	var rec rectangle1
	rec.length = 5
	rec.breadth = 3
	rec.color = "green"
	rec.geometry.area = rec.length * rec.breadth
	rec.geometry.perimeter = 2 * (rec.length + rec.breadth)
	fmt.Println("rec", rec, "\nArea: ", rec.geometry.area, "\nPermiter: ", rec.geometry.perimeter)
}
