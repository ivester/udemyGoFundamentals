package main

import (
	"fmt"
)

func main() {
	condition := false

	if food := "bread"; condition == true {
		fmt.Println(food, ", condition true")
	} else if condition == true {
		fmt.Println(food, ", condition true")
	} else {
		fmt.Println(food, ", condition false")
	}

	// food will not be available here
	// fmt.Println(food)
}
