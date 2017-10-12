package main

import (
	"fmt"
)

func main() {
	// make a slice - Reference Type
	a := make([]string, 1, 25)
	// String Type
	b := "Ives"

	fmt.Println("Slice before change:", a)
	fmt.Println("String before change:", b)

	change(a, b)

	fmt.Println("Slice after change:", a)
	fmt.Println("String after change:", b)
}

func change(x []string, y string) {
	x[0] = "Todd"
	y = "Catherine"
	fmt.Println("Slice inside func after change:", x)
	fmt.Println("String inside func after change:", y)
}
