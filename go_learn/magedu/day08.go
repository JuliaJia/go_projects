package magedu

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func PrintChars(prefix string, wait *sync.WaitGroup) {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%s: %c\n", prefix, i)
		time.Sleep(time.Millisecond)
	}
	wait.Done()
}

func MainChars(prefix string) {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%s: %c\n", prefix, i)
		time.Sleep(time.Millisecond)
	}
}

func GoRoutine1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintChars("go_learn", &wg)
	MainChars("main")
	fmt.Println("Wait")
	wg.Wait()
	fmt.Println("Over")
}

func GoRoutine2(num int) {
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	//time.Sleep(time.Second)
	fmt.Println("Wait")
	wg.Wait()
	fmt.Println("Over")
}
func GoRoutine3() {
	var a, b = 10000, 10000
	var count = 100
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)
	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if a > money {
				lock.Lock()
				a -= money
				b += money
				lock.Unlock()
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if b > money {
				lock.Lock()
				b -= money
				a += money
				lock.Unlock()
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
}

func Goroutine4(count int) {
	var wg sync.WaitGroup
	var counter int64
	counter = 0
	var ceil = 10000
	for i := 0; i < count; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < ceil; i++ {
				atomic.AddInt64(&counter, 1)
				time.Sleep(time.Millisecond)
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < ceil; i++ {
				atomic.AddInt64(&counter, -1)
				time.Sleep(time.Millisecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}

func Chan1(count int) {
	var channel chan int64
	var counter int64
	var wg sync.WaitGroup
	counter = 0
	channel = make(chan int64)
	wg.Add(2)
	go func() {
		//time.Sleep(time.Millisecond)
		for i := 0; i < count; i++ {
			atomic.AddInt64(&counter, 1)
			channel <- counter
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < count; i++ {
			num := <-channel
			fmt.Println(num)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Chan2() {
	channel := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 5)
		channel <- struct{}{}
	}()
	fmt.Println("before")
	fmt.Println(<-channel)
	fmt.Println("over")
}

func fileLine(path string) int {
	cnt := 0
	file, err := os.Open(path)
	if err != nil {
		return cnt
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		ctx, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if strings.TrimSpace(string(ctx)) == "" || strings.HasPrefix(strings.TrimSpace(string(ctx)), "//") {
			continue
		}
		cnt++
	}
	return cnt
}
func FilePathList(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			fmt.Printf("%s has %d code lines.\n", filepath.Base(path), fileLine(path))
		}
		return nil
	})
}
