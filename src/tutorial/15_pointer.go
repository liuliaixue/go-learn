package tutorial

import (
	"fmt"
)

func ponter_c() {
	a := 444222
	b := &a
	// fmt.Println("a: ", a)
	// fmt.Println("b: ", b)

	fmt.Printf("a is %v , address: %v \n", a, &a)
	fmt.Printf("b is %v , address: %v \n", b, &b)
}

func _change(a *int) {
	*a = *a + 1
}
func pointer_change() {
	a := 10
	fmt.Println("a: ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
	_change(&a)
	fmt.Println("a plus 1 ", a)
}
