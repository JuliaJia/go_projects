package main

import (
	"fmt"
)

func main() {
	var (
		a bool = true
		b bool = 1 == 1
		c bool = 1 != 1
		d bool = !b
		e bool = a && b
		f bool = b || c
	)
	fmt.Printf("a = %t,b = %t,c = %t,d = %t,e = %t,f = %t\n",a,b,c,d,e,f)

}
