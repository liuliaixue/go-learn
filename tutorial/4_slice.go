package main

import (
	"fmt"

	use "github.com/liuliaixue/go-use/src"
)

func main() {
	fmt.Println("slice")

	var a1 [10]int
	a2 := []int{1}
	a3 := [10]int{1}
	a4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)

	b := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	// fmt.Println(b)

	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c1 := c[0:10]
	c2 := c[5:]
	c3 := c[:8]
	c4 := c[:]
	fmt.Println(c)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	use.Use(c)

	use.Use(a1, a2, a3, a4, b)

}
