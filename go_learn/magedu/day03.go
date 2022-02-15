package magedu

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	taskid        = "id"
	name          = "name"
	startTime     = "start_time"
	endTime       = "end_time"
	status        = "status"
	user          = "user"
	statusNew     = "未执行"
	statusCompele = "完成"
)

var todolist = make([]map[string]string, 0)
var id = 0

func Func1() {
	calc := func(n int, m int, callback func(n1, n2 int) int) int {
		rt := callback(n, m)
		if rt >= 0 && rt <= 100 {
			return rt
		}
		return -1
	}
	sum_value := calc(50, 30, func(n1, n2 int) int {
		return n1 + n2
	})
	fmt.Println(sum_value)
	rt_value := calc(50, 30, func(n1, n2 int) int {
		return n1 * n2
	})
	fmt.Println(rt_value)
}

func Func2() {
	func() {
		fmt.Println("我只执行一次！")
	}()
}

func Func3() {
	addbase := func(base int) func(int) int {
		return func(n int) int {
			return n + base
		}
	}
	add01 := addbase(1)
	fmt.Println(add01(5))
	add10 := addbase(10)
	fmt.Println(add10(3))
}

func Func4() {
	stats := [][]int{{'A', 3}, {'B', 2}, {'D', 2}, {'C', 4}}
	sort.Slice(stats, func(i, j int) bool { return (stats[i][1] > stats[j][1]) })
	fmt.Println(stats)
	index := sort.Search(len(stats), func(i int) bool { return stats[i][1] <= 2 })
	fmt.Println(index)
	sort.Slice(stats, func(i, j int) bool { return (stats[i][1] < stats[j][1]) })
	fmt.Println(stats)
	index = sort.Search(len(stats), func(i int) bool { return stats[i][1] >= 2 })
	fmt.Println(index)
}
func newTask() map[string]string {
	task := make(map[string]string)
	id++
	task[taskid] = strconv.Itoa(id)
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	task[startTime] = ""
	task[name] = ""
	return task
}
func input(str string) string {
	var text string
	fmt.Print(str)
	fmt.Scan(&text)
	return text
}

func add() {
	task := newTask()
	fmt.Println("请输入任务信息：")
	task[name] = input("任务名：")
	task[startTime] = input("开始时间：")
	task[user] = input("负责人：")
	todolist = append(todolist, task)
}

func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID:", task[taskid])
	fmt.Println("任务名:", task[name])
	fmt.Println("开始时间:", task[startTime])
	fmt.Println("完成时间:", task[endTime])
	fmt.Println("任务负责人:", task[user])
	fmt.Println(strings.Repeat("-", 20))

}

func printTaskAll() {
	for _, todo := range todolist {
		printTask(todo)
	}
}

func query() {
	text := input("请输入你的查询信息：")
	for _, todo := range todolist {
		flag := strings.Contains(todo[name], text)
		if flag {
			printTask(todo)
		}
	}
}

func modify() {
	taskID := input("请输入你想要修改的任务ID：")
	flag := false
	for _, todo := range todolist {
		if todo[taskid] == taskID {
			newTaskName := input("新的任务名：")
			todo[name] = newTaskName
			printTask(todo)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("你输入的任务ID不存在！")
	}

}

func taskDelete() {
	taskID := input("请输入你想删除的任务ID：")
	flag := false
	for index, todo := range todolist {
		if todo[taskid] == taskID {
			copy(todolist[index:], todolist[index+1:])
			flag = true
			todolist = todolist[:len(todolist)-1]
			printTask(todo)
			break
		}
	}
	if !flag {
		fmt.Println("你输入的任务ID不存在！")
	}
}

func taskDone() {
	taskID := input("请输入完成的任务ID：")
	flag := false
	for _, todo := range todolist {
		if todo[taskid] == taskID {
			taskEndTime := input("请输入完成时间：")
			todo[endTime] = taskEndTime
			printTask(todo)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("你输入的任务ID不存在！")
	}
}

func Todolist() {
	methods := map[string]func(){
		"add":    add,
		"query":  query,
		"modify": modify,
		"all":    printTaskAll,
		"delete": taskDelete,
		"done":   taskDone,
	}
	for {
		text := input("请输入操作(add/query/modify/delete/all/done/exit):")
		if text == "exit" {
			break
		}
		method, ok := methods[text]
		if ok {
			method()
		} else {
			fmt.Println("输入指令不正确！")
		}

	}

}

func Defer1() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
}

func Panic1() (err error) {
	defer func() {
		fmt.Println("defer")
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("%s", panicErr)
		}

	}()
	fmt.Println("before")
	panic("自定义panic")
	fmt.Println("after")
	return
}

func Panic2() {
	fmt.Println("before main")
	err := Panic1()
	fmt.Println("after main", err)
}
