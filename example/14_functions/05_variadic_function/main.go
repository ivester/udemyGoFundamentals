package main

import (
	"fmt"
)

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println("The average equals:", n)
}

func average(sf ...float64) float64 {
	fmt.Println("Arguments:", sf)
	fmt.Printf("%v %T \n", "Type of Arguments:", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
