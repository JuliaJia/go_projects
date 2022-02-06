package magedu

import "fmt"

func Vars1() {
	var name string = "Ryoma"
	var zeroString string
	var typeString = "Ryoma"
	shortString := "Ryoma"
	fmt.Println(name, zeroString, typeString, shortString)
}

func Vars2() {
	var (
		name string = "Ryoma"
		msg         = "hello world"
		desc string
	)
	fmt.Println(name, msg, desc)
}

func Vars3() {
	var name string = "123"
	fmt.Println(name)
	name = "aaa"
	fmt.Println(name)
	{
		var name string = "func aaaaaaa"
		fmt.Println(name)
	}
	fmt.Println(name)
}
