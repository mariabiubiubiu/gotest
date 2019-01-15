package main

import (
	"fmt"
)

func main() {
	GO()
	PHP()
	PYTHON()
}

//Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为0了）。才使用Go中引入的Exception处理：defer, panic, recover。

//Go没有异常机制，但有panic/recover模式来处理错误
//Panic可以在任何地方引发，但recover只有在defer调用的函数中有效
func GO() {
	fmt.Println("我是GO，现在没有发生异常，我是正常执行的。")
}

func PHP() {
	// panic一般会导致程序挂掉（除非recover）  然后Go运行时会打印出调用栈
	//但是，关键的一点是，即使函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。panic就是这么简单。抛出个真正意义上的异常。
	panic("我是PHP,我要抛出一个异常了，等下defer会通过recover捕获这个异常，然后正常处理，使后续程序正常运行。")
	fmt.Println("我是PHP里panic后面要打印出的内容。")
}

func PYTHON() {
	fmt.Println("我是PYTHON，没有defer来recover捕获panic的异常，我是不会被正常执行的。")
}