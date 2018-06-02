package tutorial

import (
	"fmt"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {

	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'u' || rune == 'i' || rune == 'o' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func lookString() {
	name := MyString("My name is Alan")
	var v VowelsFinder
	v = name
	fmt.Printf("%c \n", v.FindVowels())
}

/*******************************************************************************/

type SalaryCalculator interface {
	CalculateSalary() int
}

type Pernanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (emp Pernanent) CalculateSalary() int {
	return emp.basicpay + emp.pf
}

func (emp Contract) CalculateSalary() int {
	return emp.basicpay
}

func salary() {
	emp1 := Pernanent{1, 5000, 1000}
	emp2 := Pernanent{2, 5000, 2000}
	emp3 := Contract{3, 5000}

	employees := []SalaryCalculator{emp1, emp2, emp3}
	expense := 0
	for _, v := range employees {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
	fmt.Println()

}
