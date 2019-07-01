package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func main() {
	t, err := thing()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
}

func thing() (string, error) {
	doc, err := find()
	if err != nil {
		return "", errors.New("OH NO")
	}

	var user struct {
		Name string
	}
	if err := json.Unmarshal(doc, &user); err != nil {
		return "", err
	}

	return user.Name, nil
}

func find() ([]byte, error) {

	body := `{"Name": "Rory"}`
	return []byte(body), nil

}
