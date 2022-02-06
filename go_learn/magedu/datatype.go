package magedu

import (
	"fmt"
	"strconv"
)

func Bool1() {
	isGirl := false
	fmt.Printf("%T,%#v", isGirl, isGirl)
}

func Bool2() {
	a := true
	b := false
	c := true
	d := false

	fmt.Println("a && b", a && b)
	fmt.Println("a && c", a && c)
	fmt.Println("b && c", b && c)
	fmt.Println("b && d", b && d)

	fmt.Println("a || b", a || b)
	fmt.Println("a || c", a || c)
	fmt.Println("b || c", b || c)
	fmt.Println("b || d", b || d)

	fmt.Println("!a", !a)
	fmt.Println("!b", !b)

	fmt.Println("b == c", b == c)
	fmt.Println("b == d", b == d)
	fmt.Println("b != c", b != c)
	fmt.Println("b != d", b != d)

	fmt.Printf("%t, %t", a, b)
}

func Int1() {
	var age8 int8 = 31
	fmt.Printf("%T,%#v,%d\n", age8, age8, age8)
	var age = 31
	fmt.Printf("%T,%#v,%d\n", age, age, age)
	var age1 int
	age1 = int(age8) + age
	fmt.Printf("%T,%#v,%d\n", age1, age1, age1)

}

func Byte1() {
	var chara byte = 'a'
	var aint byte = 64
	var unicodePoint rune = '中'
	fmt.Printf("%T,%#v,%c\n", chara, chara, chara)
	fmt.Printf("%T,%#v,%c\n", aint, aint, aint)
	fmt.Printf("%d,%b,%o,%x,%U,%c,%c", chara, 15, 15, 15, unicodePoint, chara, aint)
}

func Float1() {
	var height float32 = 1.68
	var heightType = 1.68
	fmt.Printf("%T,%#v,%f\n", height, height, height)
	fmt.Printf("%T,%#v,%f\n", heightType, heightType, heightType)
}

func String1() {
	msg := "中国"
	fmt.Println(msg[0])
	fmt.Println(len(msg))
	strint := "3"
	floatstr := 2.2
	num, _ := strconv.Atoi(strint)
	fmt.Printf("%T,%#v,%d\n", num, num, num)
	floatnum := strconv.FormatFloat(floatstr, 'f', 10, 64)
	vv, _ := strconv.ParseFloat(floatnum, 64)
	fmt.Printf("%T,%#v,%s\n", floatnum, floatnum, floatnum)
	fmt.Printf("%T,%#v,%f\n", vv, vv, vv)

}
