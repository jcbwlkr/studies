package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

type myTransport struct{}

func (myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("got request for", req.URL)

	if req.URL.Path == "/fail" {
		fmt.Println("failing request early")
		return &http.Response{
			StatusCode: http.StatusBadRequest,                  // 400
			Status:     http.StatusText(http.StatusBadRequest), // Bad Request
			Body:       ioutil.NopCloser(&bytes.Reader{}),      // Have to provide an io.ReadCloser
		}, nil
	}

	fmt.Println("sending upstream")
	return http.DefaultTransport.RoundTrip(req)
}

func main() {
	rp := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Point at the new location
			req.URL.Scheme = "http"
			req.URL.Host = "httpbin.org"

			// httpbin.org needs these because it's behind its own reverse proxy (heroku)
			req.Header.Set("Host", "httpbin.org")
			req.Host = "httpbin.org"
		},
		Transport: myTransport{},
	}

	fmt.Println("started")
	http.ListenAndServe("localhost:9292", &rp)
}
