package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a := time.Now().UnixNano()
	time.Sleep(190*time.Millisecond)
	time.Sleep(1*time.Millisecond)
	final := time.Now()
	b := time.Now().UnixNano()
	fmt.Println(1e9)

	fmt.Println(final.Sub(start).Seconds(), "seconds")
	fmt.Println(b -a )

	fmt.Println(time.Now())
}
