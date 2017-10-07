package main

import "net/http"
import "io/ioutil"
import "fmt"

func main() {
	res, _ := http.Get("https://www.panter.ch/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
