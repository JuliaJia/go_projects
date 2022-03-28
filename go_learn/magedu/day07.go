package magedu

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Sender interface {
	Send(to string, msg string) error
}
type SenderAll interface {
	Sender
	SendAll(to string, msg string) error
}

type EmailSender struct {
	Addr     string
	Port     int
	User     string
	Password string
}

type QQSender struct {
	Addr     string
	Port     int
	User     string
	Password string
}

func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容：%s\n", to, msg)
	return nil
}
func (s EmailSender) SendAll(to string, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容：%s\n", to, msg)
	return nil
}

func (s *EmailSender) GetAll(to string, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容：%s\n", to, msg)
	return nil
}

func (s QQSender) Send(to string, msg string) error {
	fmt.Printf("发送QQ消息给: %s, 内容：%s\n", to, msg)
	return nil
}

func EmailSenderFunc() {
	var es EmailSender
	es.Send("天才哥哥", "快来帮我做作业！我是小猪头！！！！")
	var qs QQSender
	qs.Send("天才哥哥", "我是你可爱的小猪头滴滴丫！")
	var sender Sender
	sender = es
	sender.Send("小猪头", "自己的作业自己做！！！！")
	var senderall SenderAll
	senderall = es
	senderall.SendAll("天才哥哥,段子哥哥", "我是你可爱的小猪头滴滴丫！")
}

func Reflect1() {
	var es = []interface{}{1, "test", false, time.Now()}
	for _, e := range es {
		typ := reflect.TypeOf(e)
		fmt.Println(typ.Name())
	}
}

func printType(typ reflect.Type) {
	switch typ.Kind() {
	case reflect.Int, reflect.String, reflect.Float64, reflect.Bool:
		fmt.Println(typ.Name())
	case reflect.Array:
		fmt.Println(typ.Len(), typ.Elem().Name())
	default:
		fmt.Println("unknown")
	}
}

func getType(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Int, reflect.String, reflect.Float64, reflect.Bool:
		return typ.Name()
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", typ.Len(), getType(typ.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", getType(typ.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", getType(typ.Key()), getType(typ.Elem()))
	case reflect.Struct:
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("type %s struct {\n", typ.Name()))
		builder.WriteString(fmt.Sprintf("\tField(%d)\n", typ.NumField()))
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			name := field.Name
			tag := ""
			if field.Tag != "" {
				tag = fmt.Sprintf("%s", field.Tag)
			}
			ft := field.Type
			if field.Anonymous {
				name = ""
			}
			builder.WriteString(fmt.Sprintf("\t\t%s %s %s\n", name, ft, tag))
		}
		builder.WriteString(fmt.Sprintf("\tMethods(%d)\n", typ.NumMethod()))
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			name := method.Name
			mt := method.Type
			builder.WriteString(fmt.Sprintf("\t\tfunc %s (%s)\n", name, mt))
		}
		builder.WriteString(fmt.Sprintf("}\n"))
		return builder.String()
	case reflect.Func:
		var builder strings.Builder
		builder.WriteString("(")
		for i := 0; i < typ.NumIn(); i++ {
			if i != 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("%#v", typ.In(i)))
		}
		if typ.IsVariadic() {
			if typ.NumIn() > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString("...")
		}
		builder.WriteString(fmt.Sprintf(")"))
		if typ.NumOut() > 0 {
			builder.WriteString(" (")
			for i := 0; i < typ.NumOut(); i++ {
				if i != 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(fmt.Sprintf("%#v", typ.Out(i)))
			}
		}
		builder.WriteString(fmt.Sprintf(")\n"))
		return builder.String()
	case reflect.Ptr:
		var builder strings.Builder
		builder.WriteString("* {\n")
		builder.WriteString(getType(typ.Elem()))
		builder.WriteString(fmt.Sprintf("Methods(%d)\n", typ.NumMethod()))
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			builder.WriteString(fmt.Sprintf("\t func %s%s\n", method.Name, "()"))
		}
		builder.WriteString("}\n")
		return builder.String()
	default:
		return "unknown"
	}
}

func Reflect2() {
	var es = []interface{}{1, "test", false, map[int]string{1: "林克"}, [2][]int{[]int{4}, []int{5}}, []int{1, 2, 3, 4, 1, 3, 2, 4, 3, 2, 1}, time.Now(), &EmailSender{}}
	for _, e := range es {
		//val := reflect.ValueOf(e)
		typ := reflect.TypeOf(e)
		fmt.Println(getType(typ))
	}
}
