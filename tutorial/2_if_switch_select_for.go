package main

import (
	"fmt"
	"time"
)

// Phone Phone
type Phone interface {
	call()
}

// APhone APhone
type APhone struct {
}

// IPhone IPhone
type IPhone struct {
}

func (p APhone) call() {
	fmt.Println("I am APhone")
}
func (p IPhone) call() {
	fmt.Println("I am IPhone")
}

func main() {
	// ifelse
	b := true
	if b {
		fmt.Println("b is true")
	}

	v := 20
	//switch
	switch {
	case v > 1:
		fmt.Println("v > 1")
		fallthrough
	case v > 10:
		fmt.Println("v > 10")
	default:
		fmt.Println("default")
	}

	switch v {
	case 1:
		fmt.Println("v is 1")
	case 20:
		fmt.Println("v is 20")
	default:
		fmt.Println("default")
	}

	// switch type
	var p Phone
	p = IPhone{}
	switch p.(type) {
	case IPhone:
		fmt.Printf("%v type is IPhone\n", p)
	case APhone:
		fmt.Printf("%v type is APhone\n", p)
	default:
		fmt.Printf("%v %T\n", p, p)
		fmt.Println("var p Phone default")
	}

	// select
	fmt.Println()
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	//for

	for {
		fmt.Println("无限循环")
	}

}
