package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("a's memory address in dec - %d\n", &a)
}
