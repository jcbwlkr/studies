// Comes from Michal Nicpon

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	foo := "bar"

	m := map[string]*string{
		"a": &foo,
		"b": &foo,
		"c": &foo,
	}

	start := make(chan struct{})
	go func() {
		<-start
		for {
			for k := range m {
				*m[k] = fmt.Sprintf("%d", rand.Int())
			}
		}
	}()

	close(start)

	end := time.After(100 * time.Millisecond)

	for {
		select {
		case <-end:
			return
		default:
		}

		for k, v := range m {
			fmt.Println(k, *v)
		}
	}
}
