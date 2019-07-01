package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const sessName = `myApp`

func main() {
	store := sessions.NewCookieStore([]byte("secret"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s %s", r.Method, r.URL.String())

		s, err := store.Get(r, sessName)
		if err != nil {
			log.Print(err)
		}

		c, ok := s.Values["count"]
		if !ok {
			log.Print("New user detected")
			c = 0
		}

		s.Values["count"] = c.(int) + 1

		if err := s.Save(r, w); err != nil {
			log.Print("Could not update session")
		}

		fmt.Fprintf(w, "You have visited this site %d times.", s.Values["count"])
	})

	http.ListenAndServe(":8383", nil)
}
