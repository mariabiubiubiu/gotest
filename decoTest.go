package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func dec(func1 func(s string))func(s string){
	fmt.Println("start")
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(reflect.TypeOf(func1))
	return func1
}

func decNest(func1 func(s string))func(s string){
	fmt.Println("start")
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(reflect.TypeOf(func1))
	return func1
}


func abcd(s string){
	fmt.Println(s)
}

func main() {
	a := dec(abcd)
	a("mu")
}