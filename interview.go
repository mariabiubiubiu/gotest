package main

import (
	"fmt"
)


type testStru struct {
	A int
}

//func main()  {
//	test := new(testStru)
//	//test := testStru{1}
//	test.A = 1
//	fmt.Printf("%p\n", &test)
//	testStructPoint(test)
//	fmt.Println(test)
//
//}

func testStructPoint(a *testStru)  {
	fmt.Printf("%p\n", a)
	a.A = 2
	fmt.Println(a)
}

func testStructValue(a testStru)  {
	fmt.Printf("%p\n", &a)
	a.A = 2
	fmt.Println(a)
}

func main() {
	m := make(map[int]int, 1)
	fmt.Printf("%p\n", m)
	m[1] = 1
	testMapValue(m)
	fmt.Println(m)

}

func testMapValue(m map[int]int)  {
	fmt.Printf("%p\n", m)
	m[2] = 2
	fmt.Println(m)
}


func testMapPoint(m map[int]int)  {
	m[2] = 2
	fmt.Println(m)
}


func testSliceValue(a []int){
	a[1] = 5
	a = append(a, 4)
	fmt.Println(a)
}

func testSlicePoint(a *[]int){
	*a = append(*a, 4)
	fmt.Println(a)
}
