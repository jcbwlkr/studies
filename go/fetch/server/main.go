package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Log the client's incoming request
		req, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Request:")
		fmt.Println(string(req))

		// All responses include these CORS headers.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		// POST requests get to the critical section.
		if r.Method == http.MethodPost {
			w.Write([]byte("JACKPOT"))
		}

	})

	http.ListenAndServe(":8080", nil)
}
