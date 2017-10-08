package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		// if the switch has no expression it switches on true
		switch {
		case i == 0:
			break
		case i%3 == 0 && i%5 == 0:
			fmt.Print("FizzBuzz, ")
		case i%5 == 0:
			fmt.Print("Fizz, ")
		case i%3 == 0:
			fmt.Print("Buzz, ")
		default:
			fmt.Print(i, ", ")
		}
	}
}
