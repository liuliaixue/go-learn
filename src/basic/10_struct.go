package basic

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name string
	Age  int
}

func struct_person() {
	p := person{}
	fmt.Println(p)
	p.Name = "yoyo"
	p.Age = 300
	fmt.Println(p)

	p2 := person{
		Name: "jojo",
		Age:  400,
	}
	fmt.Println(p2)

	struct_person_update(p2)
	fmt.Println(p2)

	struct_person_update_pointer(&p2)
	fmt.Println(p2)

}

func struct_person_update(per person) {
	per.Age = 500
	fmt.Println("struct_person_update", per)
}

func struct_person_update_pointer(per *person) {
	per.Age = 600
	fmt.Println("struct_person_update_pointer", per)
}

func struct_define() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "koko",
		Age:  21,
	}
	fmt.Println(a)

}

type person_2 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type contact_3 struct {
	Phone, City string
}

type person_3 struct {
	Name    string
	Age     int
	Contact contact_3
}

func struct_define_2() {
	a := person_2{Name: "wewe", Age: 22}
	fmt.Println(a)

	b := person_3{Name: "p3", Age: 33}
	b.Contact = contact_3{Phone: "1818", City: "Shanghai"}
	fmt.Println(b)

}

func struct_compare() {

	a := person_2{Name: "wewe", Age: 22}
	b := person_2{Name: "wewe", Age: 22}

	fmt.Println(a == b)

}

func struct_composite() {

	type human struct {
		Sex int
	}

	type teacher struct {
		human
		Name string
		Age  int
	}
	type student struct {
		human
		Name string
		Age  int
	}

	a := teacher{human: human{Sex: 1}, Name: "tt", Age: 40}
	b := student{Name: "ss", Age: 20}

	fmt.Println(a, b)
}

func struct_composite_2() {

	type human struct {
		Sex int
	}

	type teacher struct {
		human
		Name string
		Age  int
		Sex  int
	}

	a := teacher{human: human{Sex: 1}, Name: "tt", Age: 40}

	fmt.Println(a)
}

func structToJSON() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println("")
	fmt.Println(b)

	fmt.Println(string(b))
}

// ColorGroup ...
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type colors struct {
	color []string
}

func structToJSON2() {

	var JSONString = []byte(`{
		"ID": 1,
		"Name": "alan",
		"Colors": ["#333333"]
	}`)
	var jsonObj ColorGroup
	// fmt.Println(JSONString)
	fmt.Println(string(JSONString))
	err := json.Unmarshal(JSONString, &jsonObj)
	if err != nil {
		fmt.Println("error 2 structToJSON2 \n", err)
		return
	}
	fmt.Println(jsonObj)
	fmt.Println("ID:", jsonObj.ID)
	fmt.Println("Name:", jsonObj.Name)
	fmt.Println("Colors:", jsonObj.Colors)

}

func structToJSON3() {

	var JSONString1 = []byte(`{
		"ID": 1,
		"Name": "alan",
		"Colors":[ "#333333"]
	}`)
	var jsonObj1 ColorGroup
	// fmt.Println(JSONString1)
	fmt.Println(string(JSONString1))
	err1 := json.Unmarshal(JSONString1, &jsonObj1)
	if err1 != nil {
		fmt.Println("error 3", err1)
	}
	fmt.Println(jsonObj1)
	fmt.Println("ID:", jsonObj1.ID)
	fmt.Println("Name:", jsonObj1.Name)
	fmt.Println("Colors:", jsonObj1.Colors)
}
