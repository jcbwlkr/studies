package main

import (
	"errors"
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("vim-go")
}

func work() {
	if err := task(); err != nil {
		return handleError(err)
	}

	return
}

func handleError(err error) {
	if err != nil {
		slog.Error("failure", "error", err)
	}
}

func task() error {
	return errors.New("oh no")
}
