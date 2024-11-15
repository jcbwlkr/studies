package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User is a user
type User struct {
	Name     string `json:"name"` // Always marshalled
	Password string `json:"-"`    // Never marshalled but can be unmarshalled
}

func (u *User) UnmarshalJSON(data []byte) error {

	// Unmarshal over a modified version of User that exposes Password
	type testUser User
	var tmp struct {
		testUser
		Password string `json:"password"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// Assign our user values and password values to the value at *u
	*u = User(tmp.testUser)
	u.Password = tmp.Password

	return nil
}

// Operator is a user with a phone number
type Operator struct {
	User

	Phone string `json:"phone"`
}

func (o *Operator) UnmarshalJSON(data []byte) error {
	// First unmarshal over Operator itself to get operator specific fields
	type testOperator Operator
	var to testOperator
	if err := json.Unmarshal(data, &to); err != nil {
		return err
	}
	*o = Operator(to)

	// Next unmarshal over the User fields using custom User unmarshalling
	//var u User
	//if err := json.Unmarshal(data, u); err != nil {
	//}
	//u
	return nil
}

func main() {
	in := []byte(`{"name": "Jacob", "password": "banana", "phone": "316.555.1234"}`)

	fmt.Printf("Initial fixture data\n=> %s\n", string(in))

	var u Operator
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
}
