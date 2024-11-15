package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Product is something we sell.
type Product struct {
	Name       string      `json:"name"`
	Attributes *Attributes `json:"attributes"`
}

// Attributes defines the characteristics of a Product.
type Attributes struct {
	Size  string `json:"size"`
	Color string `json:"color"`
}

// Scan implements database/sql.Scanner which allows us to Unmarshal Attributes
// JSON data when it comes out of the database.
func (a *Attributes) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New("Attribute scan data was not []byte")
	}

	var val *Attributes
	if err := json.Unmarshal(b, &val); err != nil {
		return err
	}

	if val != nil {
		*a = *val
	}

	return nil
}

func (a *Attributes) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	b, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	//runUser()
	//return

	db, err := sql.Open("mysql", "root:@tcp(0.0.0.0:9999)/test")
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range []string{"shirt", "pants", "undefined"} {
		p, err := fetchProduct(db, name)
		if err != nil {
			log.Printf("%10s : Error(%v)", name, err)
		} else {
			log.Printf("%10s : %+v", name, p.Attributes)
		}
	}

	p1 := Product{Name: "No Attributes"}

	if err := saveProduct(db, p1); err != nil {
		log.Println("Could not save p1", err)
	}

	p2 := Product{
		Name: "Has Attributes",
		Attributes: &Attributes{
			Size:  "enormous",
			Color: "teal",
		},
	}

	if err := saveProduct(db, p2); err != nil {
		log.Println("Could not save p2", err)
	}
}

func fetchProduct(db *sql.DB, name string) (Product, error) {
	var p Product

	row := db.QueryRow("SELECT name, attributes FROM products WHERE name = ?", name)
	if err := row.Scan(&p.Name, &p.Attributes); err != nil {
		return Product{}, err
	}

	return p, nil
}

func saveProduct(db *sql.DB, p Product) error {
	_, err := db.Exec(
		"INSERT INTO products (name, attributes) VALUES (?, ?)",
		p.Name,
		p.Attributes,
	)
	return err
}
