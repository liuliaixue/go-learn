package main

import "fmt"

func main() {

	n1, n2 := 1, 200
	fmt.Println(max(n1, n2))
	fmt.Println(revert(n1, n2))

}

func max(n1 int, n2 int) int {

	if n1 > n2 {
		return n1
	}

	return n2
}

func revert(n1 int, n2 int) (int, int) {
	return n2, n1
}
