package main

import (
	"fmt"
)

func main() {
	fmt.Println(half(4))
	fmt.Println(half(3))
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
