package main

import "fmt"

type stru struct {
	A int
} 

func main() {
	//m := make([]int, 1, 10)

	//m := make(map[int]int, 1)
	//m := new(stru)
	//m.A = 2
	//fmt.Printf("%p", &m)
	m := []int{1,2,3,4}
	m[1] = 1
	fmt.Println(m)
	//m = append(m, 1,2,3,4)

	//fmt.Printf("%p", m)
	fmt.Println()
	testPoint(&m)
	fmt.Println(m)

}
func testPoint(a *[]int){
	for x, _ := range *a{
		fmt.Println(x)
	}
	*a = append(*a, 4)
	fmt.Println(a)
}

func testvalue(a []int){
	a = append(a, 4)
	fmt.Println(a)
}

func testMap(m map[int]int)  {
	fmt.Printf("%p", &m)
	m[2] = 2
	fmt.Printf("%p", &m)

	fmt.Println(m)
}

func testStruct(m *stru)  {
	fmt.Printf("%p", m)
	m.A = 1
	fmt.Println(m)
}