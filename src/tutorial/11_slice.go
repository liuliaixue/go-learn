package tutorial

import "fmt"

func slice_type() {
	var a [3]int
	fmt.Printf("%v  \n", a)
	a[0] = 12
	fmt.Printf("%v  \n", a)

	var b []int = []int{0, 2, 3}
	fmt.Printf("%v \n", b)
	fmt.Printf("%+v \n", b)
	fmt.Printf("%T \n", b)
	fmt.Printf("%t \n", true)

	fmt.Printf("%d \n", 10)
	fmt.Printf("%b \n", 10)
	fmt.Printf("%x \n", 10)

	fmt.Printf("%c \n", 65)
}

func slice_append() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
}
