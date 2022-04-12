package main

import (
	"fmt"
)

func rotate1(s []int, order int) []int {
	if order < 0 || len(s) == 0 {
		return s
	}

	r := len(s) - order % len(s)
	fmt.Println(s[r:])
	fmt.Println(s[:r])
	s = append(s[r:], s[:r]...)

	return s
}

func main() {
	nums := []int{1, 2, 3, 4 ,5 ,6, 7, 8}
	fmt.Println(rotate1(nums, 2))
}