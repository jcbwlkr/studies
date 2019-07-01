package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := fetch(); err != nil {
		log.Fatal(err)
	}
}

func fetch() error {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/golang/go", nil)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// x is a Go value that will hold the decoded response from the API. In this
	// case I don't know/care what I am going to get back. Normally I would use a
	// more meaningful type than just interface{}
	var x interface{}
	if err := json.NewDecoder(resp.Body).Decode(&x); err != nil {
		return err
	}

	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
