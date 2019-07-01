package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var t = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
  <head>
    <title>{{.Title}}</title>
  </head>
  <body>
    {{.Body}}
  </body>
</html>
`))

type htmlResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

// Write intercepts all writes to the response body so we can put it in our
// template when we call Complete.
func (w *htmlResponseWriter) Write(b []byte) (n int, err error) {
	return w.buf.Write(b)
}

// Complete writes the buffered response data to the actual http.ResponseWriter
func (w *htmlResponseWriter) Complete() error {
	data := map[string]string{
		"Title": "???",
		"Body":  w.buf.String(),
	}

	// Write the formatted result to a buffer. We do this intermediary step so we
	// can recalculate the Content-Length.
	var body bytes.Buffer
	if err := t.Execute(w.ResponseWriter, data); err != nil {
		return err
	}

	w.Header().Set("Content-Length", strconv.Itoa(body.Len()))
	_, err := body.WriteTo(w.ResponseWriter)
	return err
}

// wrap intercepts the response from http.FileServer and makes it a complete
// HTML document.
func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := &htmlResponseWriter{w, new(bytes.Buffer)}
		h.ServeHTTP(resp, r)
		if err := resp.Complete(); err != nil {
			log.Println("could not write template data", err)
		}
	})
}

func main() {
	http.ListenAndServe(":8000", wrap(http.FileServer(http.Dir("."))))
}
