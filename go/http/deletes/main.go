package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Body seen at server:")
		io.Copy(os.Stdout, r.Body)
	}))
	defer ts.Close()

	body := strings.NewReader(`This is the body of my delete`)
	req, err := http.NewRequest("DELETE", ts.URL, body)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := http.DefaultClient.Do(req); err != nil {
		log.Fatal(err)
	}
}
