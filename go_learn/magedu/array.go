package magedu

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Array1() {
	var name [32]string
	name = [32]string{"aaa", "bbb"}
	fmt.Printf("%s\n", name[1])
	names := [...]string{"123", "456"}
	for _, v := range names {
		fmt.Println(v)
	}
}

func Array2() {
	var d2 = [3][2]int{}

	d2[0] = [2]int{1, 2}
	d2[1] = [2]int{3, 4}
	d2[2] = [2]int{5, 6}

	for _, v1 := range d2 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

}

func Slice1() {
	var names []string
	fmt.Printf("%T\n", names)
	fmt.Printf("%#v\n", names)
	names = []string{}
	fmt.Printf("%#v\n", names)
	names = []string{"天才哥哥", "小猪头弟弟"}
	fmt.Printf("%#v\n", names)
	names = []string{1: "天才哥哥", 100: "小猪头弟弟"}
	fmt.Printf("%#v\n", names)
	names = make([]string, 5)
	fmt.Printf("%#v\n", names)
	names = make([]string, 0, 10)
	fmt.Printf("%#v,%d\n", names, cap(names))
	names = []string{"天才哥哥", "小猪头弟弟"}
	fmt.Printf("%#v,%d,%d\n", names, len(names), cap(names))
	names = append(names, "段子王")
	fmt.Printf("%#v,%d,%d\n", names, len(names), cap(names))
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func Slice2() {
	aSlice := []int{1, 2, 3}
	bSlice := []int{4, 5}
	fmt.Printf("%#v,%#v\n", aSlice, bSlice)
	copy(aSlice, bSlice)
	fmt.Printf("%#v,%#v\n", aSlice, bSlice)
	names := [3]string{"天才哥哥", "小猪头弟弟", "段子王"}
	new_names := names[2:3]
	fmt.Printf("%T,%#v\n", new_names, new_names)
	new_names[0] = "段子王哥哥"
	fmt.Printf("%#v,%#v\n", names, new_names)
	new_names = append(new_names, "老干部", "老干部2", "老干部3")
	fmt.Printf("%#v,%#v\n", cap(names), cap(new_names))
	fmt.Printf("%#v,%#v\n", names, new_names)
	newNamesChildren := new_names[1:3:4]
	fmt.Printf("%#v\n", newNamesChildren)
	copy(new_names[2:], new_names[3:])
	fmt.Printf("%#v\n", new_names[:len(new_names)-1])
}

func Queue1() {
	queue := []string{}
	queue = append(queue, "a")
	x := queue[0]
	queue = queue[1:]
	fmt.Println("1:", x)
}

func Stack1() {
	stack := []string{}
	stack = append(stack, "a", "b")
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射：", x)
}

func Sort1() {
	nums := []int{3, 2, 1, 6, 90, 7}
	sort.Ints(nums)
	fmt.Println(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
}

func Sort2() {
	nums := []int{3, 2, 1, 6, 90, 7}
	sort.Ints(nums)
	fmt.Println(nums[sort.SearchInts(nums, 80)] == 80)
}

func Map1() {
	var scores map[string]float64
	fmt.Printf("%T,%#v\n", scores, scores)
	scores = map[string]float64{"天才哥哥": 100}
	fmt.Printf("%T,%#v\n", scores, scores)
}

func Map2() {
	scores := make(map[string]float64)
	scores["123"] = 100
	fmt.Printf("%T,%#v\n", scores, scores)
	v, ok := scores["4332"]
	fmt.Println(v, ok)
	v, ok = scores["123"]
	fmt.Println(v, ok)
	scores["天才哥哥"] = 100
	fmt.Printf("%T,%#v\n", scores, scores)
	delete(scores, "123")
	fmt.Printf("%T,%#v\n", scores, scores)
	for k := range scores {
		fmt.Println(scores[k])
	}
	for k, v := range scores {
		fmt.Println(k, v)
	}
}

func Unicode1() {
	ascii := "abc我爱北京天安门"
	fmt.Println(utf8.RuneCountInString(ascii))
	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))
	fmt.Println(ascii)
	fmt.Println(strconv.FormatFloat(3.14544666, 'f', 10, 64))
}

func String2() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.Contains("abc", "a"))
	fmt.Println(strings.Contains("abc", "d"))
	fmt.Println(strings.Count("aabc", "a"))
	fmt.Println(strings.Count("aabc", "aa"))
	fmt.Println(strings.Count("aabc", "d"))
	fmt.Println(strings.Fields("a b\tc\nd\re\ff"))
	fmt.Println(strings.HasPrefix("abc", "ab"))
	fmt.Println(strings.HasPrefix("abc", "bc"))
	fmt.Println(strings.HasSuffix("abc", "ab"))
	fmt.Println(strings.HasSuffix("abc", "bc"))
	fmt.Println(strings.Index("abc", "bc"))
	fmt.Println(strings.Index("abc", "ddd"))
	fmt.Println(strings.Join(strings.Fields("a b\tc\nd\re\ff"), ""))
	fmt.Println(strings.Join(strings.Split("a-b-c-d-e-f-g", "-"), ""))
	fmt.Println(strings.Repeat("*", 10))
	fmt.Println(strings.Replace("aaaabbbbccccdddd", "a", "g", -1))
	fmt.Println(strings.Title("abc"))
	fmt.Println(strings.ToUpper("abcABC"))
	fmt.Println(strings.ToLower("abcABC"))
	fmt.Println(strings.Trim("abcdefabc", "abc"))
	fmt.Println(strings.TrimSpace("  \r\fabc\n\t"))
}
