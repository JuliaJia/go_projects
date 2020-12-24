package main

import (
	"fmt"
)

func main() {
	const (
		age int = 30
		age2 = iota + age
		age3
		age4
		age5
		tiancai string = "贾璐"
		bobo bool = true
		pi float64 = 3.1415926
		
	)

	fmt.Printf("%s is tiancai!This is %t,like pi is %.7f!He is %d years old this year!\n",tiancai,bobo,pi,age)
	fmt.Printf("age2 = %d,age3 = %d,age4 = %d,age5 = %d\n",age2,age3,age4,age5)
}