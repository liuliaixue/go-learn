// https://studygolang.com/articles/11872
package tutorial

import (
	"fmt"
)

func constant() {
	var name = "Wuji"
	fmt.Printf("type %T value is %v \n", name, name)
}

func const_int() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}

func const_type() {
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T \n", i, f, c)

}
