package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-r.Context().Done():
			log.Println("Request was canceled! Abort!")
			return
		case now := <-t.C:
			log.Println("tick", now)
		}
	}
}

func main() {

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}
