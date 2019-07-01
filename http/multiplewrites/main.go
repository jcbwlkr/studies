package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "var x = 123;\n")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "\nhello, world\n")
}

func main() {
	srv := httptest.NewServer(http.HandlerFunc(handler))
	defer srv.Close()

	res, err := http.Get(srv.URL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
