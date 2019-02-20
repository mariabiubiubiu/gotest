package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,5)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 5

	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				close(ch)
				break

			}
			time.Sleep(1*time.Second)
		}
	}()
	//for value := range ch{
	//	fmt.Println(value)
	//	fmt.Println("0000")
	//}
	fmt.Println("This line code wont run")
}