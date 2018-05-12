package main

import (
	"fmt"
)

func printType() {

	const a int = 1
	const A = 'A'

	const (
		b = 2
		c = 3
		d
	)
	fmt.Println("printType")
	fmt.Println(a, A, b, c, d)

	fmt.Println(^12)

}
