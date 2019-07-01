package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/evanphx/json-patch"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var sess *mgo.Session
var seed = Contact{
	ID:        bson.NewObjectId(),
	FirstName: "Jane",
	LastName:  "Doe",
	Company:   "Jane's Homemade Salsas",
	Title:     "CEO",
	Roles:     []string{"Marketing", "Food Safety", "Executive"},
}

func init() {
	var err error
	sess, err = mgo.Dial("0.0.0.0:33039")
	if err != nil {
		log.Fatalln("dialing mongo", err)
	}
}

type App struct {
	coll *mgo.Collection
}

func NewApp(name string) http.Handler {
	coll := sess.DB(name).C("contacts")

	if _, err := coll.UpsertId(seed.ID, seed); err != nil {
		log.Fatalln("inserting seed", err)
	}

	a := App{coll: coll}

	mux := http.NewServeMux()
	mux.HandleFunc("/patch", a.patch)
	mux.HandleFunc("/manual", a.manual)
	return mux
}

func (a *App) patch(w http.ResponseWriter, r *http.Request) {
	patch, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("reading body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Fetch the full version for creating the merge
	var c Contact
	if err := a.coll.FindId(seed.ID).One(&c); err != nil {
		log.Println("fetching value", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Turn original into a []byte
	orig, err := json.Marshal(c)
	if err != nil {
		log.Println("marshalling orig", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	upd, err := jsonpatch.MergePatch(orig, patch)
	if err != nil {
		log.Println("merging patch", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var c2 Contact
	if err := json.Unmarshal(upd, &c2); err != nil {
		log.Println("unmarshalling patch", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Copy generated things like ID and metadata
	// TODO this should probably be unmarshaled on top of NewContact or like
	// UpdateContact and then we set generated things
	c2.ID = seed.ID

	// TODO validate changes?

	// Send the whole thing back to the db
	if err := a.coll.UpdateId(seed.ID, c2); err != nil {
		log.Println("updating seed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the contact back to the client
	if err := json.NewEncoder(w).Encode(c2); err != nil {
		log.Println("writing response", err)
		return
	}
}

func (a *App) manual(w http.ResponseWriter, r *http.Request) {
	var upd UpdateContact
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		log.Println("unmarshalling", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO validate UpdateContact

	// Send just the changes to the db
	if err := a.coll.UpdateId(seed.ID, upd.Changes()); err != nil {
		log.Println("updating seed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Fetch the full version for sending back to the client
	var c Contact
	if err := a.coll.FindId(seed.ID).One(&c); err != nil {
		log.Println("fetching value", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(c); err != nil {
		log.Println("writing response", err)
		return
	}
}

func main() {
}
