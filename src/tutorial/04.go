package tutorial

// ref : https://studygolang.com/articles/11869

import (
	"fmt"
	"unsafe"
)

func sizeOfType() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
	fmt.Println("")

	/**
	从上面的输出，我们可以推断出 a 和 b 为 int 类型，且大小都是 32 位（4 字节）。
	如果你在 64 位系统上运行上面的代码，会有不同的输出。
	在 64 位系统下，a 和 b 会占用 64 位（8 字节）的大小。
	*/
}

func sizeOfType_2() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

func type_3() {
	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //不允许 int + float64
	fmt.Println(sum)
}
