
package main

import (
	"fmt"
)

func t(args ...int) {
	fmt.Printf("%p\n", args)
}

func testStu(s ...StudentA){
	//fmt.Println(unsafe.Pointer(&s))
	fmt.Printf("%p\n", s)
}

type StudentA struct {
	name string
}

func main() {
	a := StudentA{"b"}
	//b := []StudentA{a}
	//b := &StudentA{"a"}

	//testStu(b...)
	TestPoint(&a)
	//testStu(b)
	//fmt.Println(unsafe.Pointer(&b))
	//fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", &a)
}


func TestPoint(s *StudentA){
	fmt.Printf("%p\n", s)
}