package main

import (
	"fmt"
)

// Multiple initialization
const (
	A = iota // 0
	B        // 1
	C        // 2
)

// Multiple initialization
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10) = 2^10 = 1'024
	MB = 1 << (iota * 10) // 2 << (2 * 10) = 2^20 = 1'048'576
)

func main() {
	fmt.Printf("%v - %v - %v \n", A, B, C) // 0 - 1 - 2
	fmt.Printf("%b\t", KB)                 // 10000000000
	fmt.Printf("%d\n", KB)                 // 1024
	fmt.Printf("%b\t", MB)                 // 100000000000000000000
	fmt.Printf("%d\n", MB)                 // 1048576
}
