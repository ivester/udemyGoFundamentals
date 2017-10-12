package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("This anonymous function is immediately called")
	}()
}
