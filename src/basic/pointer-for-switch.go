package basic

import "fmt"

func pointer() {
	var a = 10
	var p = &a
	fmt.Println("pointer address is ", p)
}

func printif() {
	a := 10
	if a := 199; a > 0 {
		fmt.Println("a is bigger than", 0)
	}
	fmt.Println(a)
}

func printfor_1() {
	a := 10
	for {
		a++
		if a > 12 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("over")
}

func printfor_2() {
	a := 10
	for a < 12 {
		a++
		fmt.Println(a)
	}
	fmt.Println("over")
}

func printfor_3() {
	a := 10
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Println("over")
}

func printswitch() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a==0")
	case 1:
		fmt.Println("a==1")
	default:
		fmt.Println("default")
	}

}

func printswitch_fallthrough() {
	a := 10
	switch {
	case a > 0:
		fmt.Println("a>0")
		fallthrough
	case a > 2:
		fmt.Println("a>2")
	case a > 4:
		fmt.Println("a>4")
	default:
		fmt.Println("default")
	}

}

func print_continue_goto() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL

		}
	}

	//code below is infinite
	// LABEL1:
	// for i := 0; i < 10; i++ {
	// 	for {
	// 		fmt.Println(i)
	// 		continue LABEL1

	// 	}
	// }
}
