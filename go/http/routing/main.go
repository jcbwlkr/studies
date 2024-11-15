package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/udacity/ugo/http/web"
	"github.com/udacity/ugo/logging"
)

type v2Handler func(w http.ResponseWriter, r *http.Request) error

func v2WebAppRoute(h v2Handler) http.Handler {

	// Middleware
	// ... snip

	fn := func(w http.ResponseWriter, r *http.Request) {
		// Call handler and handle error
		err := h(w, r)
		if err != nil {
			code := http.StatusInternalServerError
			msg := http.StatusText(code)
			var webErr *web.Error
			if errors.As(err, &webErr) {
				code = webErr.Code
				msg = webErr.Error()
			}

			err := web.Encode(w, msg, code)
			if err != nil {
				slog.Error("could not encode error response", err)
			}
		}
	}
	return http.HandlerFunc(fn)
}

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type BooksHandler struct {
	db []Book
}

func (bh BooksHandler) Index(w http.ResponseWriter, r *http.Request) error {
	return web.Encode(w, bh.db, http.StatusOK)
}

func (bh BooksHandler) Get(w http.ResponseWriter, r *http.Request) error {

	// This is how we get url pattern values without httprouter
	id := r.PathValue("id")

	book, found := Book{}, false
	for _, b := range bh.db {
		if b.ID == id {
			found = true
			book = b
			break
		}
	}

	if !found {
		err := errors.New("book not found")
		return web.NewError(err, http.StatusNotFound)
	}
	return web.Encode(w, book, http.StatusOK)
}

func main() {
	db := []Book{
		{ID: "1", Title: "Goosebumps: Welcome to Dead House"},
	}

	logging.Init(logging.WithFormat("text"))

	mux := http.NewServeMux()

	books := BooksHandler{db: db}
	mux.Handle("GET /books", v2WebAppRoute(books.Index))
	mux.Handle("GET /books/{id}", v2WebAppRoute(books.Get))

	srv := http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}

	slog.Info("listening on " + srv.Addr)
	fmt.Println(srv.ListenAndServe())
}
