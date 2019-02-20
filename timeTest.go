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


	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	fmt.Println("timeNumber:", timeNumber)




}
