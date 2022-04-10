package main

import (
	"fmt"
)

type Currency int 

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var q[3]int = [3]int{}
	fmt.Println(q)
	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p)
	p[2] = 4
	fmt.Println(p)

	symbol := [...]string{USD: "$", EUR: "e", GBP: "g", RMB: "r"}
	fmt.Println(USD, symbol[USD])
}

