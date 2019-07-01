package main

import (
	"runtime"
	"sync"
)

type lengthResult struct {
	n, l int
}

func maxConcurrent(n int) result {
	// Make a channel where workers will report their results. Buffer it so
	// they can each drop their result in the channel and move on.
	c := make(chan lengthResult, 1000)

	// We'll have 1 worker per cpu. Make a channel where we'll feed them work
	// and kick off a goroutine for each. Also make a waitgroup so we can know
	// when all workers are done.
	workers := runtime.NumCPU()
	work := make(chan int, 1000)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for n := range work {
				c <- lengthResult{n: n, l: wondrousLength(n)}
			}
			wg.Done()
		}()
	}

	// Start feeding the workers
	go func() {
		for i := 1; i <= n; i++ {
			work <- i
		}
	}()

	// Have this main goroutine pull `n` results to find the max
	var m, got lengthResult
	for i := 0; i < n; i++ {
		if got = <-c; got.l > m.l {
			m = got
		}
	}

	// Close the work channel so all workers will quit then wait for them to
	// wrap up.
	close(work)
	wg.Wait()

	return wondrous(m.n)
}
