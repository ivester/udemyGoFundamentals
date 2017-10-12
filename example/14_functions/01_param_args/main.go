package main

import (
	"fmt"
)

func main() {
	greet("Catherine", "Ebneter")
}

func greet(fname, lname string) {
	fmt.Println("Hello", fname, lname)
}
