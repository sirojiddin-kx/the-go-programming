package main

import (
	"fmt"
)

func basename(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1;i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {

	withSlash := basename("a/b/c.go")
	withDot := basename("a/b.c.go")
	fmt.Println("with Slash :", withSlash)
	fmt.Println("with Dot : ", withDot)
}