package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func twoSum(a []int, b int) [2]int {

	c := [2]int{}

outline:
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == b {
				c = [2]int{i, j}
				break outline
			}
		}
	}

	fmt.Println(c)
	return c
}
