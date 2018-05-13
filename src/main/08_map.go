package main

import "fmt"

func printMap() {

	m := make(map[int]string)
	m[1] = "ok"
	fmt.Println(m)
	delete(m, 1)

	fmt.Println(m)
}

func printMap_2() {

	m := make(map[int]map[int]string)

	m[1] = make(map[int]string)
	fmt.Println(m)

	m[1][1] = "OK"
	fmt.Println(m)

	delete(m[1], 1)
	fmt.Println(m)

	delete(m, 1)
	fmt.Println(m)
}

func print_range() {
	sm := make([]map[int]string, 5)
	for i, _ := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "ok"
	}
	fmt.Println(sm)
}

func print_range_2() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := make(map[string]int, 3)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)

}
