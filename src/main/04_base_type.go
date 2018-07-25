package main

import (
	"fmt"
)

func baseTypeBoolean() {
	// boolean
	var a bool
	a = true
	b := false
	c := a && b
	fmt.Printf("true && false is %v\n", c)
}

func baseTypeInt() {
	{
		a, b, c, d, e := 12, int8(-127), int16(12), int32(12), int64(12)
		fmt.Printf("%T:%v %T:%v %T:%v %T:%v %T:%v \n", a, a, b, b, c, c, d, d, e, e)
	}
	{

		a, b, c, d, e := 12, uint8(0), uint16(12), uint32(12), uint64(12)
		fmt.Printf("%T:%v %T:%v %T:%v %T:%v %T:%v \n", a, a, b, b, c, c, d, d, e, e)
	}
}
func baseTypeExtra() {
	// byte is uint8
	// rune is uint32
}

func baseTypeFloat() {}

func baseTypeComplex() {}

func baseTypeString() {}

func baseTypeConverter() {
	a := 12
	b := 18.3
	c := a + int(b)
	fmt.Print(c)
}
