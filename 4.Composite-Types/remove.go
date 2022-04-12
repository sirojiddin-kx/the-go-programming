package main

import (
	"fmt"
)

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	fmt.Println(slice[i])
	fmt.Println(slice)
	return slice[:len(slice) -1]
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(remove(s, 2))
	fmt.Println(remove2(s, 2))
}