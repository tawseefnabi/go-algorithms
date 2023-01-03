package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e *Employee) EmpInfo() string {
	fmt.Println("Name: ", e.FirstName, e.LastName)
	fmt.Println("Age: ", e.Age)
	fmt.Println("Email: ", e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println("Basic: ", info.Basic)
		fmt.Println("HRA: ", info.HRA)
		fmt.Println("TA: ", info.TA)
	}
	return "---------"
}
func main() {
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	e.EmpInfo()

}
