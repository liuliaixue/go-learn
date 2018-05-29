package tutorial

import (
	"fmt"
)

func defer_call() {
	defer func() { fmt.Println("print a") }()
	defer func() { fmt.Println("print b") }()
	defer func() { fmt.Println("print c") }()
	panic("error")
}
