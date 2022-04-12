package main

import (
	"fmt"
)

func unique(str []string) []string {
	i := 0
	for _, s := range str {
		if str[i] == s {
			continue
		}
		fmt.Println("Here", i)
		i++
		str[i] = s
		
	}
	return str[:i+1]
}

func main() {
	str := []string{"hello", "hello", "abc", "dab", "bax"}
	result := unique(str)
	fmt.Println("Unique", result)
}