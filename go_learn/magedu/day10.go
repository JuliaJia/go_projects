package magedu

import (
	"bufio"
	"fmt"
	htmlTemplate "html/template"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"regexp"
	"text/template"
)

func Regexp1(str string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入需要检查的手机号：")
	scanner.Scan()
	phone := scanner.Text()
	fmt.Println(regexp.MatchString(str, phone))
}

func Regexp2(str string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入需要检查的字符串：")
	scanner.Scan()
	phone := scanner.Text()
	fmt.Println(regexp.MatchString(str, phone))
}

func Regexp3(str1, str2, str3 string) {
	reg, err := regexp.Compile(str1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reg.MatchString(str2))
	fmt.Println(reg.ReplaceAllString(str3, str2))
	fmt.Println(reg.FindAllString(str3, -1))
	fmt.Println(reg.Split(str3, -1))
}

type Calculator struct {
}

type CalculatorRequest struct {
	Left  int
	Right int
}

type CalculatorResponse struct {
	Result int
}

func (c *Calculator) Add(r *CalculatorRequest, rep *CalculatorResponse) error {
	log.Println("log!")
	rep.Result = r.Left + r.Right
	return nil
}

func RpcServer(addr, port, pro string) {
	addr = addr + ":" + port
	rpc.Register(&Calculator{})
	listen, err := net.Listen(pro, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func RpcClient(addr, port, pro string) {
	addr = addr + ":" + port
	conn, err := jsonrpc.Dial(pro, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	r := &CalculatorRequest{2, 5}
	rep := &CalculatorResponse{}
	err = conn.Call("Calculator.Add", r, rep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rep.Result)
}

func Template1(str string) {
	tpltext := "你好，我是{{ . }}"
	tpl, err := template.New("tpl").Parse(tpltext)
	fmt.Println(err)
	tpl.Execute(os.Stdout, str)
}

func Template2(str string) {
	tpltext := "你好，我是{{ . }}"
	htmlTpl, err := htmlTemplate.New("tpl").Parse(tpltext)
	fmt.Println(err)
	htmlTpl.Execute(os.Stdout, str)
}

func Template3() {
	//tpltext := `
	//	我的ID是: {{ .ID }}
	//	我的名字是: {{ .Name }}
	//	我的性别是: {{ if eq .Sex 1 }}男{{ else }}女{{ end }}
	//`
	tpltext := `
		{{ range .}}
			我的ID是: {{ .ID }}
			我的名字是: {{ .Name }}
			我的性别是: {{ if eq .Sex 1 }}男{{ else }}女{{ end }}
		{{ end }}
	`
	tpl := template.Must(template.New("tpl").Parse(tpltext))
	tpl.Execute(os.Stdout, []struct {
		ID   int
		Name string
		Sex  int
	}{{1, "林克", 1}, {2, "塞尔达", 0}})
}
