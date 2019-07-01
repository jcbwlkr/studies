package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	hits = make(map[string]int)
	mu   sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	hits[r.URL.Path]++
	fmt.Fprintln(w, hits)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/foo", handler)
	http.HandleFunc("/bar", handler)

	http.ListenAndServe(":8181", nil)
}
