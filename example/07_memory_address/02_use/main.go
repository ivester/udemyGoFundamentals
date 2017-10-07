package main

import (
	"fmt"
)

const metersToYards float64 = 0.09361

func main() {
	var meters float64
	fmt.Print("Enter meters: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meteres is ", yards, " yards.")
}
