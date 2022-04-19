package main

import "fmt"

type Employee struct {
	Name, Surname string
	Address string
	Salary int
}

func main() {
	var dilbert Employee = Employee{Name: "Dilbert", Surname: "Enst", Address: "Broadway", Salary: 2000}
	fmt.Println(dilbert.Salary)
	dilbert.Salary = Bonus(&dilbert, 600)
	fmt.Println(dilbert.Salary)
}

func Bonus(e *Employee, percent int) int {

	return e.Salary  * percent / 100
}