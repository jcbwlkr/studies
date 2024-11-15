package main

import (
	"fmt"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	Street string
}

type User struct {
	Name  string
	Email string
	Roles []string
	Home  *Address `bson:"home,omitempty"`
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:32768"},
		Timeout:  10 * time.Second,
		Database: "admin",
		Username: "",
		Password: "",
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}

	session.DB("test").DropDatabase()

	u := User{
		Name:  "Jacob",
		Email: "jacob@example.com",
		Roles: []string{
			"foo", "bar", "baz",
		},
		Home: &Address{Street: "1234 Woodlawn"},
	}

	coll := session.DB("test").C("foo")
	err = coll.Insert(u)
	if err != nil {
		log.Fatal(err)
	}

	u = User{}
	err = coll.Find(bson.M{"name": "Jacob"}).One(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted in db: %+v\n", u)

	err = coll.Update(nil, bson.M{"$set": bson.M{"home": nil}})
	if err != nil {
		log.Fatal(err)
	}

	u = User{}
	err = coll.Find(bson.M{"name": "Jacob"}).One(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After update: %+v\n", u)
}
