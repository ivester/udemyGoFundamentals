package main

import (
	"fmt"
)

func main() {
	// slice
	primes := []int{2, 3, 5, 7, 9, 2147483647}

	// map
	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87,
	}

	fmt.Println(primes, noteFrequency)
}
