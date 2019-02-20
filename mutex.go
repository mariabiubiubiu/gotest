package main

import (
	"fmt"
	"sync"
)

type Mute struct {
	Mu sync.Mutex
	I int
}
var x = 1

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incMu(mu *Mute, wg *sync.WaitGroup){
	mu.Mu.Lock()
	defer mu.Mu.Unlock()
	mu.I += 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	mu := &Mute{sync.Mutex{}, 0}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		incMu(mu, &w)
	}
	w.Wait()
	fmt.Println("final value of x", mu)
}
