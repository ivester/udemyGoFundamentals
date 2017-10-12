package main

import (
	"fmt"
)

func main() {
	// Func expression
	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet("Catherine")
}
