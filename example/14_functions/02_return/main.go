package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Catherine ", "Ebneter"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
