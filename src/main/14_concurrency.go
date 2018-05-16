package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrency_1() {
	go func() {
		fmt.Println("GO GO Go")
	}()
	time.Sleep(2 * time.Second)
}

func concurrency_2() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO Go")
		c <- true
	}()
	<-c
}

func concurrency_3() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO Go")
		c <- true
		c <- true
		c <- true
		c <- false
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func concurrency_4() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go PrintGo(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}

}

func concurrency_5() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go PrintGo_5(&wg, i)
	}
	for i := 0; i < 10; i++ {
	}

	wg.Wait()

}

func PrintGo(c chan bool, index int) {
	var a int
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}

func PrintGo_5(wg *sync.WaitGroup, index int) {
	var a int
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}

// func concurrency_6() {
// 	c1, c2 := make(chan int), make(chan string)
// 	o := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case v, ok := <-c1:
// 				if !ok {
// 					o <- true
// 					break
// 				}
// 				fmt.Println("c1", v)
// 			case v, ok := <-c2:
// 				if !ok {
// 					o <- true
// 					break
// 				}
// 				fmt.Println("c2", v)
// 			}
// 		}
// 	}()

// 	c1 <- 1
// 	c2 <- "ok"
// 	c1 <- 12
// 	c2 <- "false"

// 	// close(c1)
// 	// close(c2)

// 	// <-o
// }

// func concurrency_7() {
// 	c := make(chan int64)

// 	concurrency_7_producer(c, 10)
// 	concurrency_7_consumer(c)

// }
// func concurrency_7_producer(c chan int64, max int) {
// 	defer close(c)

// 	for i := 0; i < max; i++ {
// 		c <- time.Now().Unix()
// 	}
// }

// func concurrency_7_consumer(c chan int64) {
// 	var v int64
// 	ok := true
// 	for ok {
// 		if v, ok = <-c; ok {
// 			fmt.Println(v)
// 		}
// 	}
// }
