package magedu

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Net1() {
	fmt.Println(net.JoinHostPort("127.0.0.1", "6666"))
	fmt.Println(net.JoinHostPort("::1", "6666"))
	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	fmt.Println(net.SplitHostPort("[::1]:9999"))
	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupHost("localhost"))
	fmt.Println(net.LookupIP("www.baidu.com"))
	fmt.Println(net.ParseCIDR("192.168.1.1/24"))
}

func Net2() {
	fmt.Printf("%#v\n", net.ParseIP("192.168.1.1"))
	ip, ipnet, err := net.ParseCIDR("192.168.1.1/24")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(ipnet.Contains(ip))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.25")))
	fmt.Println(ipnet.Network())
}

func Net3() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.Network(), addr.String())
	}
}

func Net4() {
	intfs, _ := net.Interfaces()
	for _, intf := range intfs {
		fmt.Println(intf.Name, intf.HardwareAddr, intf.MTU)
	}
}

func Socket1(pro, ipaddr, port string) {
	addr := ipaddr + ":" + port
	listener, err := net.Listen(pro, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	} else {
		defer conn.Close()
		log.Printf("Client is connected...%s\n", conn.RemoteAddr().String())
		reader := bufio.NewReader(conn)
		scanner := bufio.NewScanner(os.Stdin)
		for {
			line, _, err := reader.ReadLine()

			if err == nil {
				if string(line) == "quit" {
					break
				}
				log.Printf("接受到的数据是：%s\n", string(line))
				fmt.Println("请输入回复：")
				scanner.Scan()
				fmt.Fprintf(conn, "%s\n", scanner.Text())
			} else if err == io.EOF {
				break
			} else {
				log.Println(err)
			}
		}

	}
}

func Socket2(pro, ipaddr, port string) {
	addr := ipaddr + ":" + port
	flag := true
	conn, err := net.Dial(pro, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("Connected!")
	for {

		if flag {
			fmt.Println("请输入消息：")
			scanner.Scan()

			fmt.Fprintf(conn, "%s\n", scanner.Text())
			line, _, err := reader.ReadLine()
			if string(line) == "quit" {
				flag = false
			}
			if err == nil {
				log.Printf("接受到的数据是：%s\n", string(line))
			} else if err == io.EOF {
				break
			} else {
				log.Println(err)
			}
		} else {
			break
		}
	}

}

func HttpWeb1() {
	timeFunc := func(response http.ResponseWriter, request *http.Request) {
		//response.Write([]byte("hello,world!"))
		fmt.Println(request)
	}
	http.HandleFunc("/time/", timeFunc)
	http.ListenAndServe(":7777", nil)
}

type timeHandleFunc struct {
}

func (h *timeHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(w, now)
}

func HttpWeb2() {
	http.Handle("/time/", &timeHandleFunc{})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL, r.Proto)
		fmt.Printf("%T,%#v\n", r.Header, r.Header)
		fmt.Println(r.Header.Get("User-Agent"))
	})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.Form.Get("a"))
		fmt.Println(r.Form["a"])
	})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb5() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.FormValue("a"))
	})
	http.ListenAndServe(":7777", nil)
}
func HttpWeb6() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.Form.Get("a"))
		fmt.Println(r.Form["a"])
		fmt.Println(r.PostForm)
	})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb7() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 1024)
		fmt.Println(r.MultipartForm)
		file, _ := r.MultipartForm.File["a"][0].Open()
		io.Copy(os.Stdout, file)
	})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb8() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, header, _ := r.FormFile("a")
		io.Copy(os.Stdout, file)
		fmt.Println(header.Filename)
		fmt.Println(header.Size)
		fmt.Println(header.Header)
	})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb9() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var info map[string]interface{}
		decoder.Decode(&info)
		fmt.Println(info)

	})
	http.ListenAndServe(":7777", nil)
}

func HttpWeb10() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("Cookie"))
		counter := 1
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter),
			HttpOnly: true,
		}
		http.SetCookie(w, counterCookie)
	})
	http.ListenAndServe(":7777", nil)
}
func HttpWeb11() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":7777", nil)
}
