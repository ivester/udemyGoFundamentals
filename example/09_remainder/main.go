package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 5; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
