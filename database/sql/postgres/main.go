package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	ID     string  `db:"user_id"`
	Name   string  `db:"name"`
	Roles  []int64 `db:"roles"`
	PW     string  `db:"pw"`
	PWHash []byte  `db:"pw_hash"`
}

/*
CREATE TABLE users (
	user_id UUID,
	name    VARCHAR(255),
	roles   INT[],
	pw      VARCHAR(255),
	pw_hash VARCHAR(255),

	PRIMARY KEY (user_id)
)
*/

func main() {

	log := log.New(os.Stderr, "SALES : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	j := user{
		ID:    "302058f0-4fe2-4256-a755-2e6a4bb21378",
		Name:  "Jacob",
		Roles: []int64{1, 2, 3},
		PW:    "banana",
	}

	var err error
	j.PWHash, err = bcrypt.GenerateFromPassword([]byte(j.PW), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	db, err := db()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`DELETE FROM users WHERE user_id = $1`, j.ID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`INSERT INTO users ("user_id", "name", "roles", "pw", "pw_hash")
		VALUES ($1, $2, $3, $4, $5)`,
		j.ID, j.Name, pq.Array(j.Roles), j.PW, j.PWHash,
	)
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow(`SELECT user_id, name, roles, pw, pw_hash FROM users WHERE user_id = $1`, j.ID)

	j = user{}
	err = row.Scan(&j.ID, &j.Name, pq.Array(&j.Roles), &j.PW, &j.PWHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %v %q %s\n", j.Name, j.Roles, j.PW, j.PWHash)
}

func db() (*sqlx.DB, error) {
	// Query parameters.
	q := url.Values{}
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	// Construct url.
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("postgres", ""),
		Host:     ":32917",
		Path:     "postgres",
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}
