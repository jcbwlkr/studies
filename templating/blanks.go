// +build blanks

package main

import (
	"fmt"
	"html/template"
	"os"
)

const str = `Names!

Something.
{{- range . }}

* Kid: {{ . }}
{{- end }}
`

func main() {
	t := template.Must(template.New("foo").Parse(str))

	fmt.Print("Got:")

	data := []string{
		"Kell",
		"Carter",
		"Rory",
	}
	t.ExecuteTemplate(os.Stdout, "foo", data)
}
