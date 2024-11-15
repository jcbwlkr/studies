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

// Operator is a user with a phone number
type Operator struct {
	User

	Phone string `json:"phone"`
}

// MarshalJSON ensures that the extra fields on Operator are included when it
// is marshaled. Without this method the inner type's MarshalJSON method would
// be called and extra fields are ignored.
func (o Operator) MarshalJSON() ([]byte, error) {
	u, err := json.Marshal(o.User)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(u, &m); err != nil {
		return nil, err
	}

	m["phone_number"] = o.Phone

	return json.Marshal(m)
}

func main() {
	op := Operator{
		User: User{
			Name:     "Operator Ophelia",
			Password: "obiwan",
			Memo:     "Hi!",
			Details:  true,
		},
		Phone: "316.555.1234",
	}

	out, err := json.Marshal(op)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled json:\n=> %s\n", string(out))
}
