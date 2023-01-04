package main

import "fmt"

// Interfaces are used in Go where polymorphism is needed.
// In a function where multiple types can be passed an interface can be used.
// Interfaces allow Go to have polymorphism. Interfaces are also used to help
// reduce duplication/boilerplate code.

// Employee is an interface for printing employee details

type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("Emp Id:", e)
	fmt.Println("Emp name:", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	fmt.Println("Salary: ", salary)
	return salary
}

func main() {
	var e1 = Emp(1)
	e1.PrintName("John Doe")
	e1.PrintSalary(25000, 5)
}
