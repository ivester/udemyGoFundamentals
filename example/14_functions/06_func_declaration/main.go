package main

import (
	"fmt"
)

func main() {
	greet("Catherine")
}

// Func declaration
func greet(name string) {
	fmt.Println("Hello", name)
}
