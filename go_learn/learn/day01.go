package learn

import (
	"fmt"
	"reflect"
	"sync"
)

type Transporter interface {
	move(string, string) (int, error)
	say(int)
}

type Car struct {
}

type Ship struct {
}

func (Car) move(string, string) (int, error) {
	return 1, nil
}

func (Car) say(a int) {

}

func (Ship) move(string, string) (int, error) {
	return 2, nil
}

func (Ship) say(a int) {

}

type add func(a, b int) int

func Op(f add, a int) int {
	b := 2 * a
	return f(a, b)
}

func Sub() func() {
	i := 10
	fmt.Printf("%p\n", &i)
	b := func() {
		fmt.Printf("i addr %p\n", &i)
		i--
		fmt.Println(i)
	}
	return b
}

func Add(a, b int) int {
	return a + b
}

func inf01(src, dest string, transporter Transporter) error {
	num, err := transporter.move(src, dest)
	fmt.Println(num)
	return err
}

func Inf01(src, dest string) {
	var car Car
	var ship Ship
	inf01(src, dest, car)
	inf01(src, dest, ship)
}

type People struct {
	User
	Username string
	Name     string
	Age      int
	Gender   byte
}

type User struct {
	Username string `json:"UserName"`
	Name     string
	Age      int
	Gender   byte
	uint8
}

type UserInt interface {
	GetInfo()
}

var (
	sUser *User
	uOnce sync.Once
)

func (u *User) GetInfo() int {
	fmt.Println(u.Username)
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	fmt.Println(u.Gender)
	return 1
}

func (User) GetInfo2() {
	fmt.Println("hello")
}

func NewUser(name, username string, age int, gender byte) *User {
	return &User{
		Username: username,
		Name:     name,
		Age:      age,
		Gender:   gender,
	}
}

func NewDefaultUser() *User {
	return &User{
		Username: "",
		Name:     "",
		Age:      -1,
		Gender:   3,
	}
}

func GetUserInstance() *User {
	uOnce.Do(func() {
		if sUser == nil {
			sUser = NewDefaultUser()
		}
	})
	return sUser
}

func Reflect01() {
	typeUser := reflect.TypeOf(&User{})
	fmt.Println(typeUser)
	fmt.Println(typeUser.Elem())
	fmt.Println(typeUser.Kind())
	fmt.Println(typeUser.Elem().Kind())
}

func Reflect02() {
	typeUser := reflect.TypeOf(User{})
	typeUserNum := typeUser.NumField()
	for i := 0; i < typeUserNum; i++ {
		field := typeUser.Field(i)
		fmt.Println(field.Name)
		fmt.Println(field.Offset)
		fmt.Println(field.Anonymous)
		fmt.Println(field.Type)
		fmt.Println(field.IsExported())
		fmt.Println(field.Tag.Get("json"))
		fmt.Println()

	}
}

func Reflect03() {
	typeUser := reflect.TypeOf(User{})
	typeUserMethodNum := typeUser.NumMethod()
	for i := 0; i < typeUserMethodNum; i++ {
		field := typeUser.Method(i)
		fmt.Println(field.Name)
		fmt.Println(field.Type)
		fmt.Println(field.IsExported())
		fmt.Println(field.IsExported())
		fmt.Println(field.Type.NumIn())
		fmt.Println(field.Type.NumOut())
		fmt.Println()

	}
}

func Reflect04() {
	typeUser := reflect.TypeOf(&User{})
	typeUserMethodNum := typeUser.NumMethod()
	for i := 0; i < typeUserMethodNum; i++ {
		field := typeUser.Method(i)
		fmt.Println(field.Name)
		fmt.Println(field.Type)
		fmt.Println(field.IsExported())
		fmt.Println(field.IsExported())
		fmt.Println(field.Type.NumIn())
		fmt.Println(field.Type.NumOut())
		fmt.Println()

	}
}

func Reflect05() {
	funcType := reflect.TypeOf(NewUser)
	fmt.Println(funcType.Kind() == reflect.Func)
	argInNum := funcType.NumIn()
	argOutNum := funcType.NumOut()
	fmt.Println()
	for i := 0; i < argInNum; i++ {
		argType := funcType.In(i)
		fmt.Println(argType)
	}
	fmt.Println()
	for i := 0; i < argOutNum; i++ {
		argType := funcType.Out(i)
		fmt.Println(argType)
	}
}

func Reflect06() {
	userType := reflect.TypeOf((*UserInt)(nil)).Elem()
	peopleType := reflect.TypeOf(&User{})
	fmt.Println(peopleType.Implements(userType))
}

func Reflect07() {
	var s string = "Hello"
	vs := reflect.ValueOf(&s)
	vs.Elem().SetString("Hello World")
	fmt.Println(s)
	fmt.Println()
	user := User{}
	vu := reflect.ValueOf(&user)
	userNameValue := vu.Elem().FieldByName("Username")
	if userNameValue.CanSet() {
		userNameValue.SetString("Ryomajia")
	}

	fmt.Println(user.Username)
}

func Reflect08() {
	vf := reflect.ValueOf(Add)
	args := []reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)}
	results := vf.Call(args)
	sum := results[0].Interface().(int)
	fmt.Println(sum)
}

func Reflect09() {
	user := User{}
	vu := reflect.ValueOf(&user)
	getMethod := vu.MethodByName("GetInfo")
	rv := getMethod.Call([]reflect.Value{})
	result := rv[0].Interface().(int)
	fmt.Println(result)
}

func Reflect10() {
	ut := reflect.TypeOf(User{})
	uv := reflect.New(ut)
	uv.Elem().FieldByName("Username").SetString("RyomeJia")
	uv.Elem().FieldByName("Age").SetInt(18)
	if user, ok := uv.Interface().(*User); ok {
		fmt.Println(user)
	}

}

func Reflect11() {
	sliceType := reflect.TypeOf([]User{})
	sliceValue := reflect.MakeSlice(sliceType, 2, 10)
	sliceValue.Index(0).Set(reflect.ValueOf(User{Age: 10, Gender: 1}))
	sliceValue.Index(1).Set(reflect.ValueOf(User{Age: 20, Gender: 2}))

	s := sliceValue.Interface().([]User)
	fmt.Println(s[0].Age)
	fmt.Println(s[1].Age)
	fmt.Println()

	sliceValue2 := reflect.ValueOf(&s)
	sliceValue2.Elem().SetLen(4)
	sliceValue2.Elem().Index(0).FieldByName("Gender").SetUint(2)
	sliceValue2.Elem().Index(2).Set(reflect.ValueOf(User{Age: 30, Gender: 1}))
	sliceValue2.Elem().Index(3).Set(reflect.ValueOf(User{Age: 40, Gender: 1}))
	s2 := sliceValue2.Elem().Interface().([]User)
	fmt.Println(s2[0].Gender)
	fmt.Println(s2[1].Age)
	fmt.Println(s2[2].Age)
	fmt.Println(s2[3].Age)
}

func Reflect12() {
	var userMap map[string]*User
	mapType := reflect.TypeOf(userMap)
	mapValue := reflect.MakeMap(mapType)
	user := &User{Username: "LongMore"}
	key := reflect.ValueOf(user.Username)
	mapValue.SetMapIndex(key, reflect.ValueOf(user))
	userMap = mapValue.Interface().(map[string]*User)
	fmt.Println(userMap)
}
