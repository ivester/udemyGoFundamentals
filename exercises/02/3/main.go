package main

import (
	"fmt"
)

func main() {
	myNums := []int{2, 5, 9, 1, 20, 0}

	fmt.Println(findMax(myNums...))
	fmt.Println(findMax(1, 3))
	fmt.Println(findMax())
}

func findMax(nums ...int) int {
	var max int

	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
}
