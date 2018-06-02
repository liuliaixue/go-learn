package tutorial

import (
	"fmt"
	"sync"
)

var x = 0
var y = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func increment_no_conflict(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

func mutex_conflict() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

func mutex_no_conflict() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment_no_conflict(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", y)
}
