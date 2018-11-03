package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println('a')
			default:
				fmt.Println("run")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("stop")
	stop <- true
	fmt.Println("ok")
}
