package tutorial

import (
	"fmt"
	"sync"
)

func defer_call() {
	defer func() { fmt.Println("print a") }()
	defer func() { fmt.Println("print b") }()
	defer func() { fmt.Println("print c") }()
	panic("error")
}

func printA(a int) {
	fmt.Println("number is ", a)
}
func defer_call_2() {
	a := 5
	defer printA(a)
	a = 10
	printA(a)
}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	wg.Done()

	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func defer_call_3() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
