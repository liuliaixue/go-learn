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

func printSlice() {
	a1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a1)
	s1 := []int{}
	s2 := a1[:len(a1)]
	s3 := a1[5:]
	fmt.Println(s1, s2, s3)
	// fmt.Println(printArray)
}

func print_slice_from_array() {
	a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'k'}

	s1 := a[3:6]
	s2 := s1[2:]
	s3 := s1[2:7]

	fmt.Println(string(s1))
	fmt.Println(string(s2))
	fmt.Println(string(s3))
	fmt.Println(string(s2[0]))

}

func print_slice_append() {

	s1 := make([]int, 3, 6)
	fmt.Printf("%p %v \n", s1, s1)

	s1 = append(s1, 3, 4, 5)
	fmt.Printf("%p %v \n", s1, s1)

	s1 = append(s1, 3, 4, 5)
	fmt.Printf("%p %v \n", s1, s1)

}

func print_slice_copy() {

	s1 := []int{0, 1, 2, 3, 4, 5}
	s2 := []int{7, 8, 9}
	fmt.Println(s1, s2)

	copy(s1, s2)
	fmt.Println(s1, s2)

	s3 := []int{0, 1, 2, 3, 4, 5}
	s4 := []int{7, 8, 9}
	copy(s3, s4)
	fmt.Println(s3, s4)
}
