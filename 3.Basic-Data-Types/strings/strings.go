package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	fmt.Println(&s)
	t := s
	fmt.Println(&t)
	s += ", again me"
	fmt.Println(&s)
	query := `Select 
				name,
				fullName
			  from 
			  	person`
	fmt.Println(query)
}
