package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID string

	UserInfo
}

type UserInfo struct {
	FirstName string
	LastName  string
}

func main() {

	u := User{
		ID: "abcdefgh",
		UserInfo: UserInfo{
			FirstName: "Rory",
			LastName:  "Walker",
		},
	}

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
