package main

import "fmt"

type A interface {
	add()
}

type B interface {
	add()
	mul()
}

type AA struct {
}

func (a AA) add() {
	fmt.Println("a")
}

type BB struct {
}

func (b BB) add() {
	fmt.Println("b")
}

func (b BB) mul() {
	fmt.Println("b mul")
}

func main() {
	var a AA

}

func check(bb BB) {

}
