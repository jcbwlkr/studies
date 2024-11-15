package main

import (
	"fmt"
	"sync"
)

type Enum struct {
	nums []int
}

func main() {
	var wg sync.WaitGroup
	wg.Add(300)
	e := Enum{}
	for i := 0; i < 300; i++ {
		go e.addNum(&wg, i)
	}
	wg.Wait()
	fmt.Println(e.nums, len(e.nums))
}

func (e *Enum) addNum(wg *sync.WaitGroup, i int) {
	e.nums = append(e.nums, i)
	wg.Done()
}
