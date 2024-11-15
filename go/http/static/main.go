package main

import "net/http"

func main() {

	http.HandleFunc("/", hello)

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("resources"))))

	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
