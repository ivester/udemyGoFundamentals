package main

import (
	"fmt"
)

func main() {
	myGreetings := make(map[string]string)

	myGreetings["german"] = "Guten Tag"
	myGreetings["english"] = "Good morning"

	//“comma ok” idiom
	val, ok := myGreetings["english"]

	fmt.Println(val, ok)
}
