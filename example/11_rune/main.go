package main

import (
	"fmt"
)

func main() {
	for i := 500; i < 503; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
