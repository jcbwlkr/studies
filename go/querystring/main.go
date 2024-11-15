package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// myTime embeds time.Time and adds additional encoding methods
type myTime struct {
	time.Time
}

// EncodeValues implements the query.Encoder interface by encoding myTime
// values in YYYY-MM-DD format.
func (t myTime) EncodeValues(name string, vals *url.Values) error {
	vals.Set(name, t.Format("2006-01-02"))
	return nil
}

// User is a type that uses myTime
type User struct {
	Name    string
	Created myTime
}

func main() {

	u := User{
		Name:    "Jacob",
		Created: myTime{time.Now()},
	}

	val, err := query.Values(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val.Encode()) // Prints Created=2017-10-19&Name=Jacob

	// When passing values of myTime to things expecting a time.Time you have to
	// pass the embedded value
	someFunc(u.Created.Time)
}

func someFunc(t time.Time) {
}
