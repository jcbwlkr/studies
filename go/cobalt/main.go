package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/ardanlabs/cobalt"
)

func hello(c *cobalt.Context) {
	c.Serve(map[string]string{
		"hello": "world",
	})
	fmt.Println("sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}

func main() {
	app := cobalt.New(new(myJSON))

	app.Get("/", hello)

	t := 5 * time.Second
	app.Run(":8080", t, t)
}

type myJSON struct{}

func (j *myJSON) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (j *myJSON) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (j *myJSON) ContentType() string {
	return "application/json"
}
