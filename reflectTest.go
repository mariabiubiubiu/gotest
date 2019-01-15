package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func (this *Student) SetName(name string) {
	this.Name = name
	fmt.Printf("set name %s\n",this.Name )
}

func (this *Student) SetAge(age int) {
	this.Age = age
	fmt.Printf("set age %d\n",age )
}

func (this *Student) String() string {
	fmt.Printf("this is %s\n",this.Name)
	return this.Name
}

func main() {
	stu1 := &Student{Name:"wd",Age:22}
	val := reflect.ValueOf(stu1)       //获取Value类型，也可以使用reflect.ValueOf(&stu1).Elem()
	val.MethodByName("String").Call(nil)  //调用String方法

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	val.MethodByName("SetAge").Call(params)  //通过名称调用方法

	params[0] = reflect.ValueOf("jack")
	val.Method(1).Call(params)    //通过方法索引调用

	fmt.Println(stu1.Name,stu1.Age)
}
//this is wd
//set age 18
//set name jack
//jack 18