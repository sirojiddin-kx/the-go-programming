package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	z := strconv.Itoa(x)
	fmt.Printf("%T, %v", y, y)
	fmt.Printf("%T, %v", z, z)
}