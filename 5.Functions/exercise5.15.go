package main

import (
	"fmt"
)

func max(nums ...int) int {
	max := nums[0]

	if len(nums) == 0 {

		return 0
	}

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {

	maxNum := max(2, 3, 1, 8, 10, 6, 4)
	fmt.Println(maxNum)
	maxNum1 := max()
	fmt.Println(maxNum1)
}