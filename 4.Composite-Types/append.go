package main

import (
	"fmt"
)

func main() {

	var runes []rune

	for _, r := range "Hello World" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)
}