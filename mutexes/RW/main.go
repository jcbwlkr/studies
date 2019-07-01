package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	m = make(map[int]string)
	l sync.RWMutex
)

func main() {

	// Fill the map
	for i := 0; i < 100; i++ {
		m[i] = strconv.Itoa(i)
	}

	var wg sync.WaitGroup

	// Fire off 3 goroutines to access the map waiting a bit between each read
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < 100; j++ {
				l.RLock()
				fmt.Printf("goroutine %d: %d=%s\n", i, j, m[j])
				l.RUnlock()
				time.Sleep(200 * time.Millisecond)
			}
			wg.Done()
		}(i)
	}

	// Have this goroutine (the main one) access the map and change it after a bit
	time.Sleep(2 * time.Second)
	l.Lock()
	fmt.Println("Changing the map! Stop the world for a bit")
	time.Sleep(2 * time.Second)
	for i := 0; i < 100; i++ {
		m[i] = strconv.Itoa(i * 2)
	}
	l.Unlock()

	wg.Wait()
	fmt.Println("done")
}
