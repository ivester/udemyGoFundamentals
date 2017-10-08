package main

import (
	"fmt"
)

func main() {
	var firstNum int
	var secondNum int
	var remainder int

	fmt.Println("Please enter two numbers:")

	fmt.Scan(&firstNum)
	fmt.Scan(&secondNum)

	if firstNum > secondNum {
		remainder = firstNum % secondNum
	} else {
		remainder = secondNum % firstNum
	}

	fmt.Println("The remainder of the larger number devided by the smaller number is:", remainder)
}
