package main

import (
	"log"
	"os"
	"text/template"
)

const tmpl = `My Template!

Data is: {{ . }}
Literal braces: &lcub;&lcub;
`

func main() {
	t := template.Must(template.New("foo").Parse(tmpl))

	if err := t.Execute(os.Stdout, "data"); err != nil {
		log.Fatal(err)
	}
}
