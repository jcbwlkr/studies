package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go Ticker(&wg)
	go Work(&wg)

	wg.Wait()
}

func Work(wg *sync.WaitGroup) {
	defer wg.Done()

}

func Ticker(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("START Ticker")
	// 3600000 milliseconds is 1 hour
	ticker := time.NewTicker(time.Duration(1000) * time.Microsecond)

	for {
		select {
		case <-ticker.C:
			fmt.Printf("TICK\n")
		case <-done:
			ticker.Stop()
			return
		}
	}
}
