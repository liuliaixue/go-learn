package main

import "fmt"

func max(n1 int, n2 int) int {

	if n1 > n2 {
		return n1
	}

	return n2
}

func revert(n1 int, n2 int) (int, int) {
	return n2, n1
}

type cb func(int) int

func plusOne(i int) int {
	return i + 1
}

func testCallback(i int, callback cb) {
	callback(i)
}

type person struct {
	name string
	age  int
}

func (p person) getName() string {
	return p.name
}

func main() {

	n1, n2 := 1, 200
	fmt.Println(max(n1, n2))
	fmt.Println(revert(n1, n2))

	// function as parameter
	testCallback(10, plusOne)

	// function methond
	var p = person{"Alan", 13}
	fmt.Println(p)
	fmt.Println(p.getName())

}
