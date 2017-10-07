package main

import (
	"fmt"
)

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Ben", "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Falko":
		fmt.Println("Wassup Falko")
	default:
		fmt.Println("Wassup Man")
	}
}
