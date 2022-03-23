package magedu

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	id   int
	name string
	addr string
	tel  string
}

type task struct {
	id        int
	name      string
	startTime *time.Time
	endTime   *time.Time
	status    int
	*User
}

type task2 struct {
	id        int
	name      *string
	startTime *time.Time
	endTime   *time.Time
	status    int
	*User
}

type FileFilter func(string) bool
type FileCallback func(string)

func TaskType() {
	var task task
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)
}

func (task *task) SetTaskName(name string) {
	task.name = name
}
func (task *task) GetTaskName() string {
	return task.name
}

func (task task) SetTask1Name(name string) {
	task.name = name
}
func (task task) GetTask1Name() string {
	return task.name
}

func (user *User) SetUserName(name string) {
	user.name = name
}

func (user *User) GetUserName() string {
	return user.name
}

func (task *task2) SetTask2Name(name string) {
	task.name = &name
}
func (task *task2) GetTask2Name() *string {
	return task.name
}

func NewTask(id int, name string, user *User) *task {
	start := time.Now()
	end := start.Add(24 * time.Hour)
	return &task{
		id,
		name,
		&start,
		&end,
		1,
		user,
	}
}
func TaskFuncUse() {
	user := &(User{
		1,
		"塞尔达",
		"",
		"",
	})
	name1 := user.GetUserName()
	task := NewTask(1, "塞尔达传说", user)
	task.User.SetUserName("林克")
	name := task.User.GetUserName()
	fmt.Println("我是" + name + "不是" + name1 + "!")
}
func Task2FuncUse() {
	task := task2{}
	task.SetTask2Name("林克")
	name := task.GetTask2Name()
	fmt.Println("我是" + *name + "!")
}

func Task3FuncUse() {
	task := task{}
	task2 := &task2{}
	methodValue1 := task.SetTaskName
	methodValue2 := task2.SetTask2Name
	methodValue1("123")
	methodValue2("456")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%v\n", *(task2.name))

}

func Task4FuncUse() {
	task1 := task{}
	task2 := &task{}
	method1 := task.SetTask1Name
	method2 := (*task).SetTaskName
	//fmt.Printf("%T\n", method1)

	method1(task1, "test1")
	method1(*task2, "test1")
	fmt.Printf("%#v\n", task1)
	fmt.Printf("%#v\n", task2)
	method2(&task1, "test2")
	method2(task2, "test2")
	fmt.Printf("%#v\n", task1)
	fmt.Printf("%#v\n", task2)

}

func File1() {
	file, err := os.Open("password.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		cxt := make([]byte, 1024)
		n, err := file.Read(cxt)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n)
			fmt.Println(string(cxt[:n]))
		}
	}

}

func File2() {
	file, err := os.OpenFile("password.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		file.Write([]byte("abc123456789"))
	}
}

func File3() {
	var (
		file_inPath  string
		file_outPath string
		new_string   string
		old_string   string
	)
	flag.StringVar(&file_inPath, "i", "", "in_path")
	flag.StringVar(&file_outPath, "o", "", "out_path")
	flag.StringVar(&new_string, "n", "", "new_string")
	flag.StringVar(&old_string, "O", "", "old_string")
	flag.Parse()
	file_out, err := os.OpenFile(file_outPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file_out.Close()
	file_in, err := os.Open(file_inPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file_in.Close()
	br := bufio.NewReader(file_in)
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		str_line := string(line)
		new_line := strings.Replace(str_line, old_string, new_string, -1)
		_, err = file_out.WriteString(new_line + "\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Println("done", index)
		index++
	}
	fmt.Println("FINISH!")
}

func File4() {
	str := "123456"
	new_str := strings.Replace(str, "123", "7788", -1)
	fmt.Println(new_str)

}

func File5() {
	var sleep_time int64
	var file_outPath string
	flag.Int64Var(&sleep_time, "s", 0, "sleep_time")
	flag.StringVar(&file_outPath, "o", "", "file_outPath")
	flag.Parse()
	file_out, err := os.OpenFile(file_outPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file_out.Close()
	file_out.Write([]byte("abc"))
	fmt.Println("sleep")
	time.Sleep(time.Duration(sleep_time) * time.Second)
	file_out.Write([]byte("123"))
}

func FileMd5(filePath string) string {
	flag.StringVar(&filePath, "p", "", "path")
	flag.Parse()
	if name == "" {
		return "请输入文件绝对路径！"
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "打开文件错误！"
	}
	defer file.Close()
	hasher := md5.New()
	ctx := make([]byte, 1024)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		hasher.Write(ctx[:n])
	}
	fmt.Printf("%x\n", hasher.Sum(nil))
	return "123"
}

func AppendLog() {
	logfile, err := os.OpenFile("append.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("appendLog")
}

func FileIsExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func Dir(path string, filter FileFilter, callback FileCallback) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, name := range names {
		fpath := path + "/" + name
		if fileInfo, err := os.Stat(fpath); err == nil {
			if fileInfo.IsDir() {
				Dir(fpath, filter, callback)
			}
			if filter == nil || filter(fpath) {
				if callback != nil {
					callback(fpath)
				}
			}
		}
	}
}

func Dir2(path string, filter FileFilter, callback FileCallback) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	fileInfos, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for _, fileInfo := range fileInfos {
		fpath := path + "/" + fileInfo.Name()
		if fileInfo.IsDir() {
			Dir2(fpath, filter, callback)
		}
		if filter == nil || filter(fpath) {
			if callback != nil {
				callback(fpath)
			}
		}
	}
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	txt := make([]byte, 0, 1024*1024)
	ctx := make([]byte, 1024)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		txt = append(txt, ctx[:n]...)
	}
	return string(txt)
}
