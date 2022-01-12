package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	c := 0
	waitg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		waitg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			defer waitg.Done()
			c++
			fmt.Println(c)
		}()
	}
	waitg.Wait()
}
