package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type app struct {
	db *sql.DB
}

func (a *app) getAllTodos(w http.ResponseWriter, r *http.Request) {
	// use a.db in code
}

func (a *app) createTodo(w http.ResponseWriter, r *http.Request) {
	// use a.db in code
}

func (a *app) getSpecificTodo(w http.ResponseWriter, r *http.Request) {
	// use a.db in code
}

func (a *app) deleteTodo(w http.ResponseWriter, r *http.Request) {
	// use a.db in code
}

func NewApp(db *sql.DB) http.Handler {

	a := app{db: db}

	r := mux.NewRouter()

	r.HandleFunc("/todo", a.getAllTodos).Methods("GET")
	r.HandleFunc("/todo", a.createTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", a.getSpecificTodo).Methods("GET")
	r.HandleFunc("/todo/{id}", a.deleteTodo).Methods("DELETE")

	return r
}

func main() {
	db, err := sql.Open("mysql", "some-dsn")
	if err != nil {
		log.Fatal(err)
	}

	a := NewApp(db)

	log.Fatal(http.ListenAndServe(":3000", a))
}
