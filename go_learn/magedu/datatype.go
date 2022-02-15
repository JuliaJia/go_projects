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

func Pointer1() {
	var (
		pointerInt    *int
		pointerString *string
	)

	fmt.Printf("%T,%#v\n", pointerInt, pointerInt)
	fmt.Printf("%T,%#v\n", pointerString, pointerString)
	age := 32
	age2 := age
	fmt.Printf("%T,%#v\n", &age, &age)
	fmt.Printf("%T,%#v\n", &age2, &age2)
	pointerInt = &age
	fmt.Println(*pointerInt)
	*pointerInt = 33000
	fmt.Println(age2, age)
	pointerString = new(string)
	*pointerString = "123"
	pp := &pointerString
	fmt.Println(**pp, *pointerString)
	**pp = "321"
	fmt.Println(*pointerString)
}
func Scan1() {
	name := ""
	age := 0
	msg := ""
	fmt.Println("请输入你的名字：")
	fmt.Scan(&name)
	fmt.Println("你的名字是：", name)
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	fmt.Println("你的年龄是：", age)
	fmt.Println("请输入你的信息：")
	fmt.Scan(&msg)
	fmt.Println("你的信息是：", msg)
}
func If1() {
	fmt.Println("去给本女王买十个包子！看到卖西瓜的就买一个！")
	var y string
	fmt.Println("有没有卖西瓜的：")
	fmt.Scan(&y)
	if y == "yes" {
		fmt.Println("老板，来一个西瓜，又大又甜的那种！")
		fmt.Println("十个包子一个大又甜西瓜给wuli宝贝小公主买回来啦！")
	} else if y == "no" {
		fmt.Println("十个包子给wuli宝贝小公主买回来啦！")
	} else {
		fmt.Println("输入错误！")
	}
}

func For1() {
	var sum = 0
	for index := 1; index <= 100; index++ {
		sum += index
	}
	fmt.Println(sum)
}

func For2() {
	var sum = 0
	var index = 1
	for index <= 1000 {

		if index == 101 {
			break
		}
		sum += index
		index++
	}
	fmt.Println(sum)
}
func For3() {
	var str = "我爱毛主席"
	for _, v := range str {
		fmt.Println(string(v))
	}
	for _, v := range str {
		fmt.Printf("%q\n", v)
	}
}
func Goto1() {
	fmt.Println("start")
	goto End
	fmt.Println("1")
End:
	fmt.Println("End")
}
func Goto2() {
	fmt.Println("Start")
	var index = 1
CC:
	fmt.Println(index)
	index++
	if index > 100 {
		goto End
	}
	goto CC
End:
	fmt.Println("End")
}
