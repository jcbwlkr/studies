package main

import (
	"fmt"
	"net/http"
	"strings"
)

func h(w http.ResponseWriter, r *http.Request) {
	required := []string{"email", "firstname", "lastname", "dob"}

	for _, req := range required {
		v := strings.TrimSpace(r.FormValue(req))
		if v == "" {
			http.Error(w, "Missing required field "+req, http.StatusBadRequest)
			return
		}
	}
}

func main() {
	fmt.Println("vim-go")
}
