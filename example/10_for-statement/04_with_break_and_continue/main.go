package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		if i > 6 {
			break
		}
		fmt.Println(i)
	}
}
