package main

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name string `validate:"required"`
}

func main() {

	var u User

	err := validator.New().Struct(u)

	fmt.Println(err)
}
