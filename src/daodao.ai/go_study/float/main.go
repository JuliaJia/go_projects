package main

import (
	"fmt"
)

func main() {
	var (
		a float32 = 0.12345678
		b float64 = 0.12345678
	)
	fmt.Printf("a = %.64f,b = %.64f\n",a,b)
}