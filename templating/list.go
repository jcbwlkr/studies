// +build list

package main

import (
	"log"
	"os"
	"text/template"
)

const tmpl = `Template!

{{ range $k, $v := . -}}
	# {{ $k }}
	{{ range $k, $v := . -}}
		* {{ $k }} = {{ $v }}
	{{ end }}
{{ end }}
`

func main() {
	data := map[string]map[string]string{
		"foo1": map[string]string{
			"item1": "baz1",
			"item2": "baz2",
		},
		"foo2": map[string]string{
			"item1": "baz1",
			"item2": "baz2",
		},
	}

	t := template.Must(template.New("foo").Parse(tmpl))

	err := t.Execute(os.Stdout, data)

	if err != nil {
		log.Fatal(err)
	}
}
