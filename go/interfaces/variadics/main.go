package main

import (
	"bytes"
	"io"
)

func main() {

	// So I have a bunch of these things which each implement io.Reader
	var bufs []bytes.Buffer

	// _ = io.MultiReader(bufs...)
	// This gives the error:
	//	cannot use bufs (type []bytes.Buffer) as type []io.Reader in argument to io.MultiReader

	// Go refuses to implicitly convert a slice of values to a slice of interface
	// values so you have to do it yourself.
	var rs []io.Reader
	for i := range bufs {
		rs = append(rs, &bufs[i])
	}

	// Also note because MultiReader is a variadic function you'd have to use the
	// ... operator to pass them all in
	_ = io.MultiReader(rs...)
}
