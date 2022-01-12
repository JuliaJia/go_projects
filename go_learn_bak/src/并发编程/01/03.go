package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	intch := make(chan int, 1)
	strch := make(chan string, 1)
	intch <- 1
	strch <- "hello"
	/*
		1.select 如果没有case会阻塞
		2.select 监测chan数据流向
		3.case 必须为IO操作
		4.select 对应异步时间处理
		5.select 超时处理
		6.如果多个case都满足条件，会随机选择其中之一来执行
	*/
	select {
	case value := <-intch:
		fmt.Println(value)
	case value := <-strch:
		fmt.Println(value)
	case <-time.After(time.Second * 5):
		fmt.Println("超时")
	}

}
