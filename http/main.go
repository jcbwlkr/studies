package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
		log.Println(r.RequestURI)

		//log.Printf("%s%s", r.Host, r.URL)
		//log.Println(r.URL.String())
	})

	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This is on page 2.\n")
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
