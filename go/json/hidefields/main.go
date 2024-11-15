package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User is a user
type User struct {
	Name     string `json:"name"`               // Always marshalled
	Password string `json:"password,omitempty"` // Never marshalled
	Memo     string `json:"memo,omitempty"`     // Sometimes marshalled

	// Set to true before marshalling to get extra information
	Details bool `json:"-"`
}

// MarshalJSON ensures that certain fields are zeroed out before marshalling to
// JSON. It implements the json.Marshaler interface.
func (u User) MarshalJSON() ([]byte, error) {
	type PublicUser User
	p := PublicUser(u)
	p.Password = ""

	if !p.Details {
		p.Memo = ""
	}

	return json.Marshal(p)
}

func main() {

	in := []byte(`{"name": "Jacob", "password": "banana", "memo": "foobar"}`)

	fmt.Printf("Initial fixture data\n=> %s\n", string(in))

	var u User
	err := json.Unmarshal(in, &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unmarshalled value:\n=> %+v\n", u)

	out, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled json:\n=> %s\n", string(out))

	u.Details = true
	out, err = json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled json with details:\n=> %s\n", string(out))
}
