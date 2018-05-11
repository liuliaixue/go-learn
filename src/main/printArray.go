package main

import (
	"fmt"
)

func printArray() {
	fmt.Println("printArray")

	var a = [2]int{}
	fmt.Println(a)

	b := [20]int{1: 1, 2: 2, 19: 1}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	fmt.Println("printArray array comparation")

	d1 := [2]int{1, 2}
	d2 := [2]int{1, 2}
	fmt.Println(&d1, &d2)
	fmt.Println(d1 == d2)

	d3 := new([10]int)
	d4 := new([10]int)
	fmt.Println(&d3, &d4)
	fmt.Println(d3 == d4)

	e := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(e)

	f := [...]int{1, 2, 9, 8, 3, 4, 6, 5, 2, 2, 2}
	fmt.Println(f)
	num := len(f)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if f[i] < f[j] {
				temp := f[j]
				f[j] = f[i]
				f[i] = temp
			}
		}
	}
	fmt.Println(f)
}
