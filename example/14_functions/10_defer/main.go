package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1. - Should be printed first but will be printed second")
	fmt.Println("2. - Should be printed second but will be printed first")
}
