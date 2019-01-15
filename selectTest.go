package main

import "fmt"

func main()  {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	go func() {
		chan1 <- 1
	}()
	select {
		case <-chan1:
			fmt.Println("1")
		case <-chan2:
			fmt.Println("2")
	}

	fmt.Println("3333")
}
