package main

import (
	"fmt"
	"sync"
)

// This is a race condition. Both write to the map concurrently.
//
// My cursor is on this line
//
// This word is missspelled
// Search for this phrase
// TODO something important...
func main() {
	m := map[int]int{}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go fillMap(wg, m, 0, 999)
	go fillMap(wg, m, 1000, 1999)
	wg.Wait()

	fmt.Println(len(m))
}

func fillMap(wg *sync.WaitGroup, m map[int]int, start, end int) {
	for i := start; i <= end; i++ {
		m[i] = i
	}
}

func Something() {
}
