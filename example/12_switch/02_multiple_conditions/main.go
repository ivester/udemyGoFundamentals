package main

import (
	"fmt"
)

func main() {
	friend := "Ben"

	switch {
	case friend == "Franz":
		fmt.Println("Hi Franz")
	case friend == "Ben", friend == "Brian":
		fmt.Println("Hi Ben")
	default:
		fmt.Println("Where are my friends?")
	}
}
