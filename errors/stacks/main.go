package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

// printBottomStack prints an error message with wrapped context but only
// prints the stack trace of the bottom-most error that has a stack.
func printBottomStack(err error) {

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	type causer interface {
		Cause() error
	}

	fmt.Println("error:", err)

	// Keep track of the last stackTracer seen as we walk down the
	// error stack.
	var finalStacker stackTracer

	// Walk down the Cause chain, updating the finalStacker.
	for err != nil {

		if stacker, ok := err.(stackTracer); ok {
			finalStacker = stacker
		}

		// Replace err with its cause if present.
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	if finalStacker == nil {
		// No stack traces found
		return
	}

	fmt.Println("Stack Trace:")
	for _, f := range finalStacker.StackTrace() {
		fmt.Printf("%+v\n", f)
	}
}

func main() {
	err := findUser(123)
	//printBottomStack(err)
	fmt.Printf("error: %v\nStack Trace:\n%+v\n", err, err)
}

////////////////////////////////////////////////////////////////////////////////
// Example Code that generates a stack of wrapped errors ///////////////////////
////////////////////////////////////////////////////////////////////////////////

func findUser(id int) error {
	return errors.Wrapf(query(id), "finding user %d", id)
}

func query(id int) error {
	return errors.Wrap(open(), "running query")
}

func open() error {
	return stderrors.New("could not open file")
}
