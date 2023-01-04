package main

import "fmt"

type Vehicle interface {
	Structure() []string
	Speed() string
}

type Human interface {
	Structure() []string
	Perf() string
}
type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200KM/hr"
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}
func (m Man) Perf() string {
	return "40KM/Day"
}

func main() {
	var v1 Vehicle
	v1 = Car("BMW")
	fmt.Println("Speed of BMW:", v1.Speed())

	var h1 Man
	h1 = Man("Softwarer developer")
	fmt.Println("Speed of Human: ", h1.Perf)

	for i, j := range v1.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, h1.Structure()[i])
	}

}
