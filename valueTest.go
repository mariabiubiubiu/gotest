package main

import "fmt"

func main() {
	Testfunc("1",1,2)
}

func Testfunc(a ...interface{}){
	fmt.Println(a)
}
