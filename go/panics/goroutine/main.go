package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	kill := make(chan struct{})

	go func() {
		fmt.Println("in another goroutine, waiting for the kill command...")
		<-kill
		fmt.Println("got kill signal")
		time.Sleep(100 * time.Millisecond)
		panic("panic from goroutine!")
	}()

	fmt.Println("starting server")
	srv := http.Server{
		Addr: ":9090",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("in handler. Sending kill signal")
			close(kill)

			panic("panic in handler. Does this kill my server?")
		}),
	}
	srv.ListenAndServe()
}
