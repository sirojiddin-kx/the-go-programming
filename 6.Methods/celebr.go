package main

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
}

func (p Person) Congratulate(s string) string {

	return p.FirstName  + " " + p.LastName + " congratulte you with " + s
}

func main() {

	p := Person{FirstName: "Sirojiddin", LastName: "Khomidov"}
	fmt.Println(p.Congratulate("Eid"))
}