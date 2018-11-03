package main

import (
	"fmt"
	"reflect"
)

type mu int

type User struct {
	Name string
	Age  int
}

func (u *User) Hello(a string) {
	fmt.Println(a)
}

func main() {
	u := User{Name: "muweijia", Age: 24}
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}
	var mu int
	mu = 1
	muvalue := reflect.ValueOf(&mu)
	muvalue.Elem().SetInt(1000)
	fmt.Println(mu)

	v := reflect.ValueOf(u)
	hello := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(hello.Call(args))
}
