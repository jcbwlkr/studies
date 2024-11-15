package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/justinas/alice"
)

func main() {
	c := alice.New(middleware)

	http.Handle("/", c.ThenFunc(handler))

	http.ListenAndServe(":8181", nil)
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := context.WithValue(r.Context(), "user", "Jacob")

		h.ServeHTTP(w, r.WithContext(c))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value("user").(string)
	if !ok {
		u = "World"
	}

	fmt.Fprintf(w, "Hello, %s\n", u)
}
