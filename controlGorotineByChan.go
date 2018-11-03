package main

var ch chan int

func main() {
	ch = make(chan int, 10)
	for i := 0; i < 10000; i++ {
		ch <- 1 //create
		go func() {
			<-ch //consume
		}()
	}
}
