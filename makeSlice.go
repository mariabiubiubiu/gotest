package main

import (
	"fmt"
	"reflect"
)

type Slice struct {
	len int
	cap int
	array interface{}
}

func NewSlice(i ...interface{})Slice{
	array := reflect.ArrayOf(len(i), reflect.TypeOf(i[0]))
	fmt.Println(array)
	temp := reflect.New(array)
	temp.Pointer()
	for _, value := range i{
		fmt.Println(reflect.ValueOf(value))
		temp.Set(reflect.ValueOf(value))
	}
	return Slice{array:&array}
}

//func (self *Slice)add(i ...int){
//	s := reflect.ArrayOf(self.len + len(i), reflect.TypeOf(i))
//	s = self.array
//}
//
//func (self *Slice)get(i int)  {
//	reflect.ArrayOf(i, int)
//}

func main() {
	a := NewSlice("5","2")
	fmt.Println(a.array)
}