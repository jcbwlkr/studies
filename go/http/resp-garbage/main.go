package main

import (
	"io"
	"net/http"
)

func NewApp() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/write", write)
	mux.HandleFunc("/writestring", writestring)
	mux.HandleFunc("/bytestostring", bytestostring)
	mux.HandleFunc("/stringtobytes", stringtobytes)
	return mux
}

func write(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world"))
}

func writestring(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func bytestostring(w http.ResponseWriter, r *http.Request) {
	body := []byte("hello, world")
	io.WriteString(w, string(body))
}

func stringtobytes(w http.ResponseWriter, r *http.Request) {
	body := "hello, world"
	w.Write([]byte(body))
}

func main() {
}
