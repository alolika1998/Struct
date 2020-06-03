package main

import "fmt"

	type Employee struct{
	Name string
	EmpID int
	Department string
	Skills []string
	}
func main() {
		aEmployee := Employee {
			Name: "Alolika",
			EmpID: 201,
			Department: "IT",
			Skills: []string {
				"C",
				"C++",
				"Java",
			},
		}
		fmt.Println(aEmployee)
	}
