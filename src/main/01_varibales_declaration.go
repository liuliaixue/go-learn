package main

import "fmt"

func variablesDeclaration() {

	var score int
	fmt.Println("var score int; score = ", score)
	score = 100
	fmt.Println("score = 100; score = ", score)

	var name = "Alan"
	fmt.Println(name)

	var array1 [2]int
	var array2 = [2]int{200, 201}
	array3 := [2]int{300, 301}
	array4 := [2]int{1: 401, 0: 400}
	fmt.Println(array1, array2, array3, array4)

	array5 := [...]int{500, 501, 502, 503, 504, 505, 506, 507, 508, 509}
	fmt.Println(array5)

	slice1 := array5[3:7]
	slice2 := array5[0:]
	slice3 := array5[:len(array5)]
	fmt.Println("\n", slice1, "\n", slice2, "\n", slice3)

}
