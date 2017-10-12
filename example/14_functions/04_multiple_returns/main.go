package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Catherine", "Ebneter"))
}

func greet(fname, lname string) (string, string, string) {
	return fmt.Sprint(fname, " ", lname), fmt.Sprint(" - "), fmt.Sprint(lname, " ", fname)
}
