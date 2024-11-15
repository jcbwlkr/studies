package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("= Start ========================================================================")
		b, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(b))
		fmt.Println("= End ==========================================================================")
	})

	fmt.Println("Server started on :9191")
	http.ListenAndServe(":9191", nil)
}
