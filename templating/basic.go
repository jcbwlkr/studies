package main

import (
	"log"
	"os"
	"text/template"
)

const tmpl = `My Template!

Name is: {{ .Name }}
ID is: {{ .ID }}
`

func main() {
	data := struct {
		ID   int
		Name string
	}{
		ID:   123,
		Name: "foobar",
	}

	t := template.Must(template.New("foo").Parse(tmpl))

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
