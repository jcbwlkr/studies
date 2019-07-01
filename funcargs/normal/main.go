package main

import (
	"fmt"
	"net/http"
	"time"
)

// ClientConfig is a function that knows how to modify a Client.
type ClientConfig func(*Client)

// Timeout gives a config func that can sets the client's request timeout to a
// new duration.
func Timeout(d time.Duration) ClientConfig {
	return func(c *Client) {
		c.c.Timeout = d
	}
}

// Client makes BLAH BLAH requests.
type Client struct {
	c *http.Client
}

// NewClient constructs a Client with default settings. Some configuration
// options can be changed by providing 1 or more ClientConfig functions.
func NewClient(opts ...ClientConfig) *Client {
	c := Client{
		c: &http.Client{
			Timeout: 15 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

func main() {

	fn := func(*Client) {
		fmt.Println("Shouldn't work")
	}

	c := NewClient(Timeout(1*time.Minute), fn)

	fmt.Println(c)
}
