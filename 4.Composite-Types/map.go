package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["Alice"] = 22
	ages["Bob"] = 30
	
	ages2 := map[string]int{
		"Alice" : 22,
		"Charllie" : 30,
	}

	fmt.Println("Ages 2 ", ages2)
	ages2["Alice"] += 1
	fmt.Println(ages2["Alice"])
}
