package basic

import (
	"fmt"
	"strconv"
)

func typeConverter() {
	a, b, c, d, e := "3", 3, int64(3), 3.0, "3.14"

	fmt.Printf("%T:%v %T:%v %T:%v %T:%v ,%T:%v\n", a, a, b, b, c, c, d, d, e, e)

	aToInt, err := strconv.Atoi(a)
	if err != nil {
		print(err)
	} else {
		fmt.Printf("%T:%v to int is %v\n", a, a, aToInt)
	}

	bToString := strconv.Itoa(b)
	fmt.Printf("%T:%v to string is %v\n", b, b, bToString)

	cToString := strconv.FormatInt(c, 10)
	fmt.Printf("%T:%v to string is %v\n", c, c, cToString)

	dToString32 := strconv.FormatFloat(d, 'E', -1, 32)
	dToString64 := strconv.FormatFloat(d, 'E', -1, 64)
	fmt.Printf("%T:%v to string is %v\n", d, d, dToString32)
	fmt.Printf("%T:%v to string is %v\n", d, d, dToString64)
}
