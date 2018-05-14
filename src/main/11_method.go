package main

import "fmt"

type methon_person struct {
	Name string
}

func (a methon_person) printName() {
	fmt.Println("name is " ,a.Name)
}

func (a *methon_person) updateName(name string) {
	a.Name = name
}

func method_struct() {
	a := methon_person{}
	a.printName()
	a.updateName("yoyo")
	fmt.Println(a)
}
/*******	*******************************/
type number int

func method_increase() {
	var a number
	a.Increase();
	fmt.Println(a)
}

func (num *number) Increase() {
	*num = number(100)
}