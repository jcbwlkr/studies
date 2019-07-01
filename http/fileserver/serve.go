package main

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
)

var t = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
  <head>
    <title>{{.Path}}</title>
  </head>
  <body>
    {{.Body}}
  </body>
</html>
`))

type page struct {
	Path string
	Body template.HTML
}

// wrap intercepts the response from http.FileServer and makes it a complete
// HTML document.
func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		resp := httptest.NewRecorder()
		h.ServeHTTP(resp, r)

		// Copy status and headers to the actual response. Skip content-length
		// because it does not account for our extra html.
		w.WriteHeader(resp.Code)
		for k := range resp.Header() {
			if k != "Content-Length" {
				w.Header().Set(k, resp.Header().Get(k))
			}
		}

		if resp.Body.Len() > 0 {
			data := page{
				Path: r.URL.Path,
				Body: template.HTML(resp.Body.String()),
			}
			if err := t.Execute(w, data); err != nil {
				log.Println("couldn't execute template", err)
			}
		}
	})
}

func main() {
	http.ListenAndServe(":8000", wrap(http.FileServer(http.Dir("."))))
}
