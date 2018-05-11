package main

import (
	"fmt"
)

const a int = 1
const A = 'A'

const (
	b = 2
	c = 3
	d
)

func printType() {
	fmt.Println("printType")
	fmt.Println(a, A, b, c, d)

	fmt.Println(^12)

}
