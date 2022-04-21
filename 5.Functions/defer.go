package main

import "fmt"

func double(x int) (result int) {

	fmt.Println("Result", result)

	return x + x
}

func doubleTwo(x int) (result int) {

	defer func() {
		fmt.Println("Result", result)
	}()

	return x + x
}

func triple(x int) (result int) {

	defer func() {
		result += x
	}()

	return double(x)
}

func main() {

	fmt.Println(double(4))
	fmt.Println(doubleTwo(4))

	fmt.Println(triple(4))
}
