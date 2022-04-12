package main

import "fmt"

func reverse(arr *[10]int) {
	l := len(arr)
	for i := 0; i < l / 2; i++ {
		j := l - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&arr)
	fmt.Println(arr)

}