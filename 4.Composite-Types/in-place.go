package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func main() {
	data := []string{"one", "two", "", "three"}
	fmt.Println(nonempty(data))	
	str := "Hello World"
	fmt.Println(str[1:])
}