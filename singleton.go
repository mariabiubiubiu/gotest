package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singeton struct {
}

var (
	instance    *singeton
	initialized uint32
	m           sync.Mutex
)

func Instance() *singeton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	m.Lock()
	defer m.Unlock()
	if instance == nil {
		atomic.StoreUint32(&initialized, 1)
		instance = &singeton{}
	}
	return instance
}

func main() {
	once := &sync.Once{}
	once.Do(func() {
		instance = &singeton{}
	})
	fmt.Println(instance)
}
