package tutorial

import (
	"fmt"
)

func printBytes() {

	s := "This is a string"
	for i := 0; i < len(s); i++ {
		// fmt.Println(s[i])
		fmt.Printf("%x ", s[i])
	}
}

func printChars() {

	s := "This is a string"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printChars_range() {
	s := "This is a string"
	for _, value := range s {
		fmt.Printf("%c ", value)
	}
	fmt.Println()

}
