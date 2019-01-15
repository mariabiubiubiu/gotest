package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {
	//s := "hello"
	//fmt.Println(s[2:5])
	//for _, value := range s{
	//	fmt.Println(string(value))
	//}
	s := "hello, world"
	//hello := s[:5]
	//world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5

	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �界abc

	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}

	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}
}
