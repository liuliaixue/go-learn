package main

import "fmt"

func printFunc_1(a int, b int) (int, string) {

	sum := a + b
	res := "OK"
	return sum, res
}

func printFunc_2() int {
	a, b, c, d, e := 1, 2, 3, 4, 5
	sum := printFunc_2_child(a, b, c, d, e)
	fmt.Println(a, b, c, d, e)
	return sum
}

func printFunc_2_child(s ...int) int {
	var sum int
	for i, _ := range s {
		s[i] = 1
		sum += s[i]
	}
	return sum
}

func printFunc_3_child(num *int) {
	*num = *num * 2

}
func printFunc_3() {
	a := 10
	fmt.Println(a)
	printFunc_3_child(&a)
	fmt.Println(a)
}

func printFunc_4_child(num int) func(int) int {
	return func(y int) int {
		return num + y
	}
}

func printFunc_4() {
	fc := printFunc_4_child(12)
	sum := fc(88)
	fmt.Println(sum)
}

func printFunc_defer() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func printFunc_defer_panic_recover_1() {

	a := func() {
		fmt.Println("func a")
	}
	b := func() {
		fmt.Println("func b")
		panic("panic b")
	}
	c := func() {
		fmt.Println("func c")
	}

	a()
	b()
	c()
}

func printFunc_defer_panic_recover_2() {

	a := func() {
		fmt.Println("func a")
	}
	b := func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover b")
			}
		}()
		fmt.Println("func b")
		panic("panic b")
	}
	c := func() {
		fmt.Println("func c")
	}

	a()
	b()
	c()
}

func printFunc_defer_2() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer_f i = ", i)
		}()
		fmt.Println(&i, i)
		fs[i] = func() {
			fmt.Println("fs i = ", i)
		}
	}
	for _, f := range fs {
		f()
	}
}
