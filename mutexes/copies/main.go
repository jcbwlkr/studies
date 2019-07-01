package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type myThing struct {
	state int64
}

func (t *myThing) Lock() {
	// Try to change 0 to 1 every 10 ms. Stop looping once it works.
	for !atomic.CompareAndSwapInt64(&t.state, 0, 1) {
		time.Sleep(10 * time.Millisecond)
	}
}

func (t *myThing) Unlock() {
	// Try to change 1 to 0. If it doesn't work then we messed up.
	swapped := atomic.CompareAndSwapInt64(&t.state, 1, 0)
	if !swapped {
		panic("OMG!")
	}
}

func main() {
	var t myThing
	work(t)
}

func work(t myThing) {
	t.Lock()
	defer t.Unlock()

	fmt.Println("Hi")
}
