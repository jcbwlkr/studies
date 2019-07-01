package main

import (
	"errors"
	"fmt"
	"strconv"
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

	// Check if any wrapped error is a certain kind of error.

	var apperror1 *AppError
	if errors.As(err, &apperror1) {
		fmt.Println("That's an app error")
	}

	// Get access to the original error and get a copy.

	var apperror2 *AppError
	if errors.As(err, &apperror2) {
		fmt.Println("That's an app error and its state is:", apperror2.State)
	}

	// Check if any wrapped error is equal to a Sentinel error.

	if errors.Is(err, ErrNotFound) {
		fmt.Println("That's a Not Found error")
	}
}

func findUser(id int) error {
	if err := query(id); err != nil {
		return fmt.Errorf("finding user %d: %w", id, err)
	}
	return nil
}

func query(id int) error {
	if err := open(); err != nil {
		return fmt.Errorf("building query: %w", err)
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
