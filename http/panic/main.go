package main

import (
	"fmt"
	"log"
	"net/http"
)

func mw(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "error:", err)
			}
		}()

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/die" {
		panic("OH NO")
	}
	fmt.Fprintln(w, "Hello, world")
}

func main() {

	http.Handle("/", mw(http.HandlerFunc(hello)))

	log.Fatal(http.ListenAndServe(":9000", nil))
}
