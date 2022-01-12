package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("协程 完成")
		close(ch)
	}()

	fmt.Println("主函数执行")
	<-ch
	fmt.Println("主函数执行完成")
}
