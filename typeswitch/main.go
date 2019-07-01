package main

import (
	"fmt"
	"io"
)

type file struct{}

func (*file) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (*file) Write([]byte) (int, error) {
	return 0, nil
}

func main() {

	var r io.Reader
	r = new(file)

	switch r.(type) {
	case io.Writer:
		fmt.Println("Concrete type also implements io.Writer")

	case *file:
		fmt.Println("Concrete type is a *file")
	}
}
