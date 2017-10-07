package main

import (
	"fmt"
)

var z = "Ives"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	x := "Ebi"

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)

	{
		fmt.Printf("%v \n", x)
		fmt.Printf("%v \n", z)
	}
}
