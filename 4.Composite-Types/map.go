package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["Alice"] = 22
	ages["Bob"] = 30
	ages["Sirojididn"] = 21
	
	ages2 := map[string]int{
		"Alice" : 22,
		"Charllie" : 30,
	}

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
