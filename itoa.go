package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexWaiterShift = iota
)

func main() {
	fmt.Println(mutexLocked, mutexWoken, mutexWaiterShift)
}