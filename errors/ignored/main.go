package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func verify(age int) error {
	if age < 18 {
		return errors.New("too young")
	}

	return nil
}

func main() {

	age := rand.Intn(100)
	if err := verify(age); err != nil {
		log.Fatal("Some problem! No idea what though.")
	}

	fmt.Println("Welcome!")
}
