package tutorial

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	fmt.Println("going to print fullName")
	defer recoverName()
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func panic_panic() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
