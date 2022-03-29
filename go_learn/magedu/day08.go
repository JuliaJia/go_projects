package magedu

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	defautCap = 1024
)

type Ts func() interface{}

type Queue struct {
	elements []interface{}
	limit    int
	locker   sync.Mutex
}

type Pool struct {
	worker  int
	tasks   *Queue
	events  chan struct{}
	Results chan interface{}
	wg      sync.WaitGroup
}

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
func FileLineCount(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			fmt.Printf("%s has %d code lines.\n", filepath.Base(path), fileLine(path))
		}
		return nil
	})
}

func FileLineTotalCount(path string) {
	count := 0
	var wg sync.WaitGroup
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			wg.Add(1)
			go func() {
				count += fileLine(path)
				wg.Done()
			}()
		}
		return nil
	})
	wg.Wait()
	fmt.Printf("%s has %d code lines.\n", path, count)
}

func FileLineTotalCountChannel(path string) {
	total := 0
	var wg sync.WaitGroup
	var wgTotal sync.WaitGroup
	channel := make(chan int, 1000)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			wg.Add(1)
			go func() {
				count := fileLine(path)
				channel <- count
				wg.Done()
			}()
		}
		return nil
	})
	wgTotal.Add(1)
	go func() {
		for cnt := range channel {
			total += cnt
		}
		wgTotal.Done()
	}()
	wg.Wait()
	close(channel)
	wgTotal.Wait()
	fmt.Printf("%s has %d code lines.\n", path, total)
}

func ChannelType(chanType string, chanNum int) (<-chan int, chan<- int) {
	channel := make(chan int, chanNum)
	if chanType == "RO" {
		var readonly <-chan int
		readonly = channel
		return readonly, nil
	} else if chanType == "WO" {
		var writeonly chan<- int
		writeonly = channel
		return nil, writeonly
	}
	return nil, nil
}

func IntChannelTypeUse(channel chan int, chanType string) (chan int, <-chan int, chan<- int) {
	if chanType == "RO" {
		var readonly <-chan int
		readonly = channel
		fmt.Println(<-readonly)
		return channel, readonly, nil
	} else if chanType == "WO" {
		var writeonly chan<- int
		writeonly = channel
		writeonly <- 1
		return channel, nil, writeonly
	}
	return channel, nil, nil
}
func RunTime1() {
	runtime.Gosched()

	wg := new(sync.WaitGroup)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			for ch := 'A'; ch < 'Z'; ch++ {
				fmt.Printf("%d: %c\n", i, ch)
				runtime.Gosched()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func RunTime2() {
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOMAXPROCS(1))
	fmt.Println(runtime.GOMAXPROCS(1))
	fmt.Println(runtime.GOMAXPROCS(1))
}

func NewQueue(limit int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, defautCap),
		limit:    limit,
	}
}

func (q *Queue) PutQueue1(e interface{}) error {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.limit != -1 && len(q.elements) >= q.limit {
		return fmt.Errorf("Queue is limit %d", q.limit)
	}
	q.elements = append(q.elements, e)
	return nil
}

func (q *Queue) TopQueue1() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("Queue is empty!")
	}
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e, nil
}

func (q *Queue) LenQueue() int {
	return q.limit
}

func NewPool(worker int) *Pool {
	return &Pool{
		worker:  worker,
		tasks:   NewQueue(-1),
		events:  make(chan struct{}, math.MaxInt32),
		Results: make(chan interface{}, worker*3),
	}
}

func (p *Pool) AddTask(task Ts) {
	p.tasks.PutQueue1(task)
	p.events <- struct{}{}

}

func (p *Pool) Start() {
	for i := 0; i < p.worker; i++ {
		p.wg.Add(1)
		go func(i int) {
			for range p.events {
				e, err := p.tasks.TopQueue1()
				if err != nil {
					continue
				}
				log.Printf("Worker %d run task", i)
				if task, ok := e.(Ts); ok {
					p.Results <- task()
				}
			}
			p.wg.Done()
		}(i)
	}
}

func (p *Pool) Wait() {
	close(p.events)
	p.wg.Wait()
	close(p.Results)
}
