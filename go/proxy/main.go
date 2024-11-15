package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	p := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			dump, _ := httputil.DumpRequest(r, false)
			r.URL.Host = r.Host
			r.URL.Scheme = "http"
			log.Println(string(dump))
		},
	}

	http.ListenAndServe(":80", p)
}
