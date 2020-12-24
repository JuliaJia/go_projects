package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		a string = "贾璐是天才！"
		e string = "我是双引号！\n"
		f string = `我是反引号!\n`
	)
	b := a
	c := "贾璐简直天才到爆表了！"
	fmt.Printf("%s%v%v\n%v%v\n",a,b,c,e,f)

	var (
		str_len int = len(a)
		str string = a + c
	)
	str_split := strings.Split(e,"\n")
	str_contains := strings.Contains(c,"天才")

	fmt.Printf("a's len is %d!\nstr is %v!\ne split is %v\nstr_contains is %v\n",str_len,str,str_split[0],str_contains)

	str_prefix := strings.HasPrefix(a,"贾璐")
	str_suffix := strings.HasSuffix(a,"天才！")
	fmt.Printf("Is a prefix 贾璐？%v\nIs a suffix 天才？%v\n",str_prefix,str_suffix)

	str_index := strings.Index(str,"天才")
	str_lastindex := strings.LastIndex(str,"天才")
	fmt.Printf("天才's index is %v,天才's lastindex is %v\n",str_index,str_lastindex)

	str_join := strings.Join(str_split,"+")
	fmt.Printf("str_split is %v\nstr_join is %v\n",str_split,str_join)

}