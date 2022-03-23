package main

import (
	"fmt"
	"go_learn/magedu"
	"strings"
)

func Add(a float64, b float64) float64 {
	return a / b
}

func NgFlow(count float64, dict map[int]float64) map[int]float64 {
	new_dict := dict
	new_value := float64(0)
	value := float64(0)
	for k := range dict {
		value = dict[k] / float64(k) / count
		fmt.Println(value)
		new_value += value
		new_dict[k] = value
		value = float64(0)
	}
	for k := range dict {
		new_dict[k] = new_dict[k] / new_value
	}
	fmt.Println(new_dict)
	return new_dict
}

func main() {
	//magedu.Vars1()
	//magedu.Vars2()
	//magedu.Vars3()
	//magedu.Const1()
	//fmt.Println(magedu.PackageMsg)
	//magedu.Print1()
	//magedu.Bool1()
	//magedu.Bool2()
	//magedu.Int1()
	//magedu.Byte1()
	//magedu.Float1()
	//magedu.String1()
	//magedu.Pointer1()
	//magedu.Scan1()
	//magedu.If1()
	//magedu.For3()
	//magedu.Goto2()
	//magedu.Array2()
	//magedu.Slice2()
	//magedu.Queue1()
	//magedu.Stack1()
	//magedu.Sort2()
	//magedu.Map2()
	//magedu.Unicode1()
	//magedu.String2()
	//magedu.Func4()
	//magedu.Todolist()
	//magedu.Defer1()
	//learntools.PrintNow()
	//test01.Test1()
	//learntools.Args1()
	//test01.Test1()
	//test01.Test2()
	//test01.TestLog()
	//test01.TestTime()
	//test01.TestBase64()
	//test01.TestHex()
	//test01.TestHash()
	//test01.TestSha()
	//test01.TestCmd()
	//fmt.Println(Add(1, 2))
	//dict := make(map[int]float64)
	//dict[2] = 0.06
	//dict[3] = 0.14
	//dict[7] = 0.80
	//new_dict := NgFlow(12, dict)
	//fmt.Printf("%#v\n", new_dict)
	//magedu.TaskType()
	//magedu.File1()
	//magedu.FileMd5("password.txt")
	//fmt.Println(magedu.FileIsExists("1"))
	//magedu.File4()
	//magedu.AppendLog()
	magedu.Dir2("magedu", func(path string) bool {
		return strings.HasSuffix(path, ".go")
	}, func(path string) {
		fmt.Println("file: ", path)
		fmt.Println("contentx: ")
		fmt.Println(magedu.ReadFile(path))
	})
}
