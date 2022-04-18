package learn

import (
	"fmt"
	"runtime"
	"sync"
)

func NumCpu() {
	fmt.Printf("cpu core num is %d\n", runtime.NumCPU)
}

func NumGoroutine() {
	fmt.Printf("Goroutine  num is %d\n", runtime.NumGoroutine())
}

func SetGOMAXPROCS(num int) {
	runtime.GOMAXPROCS(num)
}

func Mutex01() {
	var mu sync.Mutex
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

type Counter struct {
	mu          sync.Mutex
	count       uint64
	CounterType int
	Name        string
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func Mutex03() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

func Mutex02() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				counter.mu.Lock()
				counter.count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
