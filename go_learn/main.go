package main

import (
	"fmt"
	"go_learn/string_func"
)

func main() {
	s := []byte{'a', 'b', 'c'}
	str := string_func.GetStringBySlice(s)
	fmt.Println(str)
}
