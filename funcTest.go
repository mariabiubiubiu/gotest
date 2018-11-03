package main

import "fmt"

func main() {
	//fmt.Println(Ad(1,2))
	//fmt.Println(Add(1,2))
	var a = []interface{}{1, "aaa", 2}
	Print(a)
	Print(a...)
	Print(Inc())
	//因为是闭包，在for迭代语句中，每个defer语句延迟执行的函数引用的都是同一个i迭代变量，在循环结束后这个变量的值为3，因此最终输出的都是3。
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}

}

func Add(a, b int) int {
	return a + b
}

var Ad = func(a, b int) int {
	return a + b
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

//第一个Print调用时传入的参数是a...，等价于直接调用Print(123, "abc")。
//第二个Print调用传入的是未解包的a，等价于直接调用Print([]interface{}{123, "abc"})。

//...相当于解包

//不仅函数的参数可以有名字，也可以给函数的返回值命名：

func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

//如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值：
