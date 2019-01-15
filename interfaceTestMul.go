package main

import "fmt"

type duck1 interface {
	get()
	set(string2 string)
}

type instance1 struct {
	name string
}

func (this *instance1)get(){
	fmt.Println(this.name)
}

func (this *instance1)set(name string){
	this.name = name
}

func (this *instance1)delete(){
	this.name = ""
}

func getDuck(duck12 duck1){
	fmt.Println(duck12)
}

func main() {
	a := new(instance1)
	getDuck(a)
}