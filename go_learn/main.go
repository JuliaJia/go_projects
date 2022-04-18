package main

import (
	"fmt"
	"go_learn/learn"
)

//type Addable interface {
//type int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr,float32,float64,complex64,complex128,string
//}
//
//func add[T Addable](a,b T) T {
//	return a+b
//}

//func Test01(){
//	add(1,2)
//}

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
	//magedu.Dir2("magedu", func(path string) bool {
	//	return strings.HasSuffix(path, ".go")
	//}, func(path string) {
	//	fmt.Println("file: ", path)
	//	fmt.Println("contentx: ")
	//	fmt.Println(magedu.ReadFile(path))
	//})
	//fp := magedu.FilePath("password.txt")
	//fmt.Println(magedu.FileName(fp))
	//fmt.Println(magedu.FileExt(fp))
	//fp2 := "/123/456/...../////..../////./////123"
	//fmt.Println(magedu.PathClean(fp2))
	//magedu.Buffer1("我是林克！", "我不是塞尔达！")

	//magedu.CopyFile3("password.txt", "password.copy", []byte(string(1024*1024*1024)))
	//magedu.CopyFile4("password.txt", "password.copy", 1024)
	//magedu.MultiWriter1("log1.txt", "log2.txt", "我是林克！不是塞尔达！")
	//buffer := make([]byte, 1024)
	//magedu.MultiReader1("log1.txt", "log2.txt", buffer)
	//magedu.MultiReader2("log1.txt", "log2.txt")
	//magedu.Ioutil1("io1.txt", []byte("我是天才哥哥！"))
	//input := magedu.Ioutil2(".")
	//magedu.Log1("test.log", input)
	//magedu.FileList(".")
	//magedu.OsFile()
	//magedu.BufioScanner("log1.txt")
	//magedu.BufioScanner2()
	//magedu.GobEncode("gob_encode.txt")
	//magedu.GobDecode("gob_encode.txt")
	//ctx := magedu.JsonEncode()
	//magedu.JsonDecode(ctx, "jsonfile.json")
	//magedu.JsonCheck("jsonfile2.json")
	//magedu.JsonEncodeFile("jsonfile2.json")
	//magedu.EmailSenderFunc()
	//magedu.Reflect2()

	//time.Sleep(time.Second * 3)
	//magedu.GoRoutine1()
	//magedu.GoRoutine2(20)
	//magedu.GoRoutine3()
	//magedu.Goroutine4(5)
	//magedu.Chan1(30)
	//magedu.Chan2()
	//fmt.Println(magedu.FileLine("magedu/day05.go"))
	//magedu.FileLineCount(".")
	//magedu.FileLineTotalCountChannel(".")
	//magedu.FileLineTotalCount(".")
	//roc, _ := magedu.ChannelType("RO", 10)
	//_, woc := magedu.ChannelType("WO", 10)
	//channel := make(chan int, 10)
	//channel, _, _ = magedu.IntChannelTypeUse(channel, "WO")
	//magedu.IntChannelTypeUse(channel, "RO")
	//magedu.RunTime1()
	//magedu.RunTime2()
	//worker := magedu.NewPool(5)
	//worker.AddTask(func() interface{} {
	//	return 1
	//})
	//worker.AddTask(func() interface{} {
	//	return 2
	//})
	//worker.AddTask(func() interface{} {
	//	return 3
	//})
	//worker.Start()
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	for result := range worker.Results {
	//		fmt.Println(result)
	//	}
	//	wg.Done()
	//}()
	//worker.Wait()
	//wg.Wait()
	//magedu.Net4()
	//var (
	//	host     string
	//	port     string
	//	typeInfo string
	//	pro      string
	//)
	//flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	//flag.StringVar(&port, "P", "22", "连接端口")
	//flag.StringVar(&typeInfo, "t", "Server", "服务类型")
	//flag.StringVar(&pro, "p", "Tcp", "协议类型")
	//flag.Usage = func() {
	//	fmt.Println("usage: -t [Server | Client] -p [tcp | udp] [-H 127.0.0.1] [-P 22]")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()
	//if typeInfo == "Server" {
	//	magedu.Socket1(pro, host, port)
	//} else if typeInfo == "Client" {
	//	magedu.Socket2(pro, host, port)
	//}

	//magedu.HttpWeb11()

	//magedu.Regexp3("132\\d{8}", "132????????", "我的电话是13212312312，13212312313，15812312313")

	//var (
	//	host     string
	//	port     string
	//	typeInfo string
	//	pro      string
	//)
	//flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	//flag.StringVar(&port, "P", "22", "连接端口")
	//flag.StringVar(&typeInfo, "t", "Server", "服务类型")
	//flag.StringVar(&pro, "p", "Tcp", "协议类型")
	//flag.Usage = func() {
	//	fmt.Println("usage: -t [Server | Client] -p [tcp | udp] [-H 127.0.0.1] [-P 22]")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()
	//if typeInfo == "Server" {
	//	magedu.RpcServer(host, port, pro)
	//} else if typeInfo == "Client" {
	//	magedu.RpcClient(host, port, pro)
	//}

	//magedu.Template3()
	//dsn := "root:123456@tcp(192.168.33.10:3306)/go_learn?charset=utf8mb4&loc=PRC&parseTime=true"
	//magedu.Sql1("mysql", dsn)
	//var sum = func(a, b int) int {
	//	return a + b
	//}
	//nv := learn.Op(sum, 3)
	//fmt.Println(nv)
	//f := learn.Sub()
	//f()
	//f()
	//fmt.Println()
	//learn.Inf01("123", "456")
	//Test01()
	//learn.Reflect12()
	//learn.Basic3()
	//learn.Ticker()
	//learn.Timer()
	//learn.Constant()
	//rect := learn.Nan("12.435464")
	//if math.IsNaN(rect) {
	//	fmt.Println("解析出错！")
	//} else {
	//	fmt.Println(rect)
	//}

	//learn.Ceil(1.4)
	//learn.Floor(1.4)
	//learn.Trunc(-1.4)
	//learn.Modf(-1.4)
	//learn.Abs(-1.4)
	//learn.Max(3, 4)
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(learn.RanShuffleInt2(arr))
	//var str string
	//learn.Std(str)
	//learn.ReadFile2("go.mod")
	//var str string
	//learn.WriteFile2("learn.txt", str)
	//learn.FileOpInfo2("learn")
	//learn.TempLogCreate("test", "learn/")
	//learn.Compress("learn/day02.go", "learn/day02.go.zlib")
	//learn.SysCall("ls", "-a", "-l", "learn")
	//learn.CompressReader("learn/day02.go.zlib")
	//learn.JsonFunc()
	//str := learn.Base64Func("1.png", 1024)
	//learn.Base64Decode("2.png", str)
	learn.Mutex02()
}
