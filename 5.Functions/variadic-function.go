package main

import (
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}
func main() {

	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum ", s)
}