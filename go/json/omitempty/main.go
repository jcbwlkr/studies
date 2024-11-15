package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Thing struct {
	Name         string     `json:"name,omitempty" bson:"name"`
	Age          int        `json:"age,omitempty" bson:"age,omitempty"`
	Foo          string     `json:"-"`
	ConcreteTime time.Time  `json:"ctime,omitempty"`
	PointerTime  *time.Time `json:"ptime,omitempty"`
	User         *struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user,omitempty"`
	password string
}

func main() {

	var t Thing
	t.Name = "Jacob"
	//t.Age = 32
	t.Foo = "Bar"

	fmt.Printf("%+v\n", t)

	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	return

	bb, err := bson.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	var t2 Thing
	if err := bson.Unmarshal(bb, &t); err != nil {
		log.Fatal(err)
	}

	fmt.Println("T2", t2)
}
