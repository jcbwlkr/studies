// +build escaping

package main

import (
	"fmt"
	"html/template"
	"os"
)

const (
	urlStr  = `c:/tmp/test/100 Bullets 02/00.jpg`
	tmplStr = `{{define "T"}}
	<a href="file://{{ url . }}">{{ . }} </a> {{end}}`
)

func main() {
	funcs := map[string]interface{}{
		"url": func(s string) template.URL {
			return template.URL(s)
		},
	}
	t := template.Must(template.New("foo").Funcs(funcs).Parse(tmplStr))

	fmt.Print("Got:")
	t.ExecuteTemplate(os.Stdout, "T", urlStr)
}
