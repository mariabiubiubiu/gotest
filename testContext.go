package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop")
			default:
				fmt.Println("ing")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	cancel()

}
