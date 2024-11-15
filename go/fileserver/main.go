package main

import "io"
import "net/http"
import "net/http/httputil"

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/media/", http.StripPrefix("/media/", fs))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		io.WriteString(w, "That wasn't found. You said:\n")
		b, _ := httputil.DumpRequest(r, false)
		w.Write(b)
		io.WriteString(w, "\n")
	}))
	http.ListenAndServe(":5177", nil)
}
