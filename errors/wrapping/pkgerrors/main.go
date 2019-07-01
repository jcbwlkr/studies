package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// ErrNotFound is used when a User is not found.
var ErrNotFound = &AppError{404}

func main() {

	err := findUser(123)

	// Print the wrapped error string.
	fmt.Println("Normal error message")
	fmt.Printf("error: %v\n\n", err)

	// Print with details to see the call site of each wrapping.
	fmt.Println("Error with details")
	fmt.Printf("error: %+v\n\n", err)

	// Check if the original cause is a certain kind of error

	if _, ok := errors.Cause(err).(*AppError); ok {
		fmt.Println("That's an app error")
	}

	// Get access to the original error and get a copy.

	if apperror, ok := errors.Cause(err).(*AppError); ok {
		fmt.Println("That's an app error and its state is:", apperror.State)
	}

	if errors.Cause(err) == ErrNotFound {
		fmt.Println("That's a Not Found error")
	}
}

func findUser(id int) error {
	if err := query(id); err != nil {
		return errors.Wrapf(err, "finding user %d", id)
	}
	return nil
}

func query(id int) error {
	if err := open(); err != nil {
		return errors.Wrapf(err, "building query")
	}
	return nil
}

func open() error {
	return ErrNotFound
}

// AppError is a custom application error type.
type AppError struct {
	State int
}

func (a *AppError) Error() string {
	return "app error code " + strconv.Itoa(a.State)
}
