package main

import "fmt"


func MultiReturn(a int,b int) (int,int,int) {
	return a,b,a + b
}


func main() {
	var a int
	var b int
	a,b,sum := MultiReturn(3,4)
	fmt.Println(a,"+",b,"=",sum)
}