package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	sess, err := mgo.Dial("localhost:32770")
	if err != nil {
		log.Fatal(err)
	}

	if err := sess.DB("test").DropDatabase(); err != nil {
		log.Fatal(err)
	}

	c := consumer{
		ID:        bson.NewObjectId(),
		TenantID:  "b0514d98-57df-4c0e-8904-f72d8a1c7a38",
		PublicID:  "9302d959-92ef-423b-9140-7393a0e13a28",
		PrivateID: "e46eac38-9a7a-4ce9-b944-b5285fd32b75",
		Name:      "Acme Co",
	}

	// Put the consumer in
	if err := sess.DB("test").C("consumers").Insert(c); err != nil {
		log.Fatal(err)
	}

	// Pull a new consumer out
	var c2 consumer
	if err := sess.DB("test").C("consumers").FindId(c.ID).One(&c2); err != nil {
		log.Fatal(err)
	}

	if c.Name != c2.Name {
		log.Fatalf("Names didn't match! %q and %q", c.Name, c2.Name)
	}

	fmt.Println("Welcome,", c2.Name)

	// Find by name
	var c3 consumer
	q := bson.M{
		"name_hash": hash(c.PrivateID, c.Name),
	}
	if err := sess.DB("test").C("consumers").Find(q).One(&c3); err != nil {
		log.Fatal(err)
	}

	if c.Name != c3.Name {
		log.Fatal("Names didn't match!")
	}

	fmt.Println("Welcome,", c3.Name)

	qry := bson.M{"_id": c.ID}
	upd := bson.M{"$set": bson.M{"name": "Jacob Co"}}
	if err := sess.DB("test").C("consumers").Update(qry, upd); err != nil {
		log.Fatal(err)
	}
}
