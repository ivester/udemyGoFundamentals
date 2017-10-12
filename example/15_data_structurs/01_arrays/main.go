package main

import (
	"fmt"
)

func main() {
	// Number inside the square brackets signals an array
	// Empty brackets signal a slice
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	// Going from letter A 58 letters up
	for i := 65; i < 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
