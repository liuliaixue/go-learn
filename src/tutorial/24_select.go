package tutorial

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	// time.Sleep(6 * time.Second)
	ch <- "From server1 "
}

func server2(ch chan string) {

	// time.Sleep(2 * time.Second)
	ch <- "From server2 "
}

func server_select() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println("s1", s1)
	case s2 := <-output2:
		fmt.Println("s2", s2)
	}
}
