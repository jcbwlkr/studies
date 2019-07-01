package main

import (
	"fmt"
	"sync"
	"time"
)

type thing struct {
	c  chan int
	wg *sync.WaitGroup
}

func (t *thing) work() {
	defer t.wg.Done()
	fmt.Println("Start goroutine")
	for i := range t.c {
		fmt.Println("received", i)
	}
	fmt.Println("End goroutine")

}

func main() {

	t := &thing{
		c:  make(chan int),
		wg: &sync.WaitGroup{},
	}

	t.wg.Add(1)
	go t.work()

	t.c <- 2
	t.c = nil

	time.Sleep(1 * time.Second)

	fmt.Println("put it back?")
	t.c = make(chan int)
	t.c <- 3
	close(t.c)

	fmt.Println("wait for it to finish")
	t.wg.Wait()

	fmt.Println("Done")
}
