package main

import "fmt"

func main() {
	x := []int{1,2,3}

	func(arr []int) {
		arr[0] = 0
		fmt.Println(arr) //输出 [0 2 3]
	}(x)

	fmt.Println(x) //输出 [0 2 3]
}