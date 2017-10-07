package main

import (
	"fmt"
)

func zero(z int) {
	fmt.Printf("%p\n", &z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	zero(x)
	fmt.Println(x)
}
