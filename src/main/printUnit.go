package main

import (
	"fmt"
)

func printUnit() {

	const (
		_          = iota
		KB float64 = 1 << (iota * 10)
		MB
		GB
	)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
