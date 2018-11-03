//优雅一点
package main

import (
	"runtime"
	"sync"
	"time"
)

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}

func main() {
	pool := New(100)
	println(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	println(runtime.NumGoroutine())
}

//不优雅
//var wg *sync.WaitGroup
//
//func main() {
//	for i := 0 ;i < 10000; i++{
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//		}()
//	}
//	wg.Done()
//}
