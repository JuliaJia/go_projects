package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	
	ch <- 123
	time.Sleep(time.Second)
}
