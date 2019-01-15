package main

import "fmt"

func quickSort(list []int, left, right int) {
	if left > right{
		return
	}
	i := left
	j := right
	for i != j{
		for list[left] < list[j]{
			j--
		}
		for list[left] > list[i]{
			i++
		}
		if i < j {
			list[i], list[j] = list[j], list[i]
		}
	}
	list[left], list[i] = list[i], list[left]
	quickSort(list, left, i-1)
	quickSort(list, i+1, right)
}

func main() {
	aint := []int{8, 23, 1, 4, 9, 32, 2, 72}
	quickSort(aint, 0, len(aint)-1)
	fmt.Println(aint)
}