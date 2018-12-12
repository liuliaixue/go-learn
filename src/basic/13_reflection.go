package basic

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
	Sex  int
}

type Manager struct {
	User
	title string
}

func (user User) Hello(name string) {
	fmt.Println("hello, ", name, "I am ", user.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println(t)
	fmt.Println("type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
		// fmt.Println(f, t)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func reflection_1() {
	user := User{1, "good", 17, 0}
	Info(user)
}

func reflection_2() {

	m := Manager{User: User{1, "good", 17, 0}, title: "master"}
	fmt.Println(m)
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}).Name)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}).Name)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}).Name)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 3}).Name)
	fmt.Printf("%#v\n", t.Field(1))
}

func reflection_3() {
	x := 123
	v := reflect.ValueOf(x)
	fmt.Println(v)
	vv := reflect.ValueOf(&x)
	vv.Elem().SetInt(998)
	fmt.Println(x)

}

func reflection_4() {
	u := User{1, "Alice", 17, 0}
	u.Hello("C.C.")
	SetUser(&u)
	fmt.Println(u)

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("V.V.")}
	fmt.Println(args)
	mv.Call(args)
}

func SetUser(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("pointer cannot set")
		return
	} else {
		v = v.Elem()
	}

	if f := v.FieldByName("Id"); f.Kind() == reflect.Int {
		f.SetInt(999)
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("Blob")
	}
}
