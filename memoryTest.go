package main

import (
	"fmt"
	"sync"
)

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("hey hhhhhhhh")
		wg.Done()
	}()
	wg.Wait()

	print(a)
}
