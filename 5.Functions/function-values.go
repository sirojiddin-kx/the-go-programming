package main

import (
	"fmt"
)

func square(n int) int {

	return n * n
}

func product(m, n int) int {

	return m + n
}

func main() {

	f := square
	fmt.Println(f(3))

	f = product
	fmt.Println(g(3, 4))
}