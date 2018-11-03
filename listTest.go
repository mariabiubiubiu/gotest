package main

import "fmt"

func main() {
	a := []int{1,2,3,4}
	for _, value := range a{
		if value == 2{
			fmt.Println(value)
		}
	}



}
