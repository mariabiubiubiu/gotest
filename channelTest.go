package main

import "fmt"

var messages chan int
var signals chan int

func main() {
	messages = make(chan int, 1)
	messages <- 1

	select {
	case msg := <-messages:
		fmt.Println("aaa", msg)
	case sig := <-signals:
		fmt.Println("bbb", sig)
	default:
		fmt.Println("ccc")
	}
	defer close(messages)
	defer close(signals)
}
