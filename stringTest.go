package main

import "fmt"

func main() {
	a := "abd"
	var b string
	b = a
	a = "aaaaa"
	fmt.Println(b)
}
