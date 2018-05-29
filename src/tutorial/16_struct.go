package tutorial

import (
	"fmt"
)

func employee() {
	type Employee struct {
		firstName string
		lastName  string
		age       int
	}

	// var employeeA struct {
	// 	firstName string
	// 	lastName  string
	// 	age       int
	// }
	emp1 := Employee{"alan", "liu", 18}
	fmt.Println(emp1)

	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	emp4 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	emp4.age = 998
	fmt.Println("Employee 4", emp4)
}
