package main

import (
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", ":8181")

	if err != nil {
		// I couldn't listen. Must be in use!
		os.Exit(1)
	}

	// I was able to listen so I guess it's available!
	l.Close()
}
