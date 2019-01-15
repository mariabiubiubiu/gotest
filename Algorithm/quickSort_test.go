package main

import (
	"testing"
)

//func ExampleQuickSort(){
//	aint := []int{8, 23, 1, 4, 9, 32, 2, 72}
//	quickSort(aint, 0, len(aint)-1)
//	fmt.Println(aint)
//}

func TestA(t *testing.T){
	aint := []int{8, 23, 1, 4, 9, 32, 2, 72}
	quickSort(aint, 0, len(aint)-1)
	t.Log(aint)
}

