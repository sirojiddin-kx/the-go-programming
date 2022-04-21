package main

import (
	"fmt"
)

func join(sep string, strings ...string) string {
	var (
		result string
	)

	for _, str := range strings {
		result += str + sep
	}

	return result
}

func main() {

	result := join("-", "Sirojiddin", "Khomidov", "Akrom")
	fmt.Println("result : ", result[:len(result)-1])
}
