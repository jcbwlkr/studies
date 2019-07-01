package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	if err := write(time.Now().Format(time.Kitchen)); err != nil {
		log.Fatal(err)
	}
}

func write(msg string) error {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer closeAndLog(f)

	if _, err := fmt.Fprintln(f, msg); err != nil {
		return err
	}

	return nil
}

func closeAndLog(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Print(err)
	}
}
