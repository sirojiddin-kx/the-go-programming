package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["Sirojiddin"] = 22
	ages["Jahongir"] = 23

	_, ok := ages["Bob"]
	if !ok {
		fmt.Println("Bob is not exsist in map")
	}
}
